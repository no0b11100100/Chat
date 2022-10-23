package logger

import (
	"log"
	"os"
)

type Writter struct {
	prefix string
}

func (w *Writter) Write(p []byte) (int, error) {
	bytes := []byte(w.prefix)
	bytes = append(bytes, p...)
	bytes = append(bytes, []byte("\033[0m")...)
	return os.Stdout.Write(bytes)
}

var flags = log.Lshortfile | log.Lmsgprefix | log.Ltime

var (
	Info    = log.New(&Writter{"\033[0m"}, "[INFO]: ", flags)
	Warning = log.New(&Writter{"\033[31m"}, "[WARNING]: ", flags)
	Error   = log.New(&Writter{"\033[33m"}, "[ERROR]: ", flags)
)
