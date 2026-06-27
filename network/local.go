package network

import (
	"bufio"
	"fmt"
	"net"
)

func StartHost(port string) net.Conn {
	listener, _ := net.Listen("tcp", port)
	fmt.Println("Waiting for connection on", port)
	conn, _ := listener.Accept()
	fmt.Println("Player connected!")
	return conn
}

func Connect(ip string) net.Conn {
	conn, _ := net.Dial("tcp", ip)
	return conn
}

func Send(conn net.Conn, msg string) {
	fmt.Fprintln(conn, msg)
}

func Receive(conn net.Conn) string {
	scanner := bufio.NewScanner(conn)
	scanner.Scan()
	return scanner.Text()
}
