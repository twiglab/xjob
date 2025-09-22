package dbop

import (
	"context"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Open(dn, dsn string) (*Queries, error) {
	db, err := sql.Open(dn, dsn)
	if err != nil {
		return nil, err
	}

	return From(db), nil
}

func From(db *sql.DB) *Queries {
	return &Queries{db: db}
}

type CreateGmEntryParams struct {
	StoreCode string
	PickTime  string
	InTotal   int
}

func (p CreateGmEntryParams) PK() string {
	return p.StoreCode + "-" + p.PickTime + "-" + "in"
}

type Queries struct {
	db *sql.DB
}

func (q *Queries) Close() error {
	return q.db.Close()
}

func (q *Queries) CreateGmEntry(ctx context.Context, arg CreateGmEntryParams) error {
	_, err := q.db.ExecContext(ctx, createGmEntry, arg.PK(), arg.StoreCode, arg.PickTime, arg.InTotal)
	return err
}

const createGmEntry = `replace INTO g_gm_entry(pk, store_code, pick_time, in_total) VALUES (?, ?, ?, ?)`
