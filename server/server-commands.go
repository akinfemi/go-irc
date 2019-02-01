package server

import (
	"fmt"
	"strings"
)

func listChannels(sv *Server) string {
	result := []string{}
	fmt.Println("here", len(sv.Channels))
	for _, channel := range sv.Channels {
		result = append(result, channel.Name)
	}
	if len(result) > 0 {
		return strings.Join(result, "\n")
	}
	return ""
}

func listUsers(sv *Server) string {
	result := []string{}
	for _, users := range sv.Users {
		result = append(result, users.Nick)
	}
	if len(result) > 0 {
		return strings.Join(result, "\n")
	}
	return ""
}

func (sv *Server) addChannel(name string) {
	ch := &Channel{Name: name, Members: []*User{}, Messages: []*Message{}}
	sv.Channels = append(sv.Channels, ch)
}

func (sv *Server) addUser(name string) {
	u := &User{Nick: name, UserChannels: []*Channel{}, PrivateMessages: map[string][]Message{}}
	sv.Users = append(sv.Users, u)
}

func (sv *Server) handleCommand(message string) string {
	tokens := strings.Split(message, " ")
	switch tokens[0] {
	case "lc":
		return listChannels(sv)
	case "lu":
		return listUsers(sv)
	case "nc":
		if len(tokens) != 2 {
			return "Invalid number of arguments"
		}
		sv.addChannel(tokens[1])
	case "nu":
		if len(tokens) != 2 {
			return "Invalid number of arguments"
		}
		sv.addUser(tokens[1])
	default:
		return "Server doesn't understand the command"
	}
	return "Success"
}
