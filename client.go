package main

import (
	"net"
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main()  {
	//open connection:
	conn, err := net.Dial("tcp", "127.0.0.1:50001")
	if err != nil {
		fmt.Println("Error dial:", err.Error())
		return
	}


	inputReader := bufio.NewReader(os.Stdin)

	//send info to server until Quit


	for {
		fmt.Println("What do you send to the server? Type Q to quit.")
		content, _ := inputReader.ReadString('\n')
		inputContent := strings.Trim(content, "\n")
		if inputContent == "Q" {
			return
		}

		_, err := conn.Write([]byte("STA:334;TM:20160909090909;BATT:3.6V;VER:3.3;#T00:20160909090909;3.5mpa;#T01:TM:20160909090909;SN:232;V+:34L;V-:34L;E:00;#"))
		if err != nil {
			fmt.Println("Error Write:", err.Error())
			return
		}
	}
}

func main1()  {

	conn, err := net.Dial("tcp", "127.0.0.1:50001")
	if err != nil {
		fmt.Println("Error dial:", err.Error())
		return
	}


	inputReader := bufio.NewReader(os.Stdin)

	//send info to server until Quit


	for {
		fmt.Println("What do you send to the server? Type Q to quit.")
		content, _ := inputReader.ReadString('\n')
		inputContent := strings.Trim(content, "\n")
		if inputContent == "Q" {
			return
		}

		_, err := conn.Write([]byte("STA:334;TM:20160909090909;BATT:3.6V;VER:3.3;#T00:20160909090909;3.5mpa;#T01:TM:20160909090909;SN:232;V+:34L;V-:34L;E:00;#"))
		if err != nil {
			fmt.Println("Error Write:", err.Error())
			return
		}
	}
}