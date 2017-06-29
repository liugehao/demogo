package main

import (
	"net"
	"log"
	"serv1/lib"
	_ "github.com/lib/pq"
	"database/sql"
	"github.com/c4pt0r/ini"
	"fmt"
)

func main() {
	lib.Init()

	var conf = ini.NewConf("./conf.ini")
	host := conf.String("db", "host", "localhost")
	dbname := conf.String("db", "name", "dns")
	user := conf.String("db", "user", "dns")
	password := conf.String("db", "password", "123456")
	dbport := conf.String("db", "port", "5432")
	tcpPort := conf.String("server", "port", "50001")
	tcpAddr := conf.String("server", "ip", "0.0.0.0")
	conf.Parse()
	dataSourceName := "host=" + *host + " port=" + *dbport + " user=" + *user + " password=" + *password + " dbname=" + *dbname + " sslmode=disable"

	var er error;
	lib.DB, er = sql.Open("postgres", dataSourceName)
	if er != nil {
		log.Fatal("数据库连接失败")
	}
	var s string
	row:=lib.DB.QueryRow("SELECT CURRENT_TIMESTAMP ")
	row.Scan(&s)
	fmt.Println(s)
fmt.Printf("-----------")
	listener, err := net.Listen("tcp", *tcpAddr + ":" + *tcpPort)

	if err != nil {
		log.Println("Error listening:", err.Error())
		return
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error accepting:", err.Error())
			return
		}
		go doServerStuff(conn)
	}
}

func doServerStuff(conn net.Conn) {
	log.Println("new connection:", conn.LocalAddr())
	s := ""
	for {
		buf := make([]byte, 1024*512)
		length, err := conn.Read(buf)
		if err != nil {
			//log.Fatal("Error reading:", err.Error())
			fmt.Println(buf[:length])
			s += string(buf[:length])

			break
		}
		s += string(buf[:length])
	}
	lib.Parse(s)

}
