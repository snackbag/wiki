package main

import "github.com/snackbag/compass/compass"

func main() {
	server := compass.NewServer()
	server.Start()
}
