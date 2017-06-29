package lib

import "database/sql"

var DB *sql.DB

type RouteDevice struct {
	sta  string
	tm   string
	batt string
	ver  string
}
type T00 struct {
	RouteDevice
	tm    string
	value string
}
type T01 struct {
	RouteDevice
	Vup   string
	Vdown string
	err   string
	tm    string
	sn    string
}