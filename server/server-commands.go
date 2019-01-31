package server

import (
	"fmt"
	"strings"
)

func listChannels(sv *Server) {
	for _, channel := range sv.Channels {
		fmt.Println(channel.Name)
	}
}

func listUsers(sv *Server) {
	for _, users := range sv.Users {
		fmt.Println(users.Nick)
	}
}

func (sv *Server) newChannel(name string) {
	ch := &Channel{Name: name, Members: []*User{}, Messages: []*Message{}}
	sv.Channels = append(sv.Channels, ch)
}

func (sv *Server) newUser(name string) {
	u := &User{Nick: name, UserChannels: []*Channel{}, PrivateMessages: map[string][]Message{}}
	sv.Users = append(sv.Users, u)
}

func (sv *Server) handleCommand(message string) {
	switch message[:2] {
	case "lc":
		listChannels(sv)
	case "ls":
		listUsers(sv)
	case "nc":
		tokens := strings.Split(message, " ")
		if len(tokens) != 2 {
			fmt.Println("Invalid number of arguments")
			return
		}
		sv.newChannel(tokens[1])
	case "nu":
		tokens := strings.Split(message, " ")
		if len(tokens) != 2 {
			fmt.Println("Invalid number of arguments")
			return
		}
		sv.newUser(tokens[1])
	}
}
