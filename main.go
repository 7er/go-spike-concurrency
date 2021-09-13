package main

import "go-spike-concurrency/logger"
import "fmt"

func main() {
	logger := logger.NewLogger()
	logger.Log("flesk")
	fmt.Println(logger.Content())
}
