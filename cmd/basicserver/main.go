package main

import "github.com/tanveerprottoy/basic-go-server/internal/app/basicserver"

// entry point
func main() {
	a := basicserver.NewApp()
	a.Run()
}
