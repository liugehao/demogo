package lib

import (
	//"strings"
	//"fmt"
	"regexp"
	"strings"
	"errors"
	"log"
)

var RdRegSta *regexp.Regexp
var RdRegTm *regexp.Regexp
var RdRegBatt *regexp.Regexp
var RdRegVer *regexp.Regexp

var T01RegSn *regexp.Regexp
var T01RegVup *regexp.Regexp
var T01RegVdown *regexp.Regexp
var T01RegErr *regexp.Regexp

var NumberRex *regexp.Regexp

func Init() {
	RdRegSta = regexp.MustCompile("STA:(.*?);")
	RdRegTm = regexp.MustCompile("TM:(.*?);")
	RdRegBatt = regexp.MustCompile("BATT:([\\d.]+)")
	RdRegVer = regexp.MustCompile("VER:(.*?);")

	T01RegSn = regexp.MustCompile("SN:(.*?);")
	T01RegVup = regexp.MustCompile("V\\+:([\\d.]+)")
	T01RegVdown = regexp.MustCompile("V-:([\\d.]+)")
	T01RegErr = regexp.MustCompile("E:([\\d.]+)")

	NumberRex= regexp.MustCompile("([\\d.]+)")

}

func ParseRouteDevice(s string) (fl RouteDevice, err error) {

	fl.STA = rex(s, RdRegSta, 1)
	fl.TM, err = ParseDateTime(rex(s, RdRegTm, 1))
	fl.BATT = rex(s, RdRegBatt, 1)
	fl.VER = rex(s, RdRegVer, 1)

	return
}

func rex(s string, regexp2 *regexp.Regexp, index int) string {
	d := regexp2.FindStringSubmatch(s)
	return d[index]

}

func ParseT01(rd RouteDevice, s string) (t01 T01, err error) {
	t01.RouteDevice = rd
	t01.TM, err = ParseDateTime(rex(s, RdRegTm, 1))
	if err != nil {
		return
	}
	t01.Vdown = rex(s, T01RegVdown, 1)
	t01.Vup = rex(s, T01RegVup, 1)
	t01.SN = rex(s, T01RegSn, 1)
	t01.ERR = rex(s, T01RegErr, 1)

	return

}

func ParseDateTime(s string) (string, error) {
	if len(s) != 14 {
		return s, errors.New("时间格式不正确" + s)
	}
	return s[:4] + "-" + s[4:6] + "-" + s[6:8] + " " + s[8:10] + ":" + s[10:12] + ":" + s[12:], nil
}

/*func main() {
	str := "STA:334;TM:20160909090909;BATT:3.6V;VER:3.3;#T00:20160909090909;3.5mpa;#T01:TM:20160909090909;SN:232;V+:34L;V-:34L;E:00;#"

	for i := 0; i < 200000; i++ {
		Parse(str)
	}

}*/

func Parse(str string) {

	buf := strings.Split(str, "#")
	var fl RouteDevice
	var err error
	var t01 T01
	for _, s := range buf {
		if len(s) < 3 {

			break
		}
		switch s[:3] {
		case "STA":
			fl, err = ParseRouteDevice(s)
			fl.save()
		case "T01":
			t01, err = ParseT01(fl, s)
			t01.save()
		case "T00":
			tmpS := strings.Split(s[4:], ";")
			t00 := new(T00)
			t00.RouteDevice = fl
			t00.PRESSURE = NumberRex.FindString(tmpS[1])
			t00.TM, err = ParseDateTime(tmpS[0])
			if err != nil {
				log.Println(err.Error())
				return
			}
			t00.save()
		default:

		}

	}
}
