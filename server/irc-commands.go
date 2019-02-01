package server

import (
	"fmt"
	"strings"
)

func (u *User) handleIRCCommand(message string) {
	tokens := strings.Split(message, " ")
	switch tokens[0] {
	case "nick":
		usr, err := u.Sv.getUser(tokens[1])
		if err != nil {
			fmt.Println(err)
		}
		u = usr
	case "join":
		ch, err := u.Sv.getChan(tokens[1])
		if err != nil {
			fmt.Println(err)
		}
		u.joinChan(ch)
	case "who":
		ch, err := u.Sv.getChan(tokens[1])
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(ch)
	case "msg":
	case "leave":
	}
}
