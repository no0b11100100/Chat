package logger

import (
	"fmt"
	"log"
	"os"
)

type Writter struct{}

func (w *Writter) Write(p []byte) (int, error) {
	p = append(p, []byte("\033[0m")...)
	return os.Stdout.Write(p)
}

var flags = log.Lshortfile | log.Lmsgprefix | log.Ltime

var writter *Writter = &Writter{}

var (
	Info    = log.New(writter, "", flags)
	Warning = log.New(writter, "", flags)
	Error   = log.New(writter, "", flags)
)

const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorYellow = "\033[33m"
)

func InitLogger(prefix string) {
	Info.SetPrefix(fmt.Sprintf("%v%v %v: ", colorReset, "[INFO]", prefix))
	Warning.SetPrefix(fmt.Sprintf("%v%v %v: ", colorYellow, "[WARNING]", prefix))
	Error.SetPrefix(fmt.Sprintf("%v%v %v: ", colorRed, "[ERROR]", prefix))
}
