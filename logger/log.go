package logger

import (
	"log"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/robfig/cron"
)

var Err = log.New(os.Stderr, color.RedString("error "), log.LstdFlags|log.Lshortfile)
var Info = log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)

func init() {
	os.Mkdir("logs", 0777)
	color.NoColor = false
	setOutput(Err)
	setOutput(Info)
	c := cron.New()
	c.AddFunc("0 0 0 * * *", func() {
		setOutput(Err)
		setOutput(Info)
	})
	c.Start()

}
func setOutput(l *log.Logger) {
	logName := time.Now().Format("20060102") + ".log"
	file, err := os.OpenFile("./logs/"+logName, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	//defer file.Close(	)
	if err != nil {
		log.Fatalln("fail to create " + logName + ".log file!")
	}
	l.SetOutput(file)
}
