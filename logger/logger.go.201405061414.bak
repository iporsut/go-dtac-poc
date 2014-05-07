package logger

import (
	"io"
	//"io/ioutil"
	"log"
	//"os"
)

type Logger struct {
	TRACE   *log.Logger
	INFO    *log.Logger
	WARNING *log.Logger
	ERROR   *log.Logger
}

func (logger *Logger) Init(
	traceHandle io.Writer,
	infoHandle io.Writer,
	warningHandle io.Writer,
	errorHandle io.Writer) {

	logger.TRACE = log.New(traceHandle,
		"TRACE: ",
		log.Ldate|log.Lmicroseconds|log.Lshortfile)

	logger.INFO = log.New(infoHandle,
		"INFO: ",
		log.Ldate|log.Lmicroseconds|log.Lshortfile)

	logger.WARNING = log.New(warningHandle,
		"WARNING: ",
		log.Ldate|log.Lmicroseconds|log.Lshortfile)

	logger.ERROR = log.New(errorHandle,
		"ERROR: ",
		log.Ldate|log.Lmicroseconds|log.Lshortfile)
}

//func main() {

//	file, err := os.OpenFile("./file.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
//	if err != nil {
//		log.Fatalln("Failed to open log file", os.Stdout, ":", err)
//	}

//	multi := io.MultiWriter(file, os.Stdout)

//	var logger Logger

//	logger.Init(ioutil.Discard, os.Stdout, multi, os.Stderr) // multi/file

//	logger.TRACE.Println("I have something standard to say")
//	logger.INFO.Println("Special Information")
//	logger.WARNING.Println("There is something you need to know about")
//	logger.ERROR.Println("Something has failed")

//}
