package hdp

import (
	"cmp"
	"fmt"
	"slices"
	"text/template"
	"time"
)

const Tpl = `
{{ .Yestoday.Format "2006.01.02" }} {{ .Yestoday | weekday }}
{{ range .Records }}
>{{ .Store.StoreName }} {{ .Sale.Cnt }} 个商户，上报 {{ .Sale.Qty }} 单，销售额 <font color="warning"> {{ .Sale.Total | wan }} </font>万元
>>本年总欠款 <font color="warning"> {{.Fee.T7 | wan}} </font> 万元，到期已收 <font color="warning"> {{ .Fee.T8 | wan }} </font>万元，收缴率 <font color="warning"> {{ .Fee | yearRecvRate}} </font>
>>当日核销 {{.Pay.Qty}} 笔，共<font color="warning"> {{.Pay.Total | wan}} </font> 万元
>当日客流 {{.Gm.InTotal}} 人次
{{ end }}
`

// >宜悦城客流 {{.Others.yyc.InTotal}} 人次

/*
<font color="warning"> </font>
>>截至当日欠款总计 <font color="warning"> {{.Fee.T4 | wan}} </font> 万元，到期已收 <font color="warning"> {{ .Fee.T5 | wan }} </font>万元，收缴率 <font color="warning"> {{ .Fee | recvRate}} </font>
*/

type StoreX struct {
	StoreCode string
	StoreName string
}

var StoreSet = []StoreX{
	{StoreCode: "1001", StoreName: "宜悦里G66"},
	{StoreCode: "1002", StoreName: "南瑞路宜月里"},
	{StoreCode: "1006", StoreName: "金陵长乐坊"},
}

type Record struct {
	Store StoreX
	Fee   FeeRecord
	Pay   PaymentRecord
	Sale  SaleRecord
	Gm    GmRecord
}

type Outline struct {
	Yestoday time.Time

	fee  []FeeRecord
	pay  []PaymentRecord
	sale []SaleRecord
	gm   []GmRecord
}

func (o Outline) Records() (rs []Record) {
	for _, x := range StoreSet {
		var r Record
		r.Store = x

		if i, ok := slices.BinarySearchFunc(o.sale, x.StoreCode, func(sr SaleRecord, k string) int { return cmp.Compare(sr.StoreCode, k) }); ok {
			r.Sale = o.sale[i]
		}

		if i, ok := slices.BinarySearchFunc(o.fee, x.StoreCode, func(fr FeeRecord, k string) int { return cmp.Compare(fr.StoreCode, k) }); ok {
			r.Fee = o.fee[i]
		}

		if i, ok := slices.BinarySearchFunc(o.pay, x.StoreCode, func(pr PaymentRecord, k string) int { return cmp.Compare(pr.StoreCode, k) }); ok {
			r.Pay = o.pay[i]
		}

		rs = append(rs, r)
	}

	return
}

/*
func (o Outline) Records() (rs []Record) {

		for _, x := range o.sale {
			var r Record
			r.Sale = x

			if i, ok := slices.BinarySearchFunc(o.fee, x.StoreCode, func(fr FeeRecord, k string) int { return cmp.Compare(fr.StoreCode, k) }); ok {
				r.Fee = o.fee[i]
			}

			if i, ok := slices.BinarySearchFunc(o.pay, x.StoreCode, func(pr PaymentRecord, k string) int { return cmp.Compare(pr.StoreCode, k) }); ok {
				r.Pay = o.pay[i]
			}

			rs = append(rs, r)
		}

		return
	}
*/

func MakeOutline(t time.Time, fee []FeeRecord, pay []PaymentRecord, sale []SaleRecord, gm []GmRecord) Outline {
	return Outline{Yestoday: t, fee: fee, pay: pay, sale: sale, gm: gm}
}

func weekday(t time.Time) string {
	w := t.Weekday()
	switch w {
	case time.Monday:
		return "周一"
	case time.Tuesday:
		return "周二"
	case time.Wednesday:
		return "周三"
	case time.Thursday:
		return "周四"
	case time.Friday:
		return "周五"
	case time.Saturday:
		return "周六"
	case time.Sunday:
		return "周日"
	}
	return ""
}

func wan(total float64) string {
	return fmt.Sprintf("%.2f", total/10000)
}

func recvRate(fr FeeRecord) string {
	if fr.T5 == 0 {
		return fmt.Sprintf("%.2f%%", 0.0)
	}
	return fmt.Sprintf("%.2f%%", fr.T5/fr.T6*100)
}

func yearRecvRate(fr FeeRecord) string {
	if fr.T8 == 0 {
		return fmt.Sprintf("%.2f%%", 0.0)
	}
	return fmt.Sprintf("%.2f%%", fr.T8/fr.T9*100)
}

func AppTpl() *template.Template {
	tpl, _ := template.New("tpl").
		Funcs(template.FuncMap{
			"weekday":      weekday,
			"wan":          wan,
			"recvRate":     recvRate,
			"yearRecvRate": yearRecvRate,
		}).Parse(Tpl)
	return tpl
}
