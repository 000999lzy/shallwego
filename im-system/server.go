package main

import (
	"bufio"
	"fmt"
	"net"
	"sync"
)

type Server struct {
	UserMap map[string]*User
	lock    sync.RWMutex
}

func NewServer() *Server {
	server := &Server{
		UserMap: make(map[string]*User),
	}

	return server
}

func (s *Server) OnlineUser(u *User) {
	s.lock.Lock()
	s.UserMap[u.Name] = u
	s.lock.Unlock()

	msg := fmt.Sprintf("%s online\n", u.Name)
	s.Broadcast(msg, u)
}

func (s *Server) OfflineUser(u *User) {
	s.lock.Lock()
	delete(s.UserMap, u.Name)
	s.lock.Unlock()

	msg := fmt.Sprintf("%s offline\n", u.Name)
	s.Broadcast(msg, u)
}

func (s *Server) Startup() {
	fmt.Println("server gogogo")

	listen, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println("the server startup failed")
		return
	}

	defer listen.Close()

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("connect failed")
			continue
		}

		go s.HandleConn(conn)
	}
}

func (s *Server) HandleConn(c net.Conn) {
	defer c.Close()

	user := NewUser(c)
	s.OnlineUser(user)
	fmt.Println("new connect is ", user.Name)

	scanner := bufio.NewScanner(c)
	for scanner.Scan() {
		msg := scanner.Text()
		msg = fmt.Sprintf("%s: %s\n", user.Name, msg)
		user.Chan <- msg
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("user read err:", user.Name)
		s.OfflineUser(user)
	}

}

func (s *Server) Broadcast(msg string, sender *User) {
	for _, u := range s.UserMap {
		if u.Name == sender.Name {
			continue
		}

		u.Chan <- msg
	}
}
