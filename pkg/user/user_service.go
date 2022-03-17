package user

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/phihu/ponygr.am/pkg/database"
	"github.com/phihu/ponygr.am/pkg/log"
	"github.com/pkg/errors"
)

type (
	Service interface {
		ByID(ctx context.Context, id uuid.UUID) (*User, error)
		ByEmail(ctx context.Context, email string) (*User, error)
		ByHandle(ctx context.Context, handle string) (*User, error)
		Create(ctx context.Context, u *User) error
	}
	Config  interface{}
	service struct {
		cfg Config
		db  database.DB
	}
)

func NewService(cfg Config, db database.DB) (Service, error) {
	return &service{
		cfg: cfg,
		db:  db,
	}, nil
}

func (svc *service) ByID(ctx context.Context, id uuid.UUID) (*User, error) {
	var u *User
	var err error
	txErr := database.WithTxRO(ctx, svc.db, func(ctx context.Context, dbtx database.Tx) {
		tx := &Tx{Tx: dbtx}
		u, err = tx.ByID(ctx, id)
		if err != nil {
			// nolint: errcheck
			log.ErrLog(ctx).Log("msg", "database error",
				"err", err,
			)
			return
		}
	})
	if txErr != nil {
		return nil, errors.Wrap(txErr, "tx error")
	}
	return u, err
}

func (svc *service) ByEmail(ctx context.Context, email string) (*User, error) {
	var u *User
	var err error
	txErr := database.WithTxRO(ctx, svc.db, func(ctx context.Context, dbtx database.Tx) {
		tx := &Tx{Tx: dbtx}
		u, err = tx.ByEmail(ctx, email)
		if err != nil {
			// nolint: errcheck
			log.ErrLog(ctx).Log("msg", "database error",
				"err", err,
			)
			return
		}
	})
	if txErr != nil {
		return nil, errors.Wrap(txErr, "tx error")
	}
	return u, err
}

func (svc *service) ByHandle(ctx context.Context, handle string) (*User, error) {
	var u *User
	var err error
	txErr := database.WithTxRO(ctx, svc.db, func(ctx context.Context, dbtx database.Tx) {
		tx := &Tx{Tx: dbtx}
		u, err = tx.ByHandle(ctx, handle)
		if err != nil {
			// nolint: errcheck
			log.ErrLog(ctx).Log("msg", "database error",
				"err", err,
			)
			return
		}
	})
	if txErr != nil {
		return nil, errors.Wrap(txErr, "tx error")
	}
	return u, err
}

func (svc *service) Create(ctx context.Context, u *User) error {
	if u.Handle == "" {
		return errors.New("cannot create user without handle")
	}
	var err error
	u.ID, err = uuid.NewRandom()
	if err != nil {
		return errors.Wrap(err, "error generating user ID")
	}
	u.Created = time.Now().UTC()
	u.Status = StatusNew

	txErr := database.WithTxRW(ctx, svc.db, func(ctx context.Context, dbtx database.Tx) bool {
		tx := &Tx{Tx: dbtx}
		err = tx.Create(ctx, u)
		if err != nil {
			err = errors.Wrap(err, "error creating user")
			return false
		}
		err = tx.SetHandle(ctx, u)
		if err != nil {
			err = errors.Wrap(err, "error setting handle")
			return false
		}
		err = tx.SetStatus(ctx, u)
		if err != nil {
			err = errors.Wrap(err, "error setting status")
			return false
		}
		err = tx.SetSetting(ctx, u)
		if err != nil {
			err = errors.Wrap(err, "error saving user settings")
			return false
		}
		return true
	})
	if txErr != nil {
		return errors.Wrap(txErr, "tx error")
	}
	if err != nil {
		return errors.Wrap(err, "db error")
	}
	return nil
}
