package main

import (
	"fmt"
	"strings"
)

//Channel channels in the IRC
type Channel struct {
	Name     string
	Members  []*User
	Messages []*Message
}

func (ch Channel) who() {
	for _, user := range ch.Members {
		fmt.Println(user.Nick)
	}
}

func (ch Channel) String() string {
	members := []string{}
	for _, member := range ch.Members {
		members = append(members, member.Nick)
	}
	return strings.Join(members, " ")
}
