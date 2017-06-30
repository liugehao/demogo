package lib

import "database/sql"

var DB *sql.DB
const T01Insertsql string = "INSERT INTO origin.origin_data (sta,sta_tm, category, sn,sn_tm,v_up,v_down,err) VALUES ($1, $2, $3, $4, $5, $6, $7, $8);"
const StaInsertsql string = "INSERT INTO origin.origin_data (sta,sta_tm, category, batt,ver) VALUES ($1, $2, $3, $4, $5);"
const T00Insertsql string = "INSERT INTO origin.origin_data (sta,sta_tm, category,sn_tm,pressure) VALUES ($1, $2, $3, $4, $5);"

type RouteDevice struct {
	STA  string
	TM   string
	BATT string
	VER  string
}
type T00 struct {
	RouteDevice
	TM       string
	PRESSURE string
}
type T01 struct {
	RouteDevice
	Vup   string
	Vdown string
	ERR   string
	TM    string
	SN    string
}



func (t01 *T01) save() (res sql.Result, err error) {
	res, err = DB.Exec(T01Insertsql, t01.STA, t01.RouteDevice.TM, "T01", t01.SN, t01.TM, t01.Vup, t01.Vdown, t01.ERR)
	return
}

func (t00 *T00)save() (res sql.Result, err error) {
	res, err = DB.Exec(T00Insertsql, t00.STA, t00.RouteDevice.TM, "T00", t00.TM, t00.PRESSURE)
	return
}

func (sta RouteDevice) save() (res sql.Result, err error) {
	res, err = DB.Exec(StaInsertsql, sta.STA, sta.TM, "STA", sta.BATT, sta.VER);
	return
}