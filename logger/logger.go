package logger

type Command interface {
	Execute()
}

type logCommand struct {
	entry string
	doneChannel chan bool
}



type Logger struct {
	inputChannel chan * Command
}

func NewLogger() *Logger {
	input := make(chan * Command)
	result := Logger{inputChannel: input}
	result.Start()
	return &result
}

func (self *Logger) serve() {
}

func (self *Logger) Start() {
	go self.serve()
}

func (self *Logger) Log(entry string) {
	// send entry to log server
	ch := make(chan bool)
	command := logCommand{entry: entry, doneChannel: ch}
	self.inputChannel <- &command
	done := <-ch
}

func (self *Logger) Content() string {
	// suck content from channel
	return "flusk"
}
