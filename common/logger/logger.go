package logger

import (
	"fmt"
	"log"
	"os"
)

var flags = log.Lshortfile | log.Lmsgprefix | log.Ltime

const (
	info     = "[INFO]"
	warning  = "[WARNING]"
	error    = "[ERROR]"
	critical = "[CRITICAL]"
)

var (
	Info     = log.New(os.Stdout, "", flags)
	Warning  = log.New(os.Stdout, "", flags)
	Error    = log.New(os.Stdout, "", flags)
	Critical = log.New(os.Stdout, "", flags)
)

const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorYellow = "\033[33m"
)

func InitLogger(prefix string) {
	Info.SetPrefix(fmt.Sprintf("%v%v %v: ", colorReset, info, prefix))
	Warning.SetPrefix(fmt.Sprintf("%v%v %v: ", colorYellow, warning, prefix))
	Error.SetPrefix(fmt.Sprintf("%v%v %v: ", colorRed, error, prefix))
	Critical.SetPrefix(fmt.Sprintf("%v%v %v: ", colorRed, critical, prefix))
}
