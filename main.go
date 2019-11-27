package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"reg_go/config"
)

func main() {
	//lib.Lib1()
	l, err := net.Listen(config.CONNTYPE, config.CONNHOST+":"+config.CONNPORT)
	if err != nil {
		fmt.Println("ERROR listening: ", err.Error())
		os.Exit(1)
	}
	defer l.Close()
	// Новое соединение
	newConnections := make(chan net.Conn)
	// Новое сообщение
	mess := make(chan string)
	fmt.Println("Listening on " + config.CONNHOST + ":" + config.CONNPORT)

	go func() {
		for {
			conn, err := l.Accept()
			if err != nil {
				fmt.Println("Error accepting: ", err.Error())
			}
			newConnections <- conn
			go handleRequest(conn,mess)

		}
	}()
	// Ожидаем событий
	for {
		select {
			case conn := <-newConnections:
				log.Printf("Accepted new client")
				go handleRequest(conn,mess)
			case mess := <-mess:
				log.Printf(mess)
		}
	}

}

func handleRequest(conn net.Conn, mess chan<- string) {
	buf := make([]byte, 2048)
	_, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	mess <- string(buf)
	conn.Write([]byte("Message received."))
	//conn.Close()
}
