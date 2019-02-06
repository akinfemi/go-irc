package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"strings"

	"github.com/jroimartin/gocui"
)

func handleServerReply(g *gocui.Gui, reply string) error {
	var r Reply
	must(json.Unmarshal([]byte(reply), &r))
	switch r.command {
	case "lu":
		v, _ := g.View("users")
		return UpdateViewList(v, r.body)
	case "lc":
		v, _ := g.View("channels")
		return UpdateViewList(v, r.body)
	default:
		fmt.Printf("%s not implemented yet\n", r.command)
	}
	return nil
}

func runCommand(g *gocui.Gui, userInput string) {
	tokens := strings.Split(userInput, " ")
	if tokens[0] == "/connect" {
		client.mutex.Lock()
		if !client.connected {
			conn, err := net.Dial("tcp4", string(tokens[1]))
			must(err)
			client.connected = true
			client.conn = conn
		}
		client.mutex.Unlock()
	} else {
		fmt.Fprintf(client.conn, userInput)                         //send input to server
		reply, err := bufio.NewReader(client.conn).ReadString('\n') //response from server
		must(err)
		must(handleServerReply(g, reply))
	}
}

func readAndSend(g *gocui.Gui, v *gocui.View) error {
	userInput := strings.TrimSpace(v.Buffer())
	v.Clear()
	inputView, err := g.View("input")
	must(err)
	v.SetCursor(inputView.Origin())
	go runCommand(g, userInput)
	return nil
}
