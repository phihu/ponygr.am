package database

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"database/sql"
	"io/ioutil"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/go-stack/stack"
	"github.com/phihu/ponygr.am/pkg/log"

	"fmt"
	"github.com/pkg/errors"
)

type (
	contextKey int

	// Config represents a configuration type, which can configure a database connection
	Config interface {
		DBDSN() string
		DBCACert() string
		DBConnMaxLifetime() time.Duration
	}

	// DB is a minimal DB connection interface
	DB interface {
		BeginTx(context.Context, *sql.TxOptions) (*sql.Tx, error)
	}

	// RowScanner can scan a result
	RowScanner interface {
		Scan(dest ...interface{}) error
	}

	// Tx represents a transaction
	Tx interface {
		ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
		PrepareContext(context.Context, string) (*sql.Stmt, error)
		QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
		QueryRowContext(context.Context, string, ...interface{}) *sql.Row
		StmtContext(context.Context, *sql.Stmt) *sql.Stmt
		Commit() error
		Rollback() error
	}

	// Logger is a minimal logging interface
	Logger interface {
		Log(...interface{}) error
	}

	loggingTx struct {
		Tx

		logger Logger
	}
)

const (
	contextKeyTx contextKey = iota
	contextKeyTxDebug
)

// ConnectDB connects to a DB with the given config
func ConnectDB(cfg Config) (*sql.DB, error) {
	if cfg.DBCACert() != "" {
		rootCertPool := x509.NewCertPool()
		pem, err := ioutil.ReadFile(cfg.DBCACert())
		if err != nil {
			return nil, errors.Wrap(err, "error reading CA cert")
		}
		if ok := rootCertPool.AppendCertsFromPEM(pem); !ok {
			return nil, errors.New("error adding PEM")
		}
		err = mysql.RegisterTLSConfig("default", &tls.Config{
			RootCAs: rootCertPool,
		})
		if err != nil {
			return nil, errors.Wrap(err, "error registering TLS config")
		}
	}
	db, err := sql.Open("mysql", cfg.DBDSN())
	if err != nil {
		return nil, errors.Wrap(err, "error creating db connection")
	}
	db.SetMaxOpenConns(2)
	db.SetMaxIdleConns(1)
	db.SetConnMaxLifetime(cfg.DBConnMaxLifetime())
	err = db.Ping()
	if err != nil {
		return nil, errors.Wrap(err, "error connecting to db")
	}

	return db, nil
}

// ContextWithTx assigns a sql transaction to the given
// context.
//
// Subsequent calls during a process should be aware if
// they happen inside a transaction.
//
// The ContextTx related functions provide a simple way
// to have a transaction shared in a context subtree.
func ContextWithTx(ctx context.Context, tx Tx) context.Context {
	return context.WithValue(ctx, contextKeyTx, tx)
}

// ContextTx returns the active transaction from the context or nil if there is none
//
// This is a utility function, but you should use the WithTxRO and WithTxRW functions
// instead.
func ContextTx(ctx context.Context) Tx {
	if ctx.Value(contextKeyTx) == nil {
		return nil
	}
	if tx, ok := ctx.Value(contextKeyTx).(Tx); ok {
		return tx
	}
	return nil
}

// ContextWithDebug activates the debug mode on the database handling
func ContextWithDebug(ctx context.Context) context.Context {
	return context.WithValue(ctx, contextKeyTxDebug, true)
}

func WithTxRO(ctx context.Context, db DB, f func(ctx context.Context, tx Tx)) error {
	var tx Tx
	debug, ok := ctx.Value(contextKeyTxDebug).(bool)
	debug = ok && debug
	if ctxTx := ContextTx(ctx); ctxTx == nil {
		if debug {
			// nolint: errcheck
			log.DebugLog(ctx).Log("msg", "new RO transaction",
				"trace", stack.Trace(),
			)
		}
		dbtx, err := db.BeginTx(ctx, &sql.TxOptions{
			Isolation: sql.LevelRepeatableRead,
		})
		if err != nil {
			return errors.Wrap(err, "error on begin tx")
		}
		defer func() {
			err = dbtx.Rollback()
			if err != nil {
				// nolint: errcheck
				log.ErrLog(ctx).Log("msg", "error on rollback",
					"err", err,
				)
			}
		}()
		if debug {
			tx = LoggingTx(dbtx, log.DebugLog(ctx))
		} else {
			tx = dbtx
		}
		ctx = ContextWithTx(ctx, tx)
	} else {
		if debug {
			// nolint: errcheck
			log.DebugLog(ctx).Log("msg", "existing transaction",
				"trace", stack.Trace(),
			)
		}
		tx = ctxTx
	}
	f(ctx, tx)

	return nil
}

func WithTxRW(ctx context.Context, db DB, f func(ctx context.Context, t Tx) bool) error {
	debug, ok := ctx.Value(contextKeyTxDebug).(bool)
	debug = ok && debug
	var top bool
	var tx Tx
	var commit, committed bool
	if ctxTx := ContextTx(ctx); ctxTx == nil {
		if debug {
			// nolint: errcheck
			log.DebugLog(ctx).Log("msg", "new RW transaction",
				"trace", stack.Trace(),
			)
		}
		dbtx, err := db.BeginTx(ctx, &sql.TxOptions{
			Isolation: sql.LevelSerializable,
		})
		if err != nil {
			return errors.Wrap(err, "error on begin tx")
		}
		defer func() {
			if committed {
				return
			}
			err = dbtx.Rollback()
			if err != nil {
				// nolint: errcheck
				log.ErrLog(ctx).Log("msg", "error on rollback",
					"err", err,
				)
			}
		}()
		top = true
		if debug {
			tx = LoggingTx(dbtx, log.DebugLog(ctx))
		} else {
			tx = dbtx
		}
		ctx = ContextWithTx(ctx, tx)
	} else {
		if debug {
			// nolint: errcheck
			log.DebugLog(ctx).Log("msg", "existing transaction",
				"trace", stack.Trace(),
			)
		}
		tx = ctxTx
	}
	commit = f(ctx, tx)

	if top && commit {
		err := tx.Commit()
		if err != nil {
			return errors.Wrap(err, "error on commit")
		}
		committed = true
	}

	return nil
}

var _ Tx = &sql.Tx{}

func LoggingTx(tx Tx, logger Logger) Tx {
	return &loggingTx{
		Tx:     tx,
		logger: logger,
	}
}

func (tx *loggingTx) PrepareContext(ctx context.Context, query string) (*sql.Stmt, error) {
	// nolint: errcheck
	tx.logger.Log("msg", "prepare",
		"query", query,
	)
	return tx.Tx.PrepareContext(ctx, query)
}

func (tx *loggingTx) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	// nolint: errcheck
	tx.logger.Log("msg", "query",
		"query", query,
		"args", fmt.Sprintf("%v", args),
	)
	return tx.Tx.QueryContext(ctx, query, args...)
}

func (tx *loggingTx) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	// nolint: errcheck
	tx.logger.Log("msg", "query",
		"query", query,
		"args", fmt.Sprintf("%v", args),
	)
	return tx.Tx.QueryRowContext(ctx, query, args...)
}
