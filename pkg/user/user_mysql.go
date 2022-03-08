package user

import (
	"context"
	"time"

	"github.com/phihu/ponygr.am/pkg/database"
	"github.com/pkg/errors"
)

type (
	Tx struct {
		database.Tx
	}
)

func (tx *Tx) Create(ctx context.Context, u *User) error {
	const stmtStr = `
INSERT INTO user (id, created) VALUES(?, ?)
	`
	stmt, err := tx.PrepareContext(ctx, stmtStr)
	if err != nil {
		return errors.Wrap(err, "error on prepare")
	}
	defer stmt.Close()
	b, _ := u.ID.MarshalBinary()
	_, err = stmt.ExecContext(ctx, b, u.Created.Format("2006-01-02 15:04:05"))
	if err != nil {
		return errors.Wrap(err, "error on exec")
	}
	return nil
}

func (tx *Tx) SetHandle(ctx context.Context, u *User) error {
	const insertHandleStmt = `
INSERT INTO user_handle (id_user, created, handle) VALUES
(?, ?, ?)
	`
	stmt, err := tx.PrepareContext(ctx, insertHandleStmt)
	if err != nil {
		return errors.Wrap(err, "error on prepare")
	}
	defer stmt.Close()

	b, _ := u.ID.MarshalBinary()
	_, err = stmt.ExecContext(ctx,
		b,
		time.Now().Format("2006-01-02 15:04:05"),
		u.Handle,
	)
	if err != nil {
		return errors.Wrap(err, "error on exec")
	}
	return nil
}

func (tx *Tx) SetStatus(ctx context.Context, u *User) error {
	const insertStatusStmt = `
INSERT INTO user_status (id_user, created, status) VALUES
(?, ?, ?)
	`
	stmt, err := tx.PrepareContext(ctx, insertStatusStmt)
	if err != nil {
		return errors.Wrap(err, "error on prepare")
	}
	defer stmt.Close()

	b, _ := u.ID.MarshalBinary()
	_, err = stmt.ExecContext(ctx,
		b,
		time.Now().UTC().UnixNano(),
		u.Status,
	)
	if err != nil {
		return errors.Wrap(err, "error on exec")
	}
	return nil
}

func (tx *Tx) SetSetting(ctx context.Context, u *User) error {
	const insertSettingStmt = `
INSERT INTO user_setting (id_user, created, lang) VALUES
(?, ?, ?)
	`
	stmt, err := tx.PrepareContext(ctx, insertSettingStmt)
	if err != nil {
		return errors.Wrap(err, "error on prepare")
	}
	defer stmt.Close()

	b, _ := u.ID.MarshalBinary()
	_, err = stmt.ExecContext(ctx,
		b,
		time.Now().UTC().UnixNano(),
		u.Setting.Lang.String(),
	)
	if err != nil {
		return errors.Wrap(err, "error on exec")
	}
	return nil
}
