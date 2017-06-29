package lib

import (
	"fmt"
)

func Save(data interface{}) {
	fmt.Println()
	fmt.Println(data)
	/*switch data.(type) {
	case RouteDevice:
	case T00:
	case T01:lib
	default:

	}
	println(data)*/
}

func saveT01(t01 T01) {
	DB.Exec("INSERT INTO origin.origin_data (sta,sta_tm, category, sn,sn_tm,v_up,v_down,err)")
}