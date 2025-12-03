package main

import (
	"bufio"
	"fmt"
	"net"
	"sync"
)

type Client struct {
	conn     net.Conn
	username string
}

type Server struct {
	clients map[*Client]bool
	lock    sync.Mutex
}

func NewServer() *Server {
	s := &Server{
		clients: make(map[*Client]bool),
	}

	return s
}

func (s *Server) addClient(c *Client) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.clients[c] = true
}

func (s *Server) removeClient(c *Client) {
	s.lock.Lock()
	defer s.lock.Unlock()

	delete(s.clients, c)
}

func (s *Server) broadcastMessage(msg string, sender *Client) {
	s.lock.Lock()
	defer s.lock.Unlock()

	for c := range s.clients {
		if c == sender {
			continue
		}

		_, err := c.conn.Write([]byte(msg))
		if err != nil {
			fmt.Println("send msg failed:", err)
		}
	}
}

func handleConn(c net.Conn, s *Server) {
	defer c.Close()

	_, p, _ := net.SplitHostPort(c.RemoteAddr().String())
	fmt.Println("new connect is ", p)

	client := &Client{
		conn:     c,
		username: p,
	}

	s.addClient(client)

	go s.broadcastMessage("welcome user "+p+"\n", client)

	// f1 : scanner
	scanner := bufio.NewScanner(c)
	for scanner.Scan() {
		msg := scanner.Text()
		msg = fmt.Sprintf("%s : %s\n", client.username, msg)
		go s.broadcastMessage(msg, client)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("client err: ", client.username)
		s.removeClient(client)
	}

	//  f2, bufio reader
	/*reader := bufio.NewReader(c)
	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("client err: ", client.username)
			s.removeClient(client)
			break
		}

		msg = fmt.Sprintf("%s : %s", client.username, msg)
		go s.broadcastMessage(msg, client)
	}*/

	//  f3, conn.read
	/*buf := make([]byte, 1024)
	for {
		len, err := c.Read(buf)

		if err != nil {
			fmt.Println("client err: ", client.username)
			s.removeClient(client)
			break
		}

		msg := fmt.Sprintf("%s : %s\n", client.username, string(buf[:len]))
		go s.broadcastMessage(msg, client)
	}*/
}

func main() {
	fmt.Println("server gogogo")
	server := NewServer()

	listener, err := net.Listen("tcp", "localhost:1024")
	if err != nil {
		fmt.Println("startup server error:", err)
		return
	}

	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("connect to server error:", err)
			continue
		}

		go handleConn(conn, server)
	}
}
