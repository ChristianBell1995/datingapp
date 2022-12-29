package main

import (
	"github.com/ChristianBell1995/datingapp/config"
)

func main() {
	server := config.NewServer()

	server.Router.Run(":8000")
}
