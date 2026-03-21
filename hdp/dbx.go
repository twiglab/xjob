package hdp

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type DBx struct {
	db *sqlx.DB
}

func NewDBx(name, dsn string) (*DBx, error) {
	db, err := sqlx.Connect(name, dsn)
	if err != nil {
		return nil, err
	}
	return &DBx{db: db}, nil
}

func (s *DBx) Close() error {
	return s.db.Close()
}

func (s *DBx) SaleAgg(code, dt string) (r SaleRecord, err error) {
	pk := code + "_" + dt
	err = s.db.Get(&r, symmary_ads_sale_agg_per_day_sql, pk)
	return
}

func (s *DBx) PaymentAgg(code, dt string) (r PaymentRecord, err error) {
	pk := code + "_" + dt
	err = s.db.Get(&r, summary_ads_payment_agg_per_day_sql, pk)
	return
}

func (s *DBx) FeeAgg(code, dt string) (r FeeRecord, err error) {
	pk := code + "-" + dt
	err = s.db.Get(&r, summary_ads_fee_agg_per_day_sql, pk)
	return
}

func (s *DBx) GmEntry(code, dt string) (r GmRecord, err error) {
	pk := code + "-" + dt + "-in"
	err = s.db.Get(&r, summary_g_gm_entry_sql, pk)
	return
}

func (s *DBx) GatherAgg(code string) (rs []GatherRecord, err error) {
	err = s.db.Select(&rs, summary_gather_sql, code)
	return
}
