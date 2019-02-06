package main

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"strings"
)

// Server IRC Server
type Server struct {
	Users    []*User
	Channels []*Channel
}

func (sv *Server) getChan(chanName string) (*Channel, error) {
	for _, channel := range sv.Channels {
		if channel.Name == chanName {
			return channel, nil
		}
	}
	return nil, errors.New("No channel by that name")
}

func (sv *Server) getUser(username string) (*User, error) {
	for _, user := range sv.Users {
		if user.Nick == username {
			return user, nil
		}
	}
	return nil, errors.New("User doesn't exist")
}

//HandleConnection handle each client request
func (sv *Server) HandleConnection(c net.Conn) {
	fmt.Printf("Connected to %s\n", c.RemoteAddr().String())
	u := &User{Conn: c}
	for {
		data, err := bufio.NewReader(u.Conn).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		message := strings.TrimSpace(string(data))
		if message == "QUIT" {
			break
		}
		if message[0] == '/' {
			result := u.handleIRCCommand(message[1:])
			u.Conn.Write([]byte(result + "\n"))
		} else {
			result := sv.handleCommand(message)
			u.Conn.Write([]byte(result + "\n"))
		}
	}
	u.Conn.Close()
}
