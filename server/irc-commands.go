package main

import (
	"fmt"
	"strings"
)

func (u *User) handleIRCCommand(message string) string {
	tokens := strings.Split(message, " ")
	switch tokens[0] {
	case "nick":
		usr, err := u.Sv.getUser(tokens[1])
		if err != nil {
			return err.Error()
		}
		u = usr
		return "User succesfully set"
	case "join":
		ch, err := u.Sv.getChan(tokens[1])
		if err != nil {
			return err.Error()
		}
		u.joinChan(ch)
		return fmt.Sprintf("%s joined %s", u.Nick, ch.Name)
	case "who":
		ch, err := u.Sv.getChan(tokens[1])
		if err != nil {
			return err.Error()
		}
		return fmt.Sprintf("%s\n", ch)
	case "msg":
	case "leave":
	}
	return ""
}
