package hdp

// ads_fee_agg_per_day
// ads_payment_agg_per_day
// ads_sale_agg_per_day
// g_gm_entry 客流采集

type SaleRecord struct {
	StoreCode string `db:"store_code"`
	StoreName string `db:"store_name"`

	Cnt   int     `db:"cnt"`
	Qty   float64 `db:"qty"`
	Total float64 `db:"total"`
}

type PaymentRecord struct {
	StoreCode string `db:"store_code"`
	StoreName string `db:"store_name"`

	Qty   int     `db:"qty"`
	Total float64 `db:"total"`
}

type FeeRecord struct {
	StoreCode string `db:"store_code"`
	StoreName string `db:"store_name"`

	// T4 float64 // 到期未收（开店以来总欠款） t4
	// T5 float64 // 到期已收（开店以来总已收）(t5)
	// T6 float64 // 到期应收（开店以来总应该收）(t6)

	// 年度(当年)
	T7 float64 `db:"t7"`
	T8 float64 `db:"t8"`
	T9 float64 `db:"t9"`
}

type GmRecord struct {
	StoreCode string `db:"store_code"`
	StoreName string `db:"store_name"`
	InTotal   int    `db:"in_total"`
}

// dwm_base_fee
type GatherRecord struct {
	StoreCode string  `db:"store_code"`
	Fee       float64 `db:"fee"`
	Ym        string  `db:"ym"`
	IsPayment int     `db:"x_is_payment"`
}

type GatherItem struct {
	Ym   string
	Fee0 float64
	Fee1 float64
}

func (g *GatherItem) Rate() float64 {
	return g.Fee1 / (g.Fee0 + g.Fee1)
}
