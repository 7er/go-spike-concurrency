package main

import (
	"go-spike-concurrency/logger"
	"go-spike-concurrency/command"
	"sync"
)

var wg sync.WaitGroup

func logIt(server *command.CommandServer, entry string) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			logger.Log(server, entry)
		}
	}()

}

func main() {
	server := command.NewCommandServer()
	// server.Log("flesk")
	logIt(server, "flesk")
	logIt(server, "bacon")
	logIt(server, "duppe")
	wg.Wait()
	server.Close()
}
