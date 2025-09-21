package dbop

import (
	"context"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

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

func (q *Queries) CreateGmEntry(ctx context.Context, arg CreateGmEntryParams) error {
	_, err := q.db.ExecContext(ctx, createGmEntry, arg.PK(), arg.StoreCode, arg.PickTime, arg.InTotal)
	return err
}

const createGmEntry = `replace INTO g_gm_entry(pk, store_code, pick_time, in_total) VALUES (?, ?, ?, ?)`
