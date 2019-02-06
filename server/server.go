package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net"
	"strings"
)

const (
	msgMode  = 1
	chanMode = 2
	neutMode = 0
)

//Reply : Message from server
type Reply struct {
	command string
	state   int
	body    []string
}

// Server IRC Server
type Server struct {
	state    int
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
			reply := Reply{command: message, state: sv.state, body: strings.Split(result, "\n")}
			b, err := json.Marshal(reply)
			must(err)
			u.Conn.Write(b)
		} else {
			result := []string{}
			result, err = sv.handleCommand(message)
			must(err)
			reply := Reply{command: message, state: sv.state, body: result}
			b, err := json.Marshal(reply)
			must(err)
			u.Conn.Write(b)
		}
	}
	u.Conn.Close()
}

func must(err error) {
	if err != nil {
		log.Panicln(err)
	}
}
