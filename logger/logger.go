package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"time"
)

var HOST = "127.0.0.1"
var PORT = "8080"
var TYPE = "tcp"
var totalBandwidth = 0

// create map[conn] = name
var clientName = make(map[net.Conn]string)

// adapted from https://www.golinuxcloud.com/golang-tcp-server-client
func main() {
	//./logger 1234
	// parse arguments
	arguments := os.Args
	if len(arguments) != 2 {
		fmt.Println("Please provide 1 argument")
		return
	}

	PORT := arguments[1]

	// FIXME: better way?
	fmt.Fprintf(os.Stdout, "%.6f - logger started\n", float64(time.Now().UnixMicro())/1e6)

	//1610688413.743385 - node1 connected
	//1610688413.782391 node1 ce783874ba65a148930de32704cd4c809d22a98359f7aed2c2085bc1bd10f096
	listen, err := net.Listen(TYPE, HOST+":"+PORT)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	// close the listener after the server terminates
	defer listen.Close()

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	// parse the name of the client
	buffer, err := bufio.NewReader(conn).ReadBytes('\n')
	if err != nil {
		// client disconnected, print error
		fmt.Println(err.Error())
		return
	}
	// 1610688413.743385 - node1 connected
	// log to stdout seconds since 1970 - name of the client
	fmt.Fprintf(os.Stdout, "%.6f - %s connected\n",
		float64(time.Now().UnixMicro())/1e6, string(buffer[:len(buffer)-1]))
	clientName[conn] = string(buffer[:])

	// parse each line of the msg received
	for {
		// read each line sent to server
		buffer, err := bufio.NewReader(conn).ReadBytes('\n')
		if err != nil {
			// client disconnected, print error
			fmt.Println(err.Error())
			return
		}

		//1610688413.782391 node1 ce783874ba65a148930de32704cd4c809d22a98359f7aed2c2085bc1bd10f096
		// log to stdout the msg received
		clientTime, err := strconv.ParseFloat(string(buffer[:18]), 64)
		delay := float64(time.Now().UnixMicro())/1e6 - clientTime
		totalBandwidth += len(buffer)

		fmt.Fprintf(os.Stdout, "%s", string(buffer[:]))
	}

	//responseStr := fmt.Sprintf("Your message is: %v. Received curTime: %v", string(buffer[:]), curTime)
	//conn.Write([]byte(responseStr))

	// close conn
	conn.Close()
}
