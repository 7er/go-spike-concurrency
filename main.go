package main

import (
	"go-spike-concurrency/logger"
	"sync"
)

var wg sync.WaitGroup

func logIt(logger *logger.Logger, entry string) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			logger.Log(entry)
		}
	}()

}

func main() {
	logger := logger.NewLogger()
	// logger.Log("flesk")
	logIt(logger, "flesk")
	logIt(logger, "bacon")
	logIt(logger, "duppe")
	wg.Wait()
	logger.Close()
	// fmt.Println(logger.Content())
}
