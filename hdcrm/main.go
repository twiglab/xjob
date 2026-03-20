package main

import (
	"encoding/json/v2"
	"flag"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/jmoiron/sqlx"
)

const sql = `
select
	a.storeName     as store_name,
	a.serialNumber  as seria_number ,
	a.signboard     as sign_board,
-- 	a.beginDate     as begin_date,
-- 	a.endDate       as end_date,
	a.POSITIONCODES as pos_code ,
	a.positionType  as pos_type,
	a.floorName     as floor_name,
	a.floorCode     as floor_code,
	a.bizTypeName   as biz_type_name
from
	dws.dwd_contract a
where
	a.storeCode = '1006'
`

type Data struct {
	StoreName   string `json:"store_name" db:"store_name"`
	SeriaNumber string `json:"seria_number" db:"seria_number"`
	SignBoard   string `json:"sign_board" db:"sign_board"`
	//	BeginDate   time.Time `json:"begin_date" db:"begin_date"`
	//	EndDate     time.Time `json:"end_date" db:"end_date"`
	PosCode     string `json:"pos_code" db:"pos_code"`
	PosType     string `json:"pos_type" db:"pos_type"`
	FloorName   string `json:"floor_name" db:"floor_name"`
	FloorCode   string `json:"floor_code" db:"floor_code"`
	BizTypeName string `json:"biz_type_name" db:"biz_type_name"`
}

var (
	dsn  string
	addr string
)

func init() {
	flag.StringVar(&dsn, "dsn", "", "dsn")
	flag.StringVar(&addr, "addr", ":10001", "dsn")
}

func main() {
	flag.Parse()

	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("GET /crm/data/store/1006", func(w http.ResponseWriter, r *http.Request) {
		var x []Data
		if err := db.Select(&x, sql); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		_ = json.MarshalWrite(w, x)
	})

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}
