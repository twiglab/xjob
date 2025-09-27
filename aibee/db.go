package aibee

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
	StoreName string
	DT        string
	InTotal   int
}

func (p CreateGmEntryParams) PK() string {
	return p.StoreCode + "-" + p.DT + "-" + "in"
}

type Queries struct {
	db *sql.DB
}

func (q *Queries) Close() error {
	return q.db.Close()
}

func (q *Queries) CreateGmEntry(ctx context.Context, arg CreateGmEntryParams) error {
	_, err := q.db.ExecContext(ctx, createGmEntry, arg.PK(), arg.StoreCode, arg.StoreName, arg.DT, arg.InTotal)
	return err
}

const createGmEntry = `replace INTO g_gm_entry(pk, store_code, store_name, dt, in_total) VALUES (?,?,?,?,?)`
