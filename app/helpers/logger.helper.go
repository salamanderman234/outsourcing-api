package helpers

import (
	"fmt"
	"log"
)

type loggerHelper struct{}

// func (loggerHelper) writeLog(msg string) {
// 	f, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
// 	if err != nil {
// 		log.Fatalf("error opening file: %v", err)
// 	}
// 	log.SetOutput(f)
// 	defer f.Close()
// }

func (loggerHelper) log(tipe string, msg string, colorTag ...string) {
	color := "\033[0m"
	if len(colorTag) == 1 {
		color = colorTag[0]
	}
	message := fmt.Sprintf("%s[%s] \033[0m%s", color, tipe, msg)
	log.Print(message)
}

func (l loggerHelper) Info(msg string) {
	l.log("INFO", msg, "\033[36m")
}
func (l loggerHelper) Warning(msg string) {
	l.log("WARNING", msg, "\033[0;33m")
}
func (l loggerHelper) Error(msg string) {
	l.log("ERROR", msg, "\033[0;31m")
}
func (l loggerHelper) Fatal(msg string) {
	l.log("FATAL", msg, "\033[0;34m")
}

var Logger = loggerHelper{}
