package main

import (
	"golang-app/server"
	"time"
)

func main() {
	time.Sleep(5 * time.Second)
	server.NewServer()
}
