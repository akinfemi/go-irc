package server

import "fmt"

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
