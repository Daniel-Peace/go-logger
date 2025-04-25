package logger

import (
	"fmt"
	"io"
)

type GoLogger struct {
	outputSteam  io.Writer
	dateAndTime  bool
	fileName     bool
	functionName bool
	lineNumber   bool
}

var buffer []byte

// func appendToBuffer(v ...any) {
// 	buffer = append(buffer, v...)
// }

func NewGoLogger(outputStream io.Writer, dateAndTime bool, fileName bool, functionName bool, lineNumber bool) *GoLogger {
	return &GoLogger{
		outputSteam:  outputStream,
		dateAndTime:  dateAndTime,
		fileName:     fileName,
		functionName: functionName,
		lineNumber:   lineNumber,
	}
}

func (l *GoLogger) generateLog(buffer *[]byte) {
	l.formatHeader(buffer)
}

func (l *GoLogger) formatHeader(buffer *[]byte) {
	showDateAndTime := l.dateAndTime
	showFileName := l.fileName
	showFunctionName := l.functionName
	showLineNumber := l.lineNumber

	if showDateAndTime {
		*buffer = append(*buffer, "[date and time]"...)
	}

	if showFileName {
		*buffer = append(*buffer, "[file name]"...)
	}

	if showFunctionName {
		*buffer = append(*buffer, "[function name]"...)
	}

	if showLineNumber {
		*buffer = append(*buffer, "[line number]"...)
	}
}

func (l *GoLogger) Print(a ...any) {
	var buffer []byte
	l.formatHeader(&buffer)
	fmt.Append(buffer, a...)
	buffer = fmt.Append(buffer, a...)
	l.WriteToStream(&buffer)
}

func (l *GoLogger) Printf(format string, a ...any) {
	var buffer []byte
	l.formatHeader(&buffer)
	fmt.Append(buffer, a...)
	buffer = fmt.Appendf(buffer, format, a...)
	l.WriteToStream(&buffer)
}

func (l *GoLogger) Println(a ...any) {
	var buffer []byte
	l.formatHeader(&buffer)
	fmt.Append(buffer, a...)
	buffer = fmt.Appendln(buffer, a...)
	l.WriteToStream(&buffer)
}

// Takes a buffer and writes to the specified output stream.
// It applies a newline to the end of the buffer if one does not already exist
func (l *GoLogger) WriteToStream(buffer *[]byte) {
	if (*buffer)[len(*buffer)-1] != '\n' {
		*buffer = append(*buffer, '\n')
	}
	l.outputSteam.Write(*buffer)
}
