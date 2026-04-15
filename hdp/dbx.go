package hdp

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/twiglab/xjob/pfsdk"
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

func (s *DBx) GmEntry(code string, yestoday time.Time) (r GmRecord, err error) {
	var rs []GmRecord

	pkNow := code + "-" + pfsdk.DT(yestoday) + "-in"
	last := yestoday.Add(-7 * 24 * time.Hour)
	pkLast := code + "-" + pfsdk.DT(last) + "-in"
	if err = s.db.Select(&rs, summary_g_gm_entry_sql, pkNow, pkLast); err != nil {
		return
	}

	switch len(rs) {
	case 1:
		r = rs[0]
		if r.Pk == pkLast {
			// 当前的没取到，取到是上次的（7天前的)
			r.InTotal = 0
			r.InTotalLast = rs[0].InTotal
		}
	case 2:
		r = rs[0]
		r.InTotalLast = rs[1].InTotal
	}
	return
}

func (s *DBx) GatherAgg(code string) (rs []GatherRecord, err error) {
	err = s.db.Select(&rs, summary_gather_sql, code)
	return
}
