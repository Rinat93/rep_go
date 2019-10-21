package main

import (
	"fmt"
	"net"
	"os"
	"reg_go/config"
	"reg_go/lib"
)

func main() {
	lib.Lib1()
	l, err := net.Listen(config.CONNTYPE, config.CONNHOST+":"+config.CONNPORT)
	if err != nil {
		fmt.Println("ERROR listening: ", err.Error())
		os.Exit(1)
	}
	defer l.Close()
	fmt.Println("Listening on " + config.CONNHOST + ":" + config.CONNPORT)
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
		}
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	buf := make([]byte, 2048)
	_, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	fmt.Println("INFORMATIONS: ", string(buf))
	conn.Write([]byte("Message received."))
	conn.Close()
}
