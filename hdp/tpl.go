package hdp

import (
	"fmt"
	"text/template"
	"time"
)

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

func yearRecvRate(fr FeeRecord) string {
	if fr.T8 == 0 {
		return fmt.Sprintf("%.2f%%", 0.0)
	}
	return fmt.Sprintf("%.2f%%", fr.T8/fr.T9*100)
}

func rate(f float64) string {
	if f == 0 {
		return fmt.Sprintf("%.2f%%", 0.0)
	}
	return fmt.Sprintf("%.2f%%", f*100)
}

const summaryTpl = `
# {{ .Param.StoreName }} 运营日报 {{ .Yestoday.Format "2006.01.02" }} {{ .Yestoday | weekday }}
>**{{ .Sale.Cnt }}** 个商户，上报 **{{ .Sale.Qty }}** 单，销售额 **{{ .Sale.Total | wan }}** 万元
>本年总欠款 **{{.Fee.T7 | wan}}** 万元，到期已收 **{{ .Fee.T8 | wan }}** 万元，收缴率 **{{ .Fee | yearRecvRate}}**
>当日核销 **{{.Pay.Qty}}** 笔，共 **{{.Pay.Total | wan}}** 万元
{{- if .Gm.InTotal }}
>营业期间总客流 **{{.Gm.InTotal}}** 人次（入）
{{ end }}

## 前12个月收缴率（单位：万元）
| 月份 | 未收 | 已收 | 收缴率 |
| :----- | :----: | -------: | :----- |
{{ range .Gr.Table -}}
| {{.Ym}} | {{.Fee0 | wan}} | {{.Fee1 | wan}} | **{{.Rate | rate}}** |
{{ end }}
`

func SummaryTpl() *template.Template {
	tpl, _ := template.New("summary").
		Funcs(template.FuncMap{
			"weekday":      weekday,
			"wan":          wan,
			"yearRecvRate": yearRecvRate,
			"rate":         rate,
		}).Parse(summaryTpl)
	return tpl
}
