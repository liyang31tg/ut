package logger

import (
	"log"
	"os"
	"github.com/fatih/color"
)

var Err *log.Logger
var Info *log.Logger

func init() {
	color.NoColor = false
	Err = log.New(os.Stderr,color.RedString("error "),log.LstdFlags|log.Lshortfile)
	Info = log.New(os.Stdout,"",log.LstdFlags|log.Lshortfile)
}
