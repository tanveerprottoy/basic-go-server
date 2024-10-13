package main

import "github.com/tanveerprottoy/basic-go-server/internal/server"

// entry point
func main() {
	a := server.NewApp()
	a.Run()
}
