package main

import "net"

type User struct {
	Name string
	Addr string
	Conn net.Conn
	Chan chan string
}

func NewUser(conn net.Conn) *User {
	addr := conn.RemoteAddr().String()
	_, port, _ := net.SplitHostPort(addr)

	user := &User{
		Name: port,
		Addr: addr,
		Conn: conn,
		Chan: make(chan string),
	}

	go user.HandleMsg()

	return user
}

func (user *User) HandleMsg() {
	for {
		msg := <-user.Chan
		user.Conn.Write([]byte(msg + "\n"))
	}
}
