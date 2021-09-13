package logger

import (
	"fmt"
	"go-spike-concurrency/command"
)


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

func Log(server *command.CommandServer, entry string) {
	// send entry to log server
	ch := make(chan bool)
	command := logCommand{entry: entry, doneChannel: ch}
	server.ScheduleCommand(&command)
	<-ch
}
