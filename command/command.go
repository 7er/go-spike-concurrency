package command

import (
	"sync"
)

type Command interface {
	Execute()
}


type CommandServer struct {
	inputChannel chan Command
	wg           sync.WaitGroup
}

func NewCommandServer() *CommandServer {
	input := make(chan Command)
	result := CommandServer{inputChannel: input}
	result.Start()
	return &result
}

func (self *CommandServer) serve() {
	for command := range self.inputChannel {
		command.Execute()
	}
	self.wg.Done()

}

func (self *CommandServer) Start() {
	self.wg.Add(1)
	go self.serve()
}

func (self *CommandServer) Close() {
	close(self.inputChannel)
	self.wg.Wait()

}

func (self *CommandServer) ScheduleCommand(command Command ) {
	self.inputChannel <- command
}

