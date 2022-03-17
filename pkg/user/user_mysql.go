package user

import (
	"context"
	"database/sql"

	"github.com/doug-martin/goqu/v9"
	"github.com/google/uuid"
	"golang.org/x/text/language"

	"time"

	"github.com/phihu/ponygr.am/pkg/database"
	"github.com/pkg/errors"
)

type (
	Tx struct {
		database.Tx
	}

	selectFunc func(*goqu.SelectDataset) *goqu.SelectDataset
)

func selectUser(fs ...selectFunc) *goqu.SelectDataset {
	q := goqu.Dialect("mysql").From(goqu.T("user").As("u")).Prepared(true).
		Select(
			"u.id",
			"e.email",
			"h.handle",

			"u.created",

			"s.status",

			"setting.lang",
		)
	for _, f := range fs {
		q = f(q)
	}
	return q
}

func withEmail(email string) selectFunc {
	return func(q *goqu.SelectDataset) *goqu.SelectDataset {
		return q.
			Join(goqu.T("user_email").As("e"),
				goqu.On(
					goqu.And(
						goqu.I("e..id_user").Eq(goqu.I("u.id")),
						goqu.I("e..handle").Eq(goqu.V(email)),
					),
				))
	}
}

func joinEmail() selectFunc {
	return func(q *goqu.SelectDataset) *goqu.SelectDataset {
		return q.
			Join(goqu.T("user_email").As("e"),
				goqu.On(
					goqu.And(
						goqu.I("e.id_user").Eq(goqu.I("u.id")),
						goqu.I("e.created").Eq(
							goqu.Select(goqu.MAX("created")).
								From("user_email").Where(goqu.I("id_user").Eq(goqu.I("u.id"))),
						),
					),
				))
	}
}

func withHandle(handle string) selectFunc {
	return func(q *goqu.SelectDataset) *goqu.SelectDataset {
		return q.
			Join(goqu.T("user_handle").As("h"),
				goqu.On(
					goqu.And(
						goqu.I("h.id_user").Eq(goqu.I("u.id")),
						goqu.I("h.handle").Eq(goqu.V(handle)),
					),
				))
	}
}

func joinHandle() selectFunc {
	return func(q *goqu.SelectDataset) *goqu.SelectDataset {
		return q.
			Join(goqu.T("user_handle").As("h"),
				goqu.On(
					goqu.And(
						goqu.I("h.id_user").Eq(goqu.I("u.id")),
						goqu.I("h.created").Eq(
							goqu.Select(goqu.MAX("created")).
								From("user_handle").Where(goqu.I("id_user").Eq(goqu.I("u.id"))),
						),
					),
				))
	}
}

func selectDefaults() selectFunc {
	return func(q *goqu.SelectDataset) *goqu.SelectDataset {
		return q.
			Join(goqu.T("user_status").As("s"),
				goqu.On(
					goqu.And(
						goqu.I("s.id_user").Eq(goqu.I("u.id")),
						goqu.I("s.created").Eq(
							goqu.Select(goqu.MAX("created")).
								From("user_status").Where(goqu.I("id_user").Eq(goqu.I("u.id"))),
						),
					),
				)).
			Join(goqu.T("user_setting").As("setting"),
				goqu.On(
					goqu.And(
						goqu.I("setting.id_user").Eq(goqu.I("u.id")),
						goqu.I("setting.created").Eq(
							goqu.Select(goqu.MAX("created")).
								From("user_setting").Where(goqu.I("id_user").Eq(goqu.I("u.id"))),
						),
					),
				))
	}
}

func scanSingle(row *sql.Row) (*User, error) {
	u := &User{}
	var lang string
	err := row.Scan(
		&u.ID,
		&u.Email,
		&u.Handle,
		&u.Created,
		&u.Status,

		&lang,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, errors.Wrap(err, "error on scan")
	}

	u.Setting.Lang = language.Make(lang)

	return u, nil
}

func (tx *Tx) ByID(ctx context.Context, id uuid.UUID) (*User, error) {
	b, _ := id.MarshalBinary()
	q := selectUser(
		joinEmail(),
		joinHandle(),
		selectDefaults(),
	).
		Where(goqu.I("u.id").Eq(goqu.V(b)))

	sqlQuery, args, err := q.ToSQL()
	if err != nil {
		return nil, errors.Wrap(err, "error generating query")
	}
	return scanSingle(tx.QueryRowContext(ctx, sqlQuery, args...))
}

func (tx *Tx) ByEmail(ctx context.Context, email string) (*User, error) {
	q := selectUser(
		withEmail(email),
		joinHandle(),
		selectDefaults(),
	)

	sqlQuery, args, err := q.ToSQL()
	if err != nil {
		return nil, errors.Wrap(err, "error generating query")
	}
	return scanSingle(tx.QueryRowContext(ctx, sqlQuery, args...))
}

func (tx *Tx) ByHandle(ctx context.Context, handle string) (*User, error) {
	q := selectUser(
		joinEmail(),
		withHandle(handle),
		selectDefaults(),
	)

	sqlQuery, args, err := q.ToSQL()
	if err != nil {
		return nil, errors.Wrap(err, "error generating query")
	}
	return scanSingle(tx.QueryRowContext(ctx, sqlQuery, args...))
}

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
