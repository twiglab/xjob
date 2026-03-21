package hdp

import (
	"context"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Store struct {
	db *sqlx.DB
}

func NewStore(name, dsn string) (*Store, error) {
	db, err := sqlx.Connect(name, dsn)
	if err != nil {
		return nil, err
	}
	return &Store{db: db}, nil
}

func (s *Store) Close() error {
	return s.db.Close()
}

func (s *Store) SaleAgg(dt string) ([]SaleRecord, error) {
	rs, err := s.db.QueryContext(context.Background(), ads_sale_agg_per_day_sql, dt)
	if err != nil {
		return nil, err
	}
	defer rs.Close()

	var res []SaleRecord
	for rs.Next() {
		var sr SaleRecord
		err := rs.Scan(
			&sr.StoreCode,
			&sr.StoreName,
			&sr.Cnt,
			&sr.Qty,
			&sr.Total,
		)
		if err != nil {
			return res, err
		}
		res = append(res, sr)
	}
	return res, err
}

func (s *Store) PaymentAgg(dt string) ([]PaymentRecord, error) {
	rs, err := s.db.QueryContext(context.Background(), ads_payment_agg_per_day_sql, dt)
	if err != nil {
		return nil, err
	}
	defer rs.Close()

	var res []PaymentRecord
	for rs.Next() {
		var sr PaymentRecord
		err := rs.Scan(
			&sr.StoreCode,
			&sr.StoreName,
			&sr.Qty,
			&sr.Total,
		)
		if err != nil {
			return res, err
		}
		res = append(res, sr)
	}
	return res, err
}

func (s *Store) FeeAgg(dt string) ([]FeeRecord, error) {
	rs, err := s.db.QueryContext(context.Background(), ads_fee_agg_per_day_sql, dt)
	if err != nil {
		return nil, err
	}
	defer rs.Close()

	var res []FeeRecord
	for rs.Next() {
		var sr FeeRecord
		err := rs.Scan(
			&sr.StoreCode,
			&sr.StoreName,

			&sr.T7,
			&sr.T8,
			&sr.T9,
		)
		if err != nil {
			return res, err
		}
		res = append(res, sr)
	}
	return res, err
}

func (s *Store) GmEntry(dt string) (a []GmRecord, err error) {
	err = s.db.Select(&a, g_gm_entry_sql, dt)
	return
}
