package server

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Usage: ./server <port>")
		return
	}
	port := ":" + args[1]
	launchServer(port)
}
