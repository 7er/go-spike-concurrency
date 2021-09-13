package logger

import (
	"fmt"
	"sync"
)

type Command interface {
	Execute()
}

type logCommand struct {
	entry       string
	doneChannel chan bool
}

func (self *logCommand) Execute() {
	for _, c := range self.entry {
		fmt.Printf("%c", c)
	}
	fmt.Println("")
	self.doneChannel <- true
}

type Logger struct {
	inputChannel chan Command
	wg           sync.WaitGroup
}

func NewLogger() *Logger {
	input := make(chan Command)
	result := Logger{inputChannel: input}
	result.Start()
	return &result
}

func (self *Logger) serve() {
	for command := range self.inputChannel {
		command.Execute()
	}
	self.wg.Done()

}

func (self *Logger) Start() {
	self.wg.Add(1)
	go self.serve()
}

func (self *Logger) Close() {
	close(self.inputChannel)
	self.wg.Wait()

}

func (self *Logger) Log(entry string) {
	// send entry to log server
	ch := make(chan bool)
	command := logCommand{entry: entry, doneChannel: ch}
	self.inputChannel <- &command
	<-ch
}
