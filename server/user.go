package main

import (
	"fmt"
	"net"
)

//User each user (connection)
type User struct {
	Nick            string
	Sv              *Server
	Conn            net.Conn
	UserChannels    []*Channel
	PrivateMessages map[string][]Message
}

func (u User) String() string {
	return fmt.Sprintf("Name: %s\n", u.Nick)
}

func (u *User) leaveChan(ch *Channel) {
	i := -1
	for idx, channel := range u.UserChannels {
		if channel.Name == ch.Name {
			i = idx
			break
		}
	}
	if i != -1 {
		u.UserChannels = append(u.UserChannels[:i], u.UserChannels[i+1:]...)
		i = -1
	}
	for idx, user := range ch.Members {
		if user.Nick == u.Nick {
			i = idx
			break
		}
	}
	if i != -1 {
		ch.Members = append(ch.Members[:i], ch.Members[i+1:]...)
	}
}

func (u *User) joinChan(ch *Channel) {
	for _, channel := range u.UserChannels {
		if channel.Name == ch.Name {
			return
		}
	}
	ch.Members = append(ch.Members, u)
	u.UserChannels = append(u.UserChannels, ch)
}

func (u *User) invite(usr *User, ch *Channel) {
	usr.joinChan(ch)
}

func (u *User) sendMessage(msg Message, dest interface{}) {
	switch v := dest.(type) {
	case Channel:
		v.Messages = append(v.Messages, &msg)
	case User:
		if _, ok := v.PrivateMessages[v.Nick]; !ok {
			v.PrivateMessages[v.Nick] = []Message{}
		}
		v.PrivateMessages[u.Nick] = append(v.PrivateMessages[v.Nick], msg)
	default:
		fmt.Printf("Random type")
	}
}
