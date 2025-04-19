package logger

import (
	"fmt"
	"io"
)

type GoLogger struct {
	OutputSteam  io.Writer
	DateAndTime  bool
	FunctionName bool
}

func NewGoLogger() {

}

func (l *GoLogger) Print(a ...any) {
	fmt.Print(a...)
}

func (l *GoLogger) Printf() {

}

func (l *GoLogger) Println() {

}

func (l *GoLogger) WriteToStream(buffer *[]byte) {
	l.OutputSteam.Write(*buffer)
}
