package server

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

// Server IRC Server
type Server struct {
	Users    []*User
	Channels []*Channel
}

func (sv *Server) handleMessage(message string) {
	if message[0] == '/' {
		sv.handleIRCCommand(message[1:])
	} else {
		sv.handleCommand(message)
	}
}

func (sv *Server) handleConnection(c net.Conn) {
	fmt.Printf("Connected to %s\n", c.RemoteAddr().String())
	for {
		data, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		message := strings.TrimSpace(string(data))
		if message == "QUIT" {
			break
		}
		sv.handleMessage(message)
		c.Write([]byte("From server: " + message + "\n"))
	}
	c.Close()
}

func launchServer(port string) {
	conn, err := net.Listen("tcp4", port) //support (for now ipv4)
	if err != nil {
		fmt.Println(err)
		return
	}
	sv := &Server{Users: []*User{}, Channels: []*Channel{}}
	defer conn.Close()
	for { //accept connections
		fmt.Println("Waiting for connections...")
		c, err := conn.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go sv.handleConnection(c) //multi-client (non-blocking)
	}
}
