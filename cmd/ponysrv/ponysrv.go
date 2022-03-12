package main

import (
	"context"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-kit/log"
	"github.com/phihu/ponygr.am/cmd/ponysrv/config"
	"github.com/phihu/ponygr.am/pkg/database"
	plog "github.com/phihu/ponygr.am/pkg/log"
	"github.com/phihu/ponygr.am/pkg/session"
	"github.com/phihu/ponygr.am/pkg/user"
	"github.com/pkg/errors"
)

const (
	AppName = "ponysrv"
)

var (
	Version = "dev"
	BuildID = "detached"
)

func main() {
	os.Exit(runWithCode())
}

// app is the main application data type
//
// It contains all necessary dependencies. Everything is initialized here.
type app struct {
	ctx    context.Context
	cancel context.CancelFunc

	// our loggers
	errLog, infoLog, debugLog log.Logger

	cfg *config.Config

	// The default HTTP client, sharing connection pool
	cl *http.Client

	srv    *http.Server
	router chi.Router

	recoverer func()

	db database.DB

	sessionSvc session.Service
	userSvc    user.Service
}

func runWithCode() int {
	app := &app{}
	app.ctx, app.cancel = context.WithCancel(context.Background())
	defer app.cancel()

	app.errLog = log.NewSyncLogger(log.NewLogfmtLogger(os.Stderr))
	app.errLog = log.WithPrefix(app.errLog,
		"t", log.DefaultTimestampUTC,
		"level", "error",
		"appName", AppName,
		"appVersion", Version,
		"appBuild", BuildID,
		"pid", os.Getpid(),
		"caller", log.Caller(5),
	)
	app.infoLog = log.NewSyncLogger(log.NewLogfmtLogger(os.Stdout))
	app.infoLog = log.WithPrefix(app.infoLog,
		"t", log.DefaultTimestampUTC,
		"level", "info",
		"appName", AppName,
		"appVersion", Version,
		"appBuild", BuildID,
		"pid", os.Getpid(),
		"caller", log.Caller(5),
	)

	var err error
	app.cfg, err = config.LoadConfig()
	if err != nil {
		// nolint: errcheck
		app.errLog.Log("msg", "error loading config",
			"err", err)
		return 1
	}

	// setup base context
	app.ctx = plog.ContextWithErrLog(app.ctx, app.errLog)
	app.ctx = plog.ContextWithInfoLog(app.ctx, app.infoLog)
	if app.cfg.Debug {
		app.debugLog = log.With(app.infoLog,
			"level", "debug")
		// nolint: errcheck
		app.debugLog.Log("msg", "debug log enabled")
		app.ctx = plog.ContextWithDebugLog(app.ctx, app.debugLog)
	}

	app.cl = &http.Client{
		Timeout: app.cfg.HTTP.Client.Timeout,
	}

	app.srv = &http.Server{
		Addr: app.cfg.HTTP.Addr,

		ReadTimeout:  app.cfg.HTTP.ReadTimeout,
		WriteTimeout: app.cfg.HTTP.WriteTimeout,
		IdleTimeout:  app.cfg.HTTP.IdleTimeout,
	}
	app.srv.BaseContext = func(_ net.Listener) context.Context {
		return app.ctx
	}
	app.router = chi.NewRouter()
	if app.cfg.Debug {
		app.router.Use(plog.DebugMiddleware())
	}
	app.recoverer = func() {
		if err := recover(); err != nil {
			// nolint: errcheck
			app.errLog.Log("msg", "recovered from panic",
				"panic", err,
			)
		}
	}

	app.db, err = database.ConnectDB(app.cfg)
	if err != nil {
		// nolint: errcheck
		plog.ErrLog(app.ctx).Log("msg", "error connecting to db",
			"err", err)
		return 1
	}

	app.sessionSvc, err = session.NewService(app.cfg)
	if err != nil {
		// nolint: errcheck
		plog.ErrLog(app.ctx).Log("msg", "error initializing session service",
			"err", err)
		return 1
	}

	app.userSvc, err = user.NewService(app.cfg, app.db)
	if err != nil {
		// nolint: errcheck
		plog.ErrLog(app.ctx).Log("msg", "error initializing user service",
			"err", err)
		return 1
	}

	err = Mount(app.router,
		app.cfg,
		app.sessionSvc,
		app.userSvc,
	)
	if err != nil {
		// nolint: errcheck
		plog.ErrLog(app.ctx).Log("msg", "error mounting endpoints",
			"err", err,
		)
		return 1
	}

	app.srv.Handler = app.router

	err = serve(app.ctx, app.srv)
	if err != nil {
		// nolint: errcheck
		plog.ErrLog(app.ctx).Log("msg", "error serving",
			"err", err,
		)
		return 1
	}

	return 0
}

func serve(ctx context.Context, srv *http.Server) error {
	// nolint: errcheck
	plog.InfoLog(ctx).Log("msg", "starting HTTP server",
		"addr", srv.Addr,
	)

	errchan := make(chan error, 1)
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			// nolint: errcheck
			plog.ErrLog(ctx).Log("msg", "error on listen and serve",
				"err", err,
			)
			errchan <- err
		}
	}()

waitloop:
	for {
		select {
		case <-ctx.Done():
			break waitloop

		case err := <-errchan:
			if err == http.ErrServerClosed {
				continue
			}
			return errors.Wrap(err, "error on server start")
		}
	}

	var err error

	ctxShutdown, cancelShutdown := context.WithTimeout(ctx, 5*time.Second)
	defer cancelShutdown()

	if err = srv.Shutdown(ctxShutdown); err != nil {
		// nolint: errcheck
		plog.ErrLog(ctx).Log("msg", "error during shutdown",
			"err", err,
		)
	}

	// nolint: errcheck
	plog.InfoLog(ctx).Log("msg", "server shut down")

	if err == http.ErrServerClosed {
		err = nil
	}

	return err
}
