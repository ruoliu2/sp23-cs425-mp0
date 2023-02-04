package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

var HOST = "127.0.0.1"
var PORT = "8080"
var TYPE = "tcp"

// adapted from https://www.golinuxcloud.com/golang-tcp-server-client
func main() {
	// ./node node1 10.0.0.1 1234
	// ./node node-name address port

	// parse arguments
	arguments := os.Args
	if len(arguments) != 4 {
		fmt.Println("Please provide 3 arguments")
		return
	}

	NAME := arguments[1]
	HOST := arguments[2]
	PORT := arguments[3]
	tcpServer, err := net.ResolveTCPAddr(TYPE, HOST+":"+PORT)
	if err != nil {
		println("ResolveTCPAddr failed:", err.Error())
		os.Exit(1)
	}
	conn, err := net.DialTCP(TYPE, nil, tcpServer)
	if err != nil {
		println("Dial failed:", err.Error())
		os.Exit(1)
	}
	fmt.Fprint(os.Stdout, "Connected to server")

	// send the name of client
	fmt.Fprintf(conn, "%s\n", NAME)

	// open stdin, while there's new line, get and send to server
	reader := bufio.NewReader(os.Stdin)
	for {
		// 1610688413.782391 ce783874ba65a148930de32704cd4c809d22a98359f7aed2c2085bc1bd10f096
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Read from stdin failed:", err.Error())
			os.Exit(1)
		}
		// 1610688413.782391 node1 ce783874ba65a148930de32704cd4c809d22a98359f7aed2c2085bc1bd10f096
		// insert NAME in between the time and event
		// write to server
		fmt.Fprintf(conn, "%s %s %s", text[:17], NAME, text[18:])
		fmt.Fprintf(os.Stdout, "Sending: %s %s %s", text[:17], NAME, text[18:])

		// read from server
		// message, err := bufio.NewReader(conn).ReadString('\n')
		// if err != nil {
		// 	fmt.Println("Read from server failed:", err.Error())
		// 	os.Exit(1)
		// }
		// fmt.Print("Message from server: " + message)
	}

	conn.Close()
}
