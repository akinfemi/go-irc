package main

import (
	"fmt"
	"net"
	"os"

	"github.com/akinfemi/go-irc/server"
)

func launchServer(port string) {
	conn, err := net.Listen("tcp4", port) //support (for now ipv4)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	sv := &server.Server{Users: []*server.User{}, Channels: []*server.Channel{}}
	for { //accept connections
		fmt.Println("Waiting for connections...")
		c, err := conn.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go sv.HandleConnection(c) //multi-client (non-blocking)
	}
}

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Usage: ./server <port>")
		return
	}
	port := ":" + args[1]
	launchServer(port)
}
