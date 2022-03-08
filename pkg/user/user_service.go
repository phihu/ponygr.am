package user

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/phihu/ponygr.am/pkg/database"
	"github.com/pkg/errors"
)

type (
	Service interface {
		Create(ctx context.Context, u *User) error
	}
	Config  interface{}
	service struct {
		cfg Config
		db  database.DB
	}
)

func NewService(cfg Config, db database.DB) Service {
	return &service{
		cfg: cfg,
		db:  db,
	}
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
