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

func holiday(h Holiday) string {
	if h.NotFound {
		return ""
	}

	if h.IsOffDay {
		return h.Name + "（休）"
	}
	return h.Name + "（班）"
}

const summaryTpl = `
# {{ .Param.StoreName }}（{{ .Param.StoreCode }}）运营日报 {{ .Yestoday.Format "2006.01.02" }} {{ .Yestoday | weekday }} {{.Holiday | holiday}}
>**{{ .Sale.Cnt }}** 个商户，上报 **{{ .Sale.Qty }}** 单，销售额 **{{ .Sale.Total | wan }}** 万元
>当日核销 **{{.Pay.Qty}}** 笔，共 **{{.Pay.Total | wan}}** 万元
{{- if .Gm.InTotal }}
>营业期间总客流 **{{.Gm.InTotal}}** 人次（入）{{- if .Gm.InTotalLast }} 上周同期 **{{.Gm.InTotalLast}}** {{end}} 人次
{{ end }}

## 前15个月收缴率（单位：万元）
| 月份 | 收缴率  | 未收 | 已收 | 应收 |
| :-----: | :----: | :----: | :-----: | :-----: |
{{ range .Gr.Table -}}
| {{.Ym}} | **{{.Rate | rate}}** | {{.Fee0 | wan}} | {{.Fee1 | wan}} | {{.Total | wan}} |
{{ end }}
>{{ .Yestoday.Year }} 年度，已出账金额（总应收）**{{.Fee.T9 | wan}}** 万元，当前欠款 **{{.Fee.T7 | wan}}** 万元，到期已收 **{{ .Fee.T8 | wan }}** 万元，收缴率为 **{{ .Fee | yearRecvRate}}**
`

func SummaryTpl() *template.Template {
	tpl, _ := template.New("summary").
		Funcs(template.FuncMap{
			"weekday":      weekday,
			"wan":          wan,
			"yearRecvRate": yearRecvRate,
			"rate":         rate,
			"holiday":      holiday,
		}).Parse(summaryTpl)
	return tpl
}
