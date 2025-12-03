package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	fmt.Println("client gogogo")

	conn, err := net.Dial("tcp", "localhost:8888")
	if err != nil {
		fmt.Println("connect to server failed:", err)
		return
	}

	defer conn.Close()

	go func() {
		reader := bufio.NewReader(conn)
		for {
			msg, _ := reader.ReadString('\n')
			fmt.Printf(msg)
		}
	}()

	for {
		reader := bufio.NewReader(os.Stdin)
		msg, _ := reader.ReadString('\n')
		msg = strings.TrimSpace(msg)

		conn.Write([]byte(msg + "\n"))
	}
}
