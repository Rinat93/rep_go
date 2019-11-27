package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"net"
	"os"
	"reg_go/config"
	"strconv"
)

func Eval(exp ast.Expr) int {
	switch exp := exp.(type) {
	case *ast.BinaryExpr:
		return EvalBinaryExpr(exp)
	case *ast.BasicLit:
		switch exp.Kind {
		case token.INT:
			i, _ := strconv.Atoi(exp.Value)
			return i
		}
	}

	return 0
}

func EvalBinaryExpr(exp *ast.BinaryExpr) int {
	left := Eval(exp.X)
	right := Eval(exp.Y)

	switch exp.Op {
		case token.ADD:
			return left + right
		case token.SUB:
			return left - right
		case token.MUL:
			return left * right
		case token.QUO:
			return left / right
	}

	return 0
}

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
	// Все пользователи
	allClients := make(map[net.Conn]int)
	fmt.Println("Listening on " + config.CONNHOST + ":" + config.CONNPORT)

	go func() {
		for {
			conn, err := l.Accept()
			if err != nil {
				fmt.Println("Error accepting: ", err.Error())
			}
			newConnections <- conn
			allClients[conn] = 1

		}
	}()

	// Ожидаем событий
	for {
		select {
			case conn := <-newConnections:
				log.Printf("Accepted new client")
				go handleRequest(conn,mess)
			case mess := <-mess:
				exp, err := parser.ParseExpr(mess)
				if err != nil {
					fmt.Printf("parsing failed: %s\n", err)
					continue
				}
				res := Eval(exp)
				fmt.Printf("%d\n", res)
				for user,_ := range allClients {
					user.Write([]byte(fmt.Sprintf("%s \n",strconv.Itoa(res))))
				}
		}
	}

}

func handleRequest(conn net.Conn, mess chan<- string) {

	for {
		buf := make([]byte, 2048)
		_, err := conn.Read(buf)
		buf = bytes.Trim(buf, "\x00")
		if err != nil{
			fmt.Println("Error reading:", err.Error())
		}
		mess <- string(buf)
	}

	//conn.Close()
}
