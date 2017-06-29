package main

import (
	"net"
	"fmt"
	"bufio"
	"os"
	"strings"
	"mydb"
)

func main()  {
	//open connection:
	conn, err := net.Dial("tcp", "127.0.0.1:50000")
	if err != nil {
		fmt.Println("Error dial:", err.Error())
		return
	}


	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please input your name:")
	clientName, _ := inputReader.ReadString('\n')
	inputClientName := strings.Trim(clientName, "\n")

	//send info to server until Quit
	for {
		fmt.Println("What do you send to the server? Type Q to quit.")
		content, _ := inputReader.ReadString('\n')
		inputContent := strings.Trim(content, "\n")
		if inputContent == "Q" {
			return
		}

		_, err := conn.Write([]byte(inputClientName + " says " + inputContent))
		if err != nil {
			fmt.Println("Error Write:", err.Error())
			return
		}
	}
}
