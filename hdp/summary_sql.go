package hdp

const (
	symmary_ads_sale_agg_per_day_sql = `
select store_code, store_name, cnt, qty, total from ads_sale_agg_per_day where pk = ?
`

	summary_ads_payment_agg_per_day_sql = `
select store_code, store_name, qty, total from ads_payment_agg_per_day where pk = ?
`

	summary_ads_fee_agg_per_day_sql = `
select store_code, store_name, t7, t8, t9 from ads_fee_agg_per_day where pk = ?
`

	summary_g_gm_entry_sql = `
select store_code, store_name, in_total from g_flow_entry where pk in (?, ?) order by pk desc
`

	summary_gather_sql = `
select
	storecode as store_code,
	x_is_payment,
	sum(total) as fee,
	ym
from
	dwm_base_fee
where
	ym > date_format(DATE_SUB(CURDATE(), INTERVAL 15 MONTH), '%Y%m') and ym <= date_format(CURDATE(), '%Y%m')
	and storecode = ?
group by storeCode,  ym, x_is_payment
order by ym, x_is_payment
`
)
