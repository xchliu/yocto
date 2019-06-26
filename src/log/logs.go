package log

import (
	// "io"
	"fmt"
	"log"
	"os"
)

var (
	Trace   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

func LogInit(logpath string) {
	/*	traceHandle io.Writer,
		infoHandle io.Writer,
		warningHandle io.Writer,
		errorHandle io.Writer) */
	logfile, err := os.OpenFile(logpath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
	Trace = log.New(logfile, "TRACE: ", log.Ldate|log.Ltime|log.Lshortfile)
	Info = log.New(logfile, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	Warning = log.New(logfile, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(logfile, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}
