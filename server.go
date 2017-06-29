package main

import (
	"net"
	"log"
	"serv1/lib"
	"fmt"
)

func main() {
	lib.Init()
	listener, err := net.Listen("tcp", "127.0.0.1:50001")

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
    s:=""
	for {
		buf := make([]byte, 1024*512)
		length, err := conn.Read(buf)
		if err != nil {
			//log.Fatal("Error reading:", err.Error())
			//fmt.Println(buf[:length])
			s += string( buf[:length])

			break
		}
		s += string( buf[:length])
	}
	lib.Parse(s)

}
