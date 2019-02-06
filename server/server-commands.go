package main

import (
	"errors"
	"fmt"
	"strings"
)

func (sv *Server) listChannels() []string {
	result := []string{}
	fmt.Println("here", len(sv.Channels))
	for _, channel := range sv.Channels {
		result = append(result, channel.Name)
	}
	return result
}

func (sv *Server) listUsers() []string {
	result := []string{}
	for _, users := range sv.Users {
		result = append(result, users.Nick)
	}
	return result
}

func (sv *Server) addChannel(name string) {
	ch := &Channel{Name: name, Members: []*User{}, Messages: []*Message{}}
	sv.Channels = append(sv.Channels, ch)
}

func (sv *Server) addUser(name string) {
	u := &User{Nick: name, UserChannels: []*Channel{}, PrivateMessages: map[string][]Message{}}
	sv.Users = append(sv.Users, u)
}

func (sv *Server) handleCommand(message string) ([]string, error) {
	tokens := strings.Split(message, " ")
	switch tokens[0] {
	case "lc":
		return sv.listChannels(), nil
	case "lu":
		return sv.listUsers(), nil
	case "nc":
		if len(tokens) != 2 {
			break
		}
		sv.addChannel(tokens[1])
	case "nu":
		if len(tokens) != 2 {
			break
		}
		sv.addUser(tokens[1])
	}
	return []string{}, errors.New("Invalid Command + Args")
}
