package logger

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"time"
)

// This type is here to simply enforce types when using the colorizeString function
type AnsiColor string

// These are all of the ansi colors and the reset character
var (
	Reset   AnsiColor = "\033[0m"
	Red     AnsiColor = "\033[31m"
	Green   AnsiColor = "\033[32m"
	Yellow  AnsiColor = "\033[33m"
	Blue    AnsiColor = "\033[34m"
	Magenta AnsiColor = "\033[35m"
	Cyan    AnsiColor = "\033[36m"
	Gray    AnsiColor = "\033[37m"
	White   AnsiColor = "\033[97m"
)

type Status string

var (
	GOAL        Status = "GOAL"
	IN_PROGRESS Status = "IN PROGRESS"
	ERROR       Status = "ERROR"
	SUCCESS     Status = "SUCCESS"
)

// This function takes a string to colorize and the desired color.
// It returns a string with the color code as a prefix and the reset code as a postfix.
func ColorizeString(s string, c AnsiColor) string {
	return string(c) + s + string(Reset)
}

type GoLogger struct {
	prefix       string
	outputSteam  io.Writer
	dateAndTime  bool
	fileName     bool
	functionName bool
}

func NewGoLogger(prefix string, outputStream io.Writer, dateAndTime bool, fileName bool, functionName bool) *GoLogger {
	return &GoLogger{
		prefix:       prefix,
		outputSteam:  outputStream,
		dateAndTime:  dateAndTime,
		fileName:     fileName,
		functionName: functionName,
	}
}

func (l *GoLogger) generateLog(buffer *[]byte) {
	l.formatHeader(buffer)
}

func (l *GoLogger) formatHeader(buffer *[]byte) {
	showDateAndTime := l.dateAndTime
	showFileName := l.fileName
	showFunctionName := l.functionName

	pc, file, _, ok := runtime.Caller(2)

	if l.prefix != "" {
		formattedPrefix := "[" + ColorizeString(l.prefix, Blue) + "]"
		*buffer = append(*buffer, []byte(formattedPrefix)...)
	}

	if showDateAndTime {
		currentTime := time.Now()
		formattedDate := fmt.Sprintf("%d/%d/%d", currentTime.Year(), currentTime.Month(), currentTime.Day())
		formattedTime := fmt.Sprintf("%d:%d:%d", currentTime.Local().Hour(), currentTime.Local().Minute(), currentTime.Local().Second())
		formattedDateAndTime := formattedDate + " " + formattedTime
		*buffer = append(*buffer, []byte("["+ColorizeString(formattedDateAndTime, Cyan)+"]")...)
	}

	if showFileName && ok {
		var shortendFileName []byte
		for index := range file {
			if file[len(file)-(1+index)] != '/' {
				var tempBuffer []byte
				tempBuffer = append(tempBuffer, file[len(file)-(1+index)])
				shortendFileName = append(tempBuffer, shortendFileName...)
			} else {
				break
			}
		}
		*buffer = append(*buffer, []byte("["+ColorizeString(string(shortendFileName), Magenta)+"]")...)
	}

	if showFunctionName && ok {
		fn := runtime.FuncForPC(pc)
		functionName := fn.Name()

		var shortendFunctionName []byte

		for index, _ := range functionName {
			if functionName[len(functionName)-(1+index)] != '.' {
				var tempBuffer []byte
				tempBuffer = append(tempBuffer, functionName[len(functionName)-(1+index)])
				shortendFunctionName = append(tempBuffer, shortendFunctionName...)
			} else {
				break
			}
		}
		*buffer = append(*buffer, []byte("["+ColorizeString(string(shortendFunctionName), Magenta)+"]")...)
	}

	if l.prefix != "" || l.dateAndTime || l.fileName || l.functionName {
		*buffer = append(*buffer, " - "...)
	}

}

// Prints a log message with the same formatting as fmt.Print() but with log header prefix
func (l *GoLogger) Print(a ...any) {
	var buffer []byte
	l.formatHeader(&buffer)
	buffer = fmt.Append(buffer, a...)
	l.WriteToStream(&buffer)
}

// Prints a log message with the same formatting as fmt.Printf() but with log header prefix
func (l *GoLogger) Printf(format string, a ...any) {
	var buffer []byte
	l.formatHeader(&buffer)
	buffer = fmt.Appendf(buffer, format, a...)
	l.WriteToStream(&buffer)
}

// Prints a log message with the same formatting as fmt.Println() but with log header prefix
func (l *GoLogger) Println(a ...any) {
	var buffer []byte
	l.formatHeader(&buffer)
	buffer = fmt.Appendln(buffer, a...)
	l.WriteToStream(&buffer)
}

// Prints a log message with the same formatting as fmt.Print() but with log header prefix
func (l *GoLogger) StatusPrint(status Status, a ...any) {
	var buffer []byte
	l.formatHeader(&buffer)
	buffer = fmt.Append(buffer, a...)
	l.WriteToStream(&buffer)
}

// Prints a log message with the same formatting as fmt.Printf() but with log header prefix
func (l *GoLogger) StatusPrintf(status Status, format string, a ...any) {
	var buffer []byte
	l.formatHeader(&buffer)
	buffer = fmt.Appendf(buffer, format, a...)
	l.WriteToStream(&buffer)
}

// Prints a log message with the same formatting as fmt.Println() but with log header prefix
func (l *GoLogger) StatusPrintln(status Status, a ...any) {
	var buffer []byte
	l.formatHeader(&buffer)
	buffer = fmt.Appendln(buffer, a...)
	l.WriteToStream(&buffer)
}

// Prints a log message with the same formatting as fmt.Print() but with log header prefix and then exits with status 1
func (l *GoLogger) Fatal(a ...any) {
	var buffer []byte
	l.formatHeader(&buffer)
	buffer = fmt.Append(buffer, a...)
	l.WriteToStream(&buffer)
	os.Exit(1)
}

// Prints a log message with the same formatting as fmt.Printf() but with log header prefix and then exits with status 1
func (l *GoLogger) Fatalf(format string, a ...any) {
	var buffer []byte
	l.formatHeader(&buffer)
	buffer = fmt.Appendf(buffer, format, a...)
	l.WriteToStream(&buffer)
	os.Exit(1)
}

// Prints a log message with the same formatting as fmt.Println() but with log header prefix and then exits with status 1
func (l *GoLogger) Fatalln(a ...any) {
	var buffer []byte
	l.formatHeader(&buffer)
	buffer = fmt.Appendln(buffer, a...)
	l.WriteToStream(&buffer)
	os.Exit(1)
}

// Prints a log message with the same formatting as fmt.Print() but with log header prefix and then panics
func (l *GoLogger) Panic(a ...any) {
	var buffer []byte
	l.formatHeader(&buffer)
	buffer = fmt.Append(buffer, a...)
	l.WriteToStream(&buffer)
	panic(fmt.Sprint(a...))
}

// Prints a log message with the same formatting as fmt.Printf() but with log header prefix and then panics
func (l *GoLogger) Panicf(format string, a ...any) {
	var buffer []byte
	l.formatHeader(&buffer)
	buffer = fmt.Appendf(buffer, format, a...)
	l.WriteToStream(&buffer)
	panic(fmt.Sprintf(format, a...))
}

// Prints a log message with the same formatting as fmt.Println() but with log header prefix and then panics
func (l *GoLogger) Panicln(a ...any) {
	var buffer []byte
	l.formatHeader(&buffer)
	buffer = fmt.Appendln(buffer, a...)
	l.WriteToStream(&buffer)
	panic(fmt.Sprintln(a...))
}

// Takes a buffer and writes to the specified output stream.
// It applies a newline to the end of the buffer if one does not already exist
func (l *GoLogger) WriteToStream(buffer *[]byte) {
	if (*buffer)[len(*buffer)-1] != '\n' {
		*buffer = append(*buffer, '\n')
	}
	l.outputSteam.Write(*buffer)
}
