package main

import (
	"fmt"
	"net"
	"log"
	"os"
)

func main() {

	fileName := "xxx_debug.log"
	logFile,err  := os.Create(fileName)
	defer logFile.Close()
	if err != nil {
		log.Fatalln("open file error !")
	}
	//debugLog := log.New(logFile,"[Debug]",log.Llongfile)

	fmt.Printf("hello world！")
	listener, err := net.Listen("tcp", "127.0.0.1:50000")

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
		//create a goroutine for each request.
		go doServerStuff(conn)
	}
}

func doServerStuff(conn net.Conn) {
	log.Println("new connection:", conn.LocalAddr())
	for {
		buf := make([]byte, 1024)
		length, err := conn.Read(buf)
		if err != nil {
			log.Fatal("Error reading:", err.Error())
			return
		}

		log.Println("Receive data from client:", string(buf[:length]))
	}
}
