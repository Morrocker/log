package log

import (
	"fmt"
	"sync"
)

const (
	None = iota
	Regular
	Verbose
	Debug
)

type Logger struct {
	write      bool
	silent     bool
	timestamp  bool
	preNote    bool
	dualMode   bool
	color      bool
	mode       uint8
	writeScope Scope
	timeFormat string
	outputFile string
	writeLock  sync.Mutex
	preMsg     preMessages
}

type preMessages struct {
	red, yellow, green, blue, cyan, magenta string
}

type Scope struct {
	Regular bool
	Verbose bool
	Debug   bool
}

func New() *Logger {
	newLogger := &Logger{
		timestamp:  true,
		preNote:    true,
		color:      true,
		mode:       Regular,
		timeFormat: "2006-01-02 15:04:05",
		preMsg: preMessages{
			red:     "[ERROR]",
			cyan:    "[INFO]",
			green:   "[TASK]",
			yellow:  "[ALERT]",
			blue:    "[NOTE]",
			magenta: "[BENCH]",
		},
		writeScope: Scope{true, true, true},
	}
	return newLogger
}

// Info works like a fmt.Printf however it adds, datetime, a prefix label and a return at the end.
func (l *Logger) Info(format string, a ...interface{}) {
	l.doLog(1, "cyan", format, a...)
}

// Info works like a fmt.Printf however it adds, datetime, a prefix label and a return at the end.
func (l *Logger) Infoln(a ...interface{}) {
	l.doLog(1, "cyan", "%s", fmt.Sprintln(a...))
}

// InfoV same as Info(), but will only print when verbose or debug options are set
func (l *Logger) InfoV(format string, a ...interface{}) {
	l.doLog(2, "cyan", format, a...)
}

// InfoV same as Info(), but will only print when verbose or debug options are set
func (l *Logger) InfolnV(format string, a ...interface{}) {
	l.doLog(2, "cyan", "%s", fmt.Sprintln(a...))
}

// InfoD same as Info(), but will only print when the debug options is set
func (l *Logger) InfoD(format string, a ...interface{}) {
	l.doLog(3, "cyan", format, a...)
}

// InfoD same as Info(), but will only print when the debug options is set
func (l *Logger) InfolnD(format string, a ...interface{}) {
	l.doLog(3, "cyan", "%s", fmt.Sprintln(a...))
}

// Alert works like a fmt.Printf however it adds, datetime, a prefix label and a return at the end
func (l *Logger) Alert(format string, a ...interface{}) {
	l.doLog(1, "yellow", format, a...)
}

// Alert works like a fmt.Printf however it adds, datetime, a prefix label and a return at the end
func (l *Logger) Alertln(format string, a ...interface{}) {
	l.doLog(1, "yellow", "%s", fmt.Sprintln(a...))
}

// AlertV same as Alert(), but will only print when verbose or debug options are set
func (l *Logger) AlertV(format string, a ...interface{}) {
	l.doLog(2, "yellow", format, a...)
}

// Alert works like a fmt.Printf however it adds, datetime, a prefix label and a return at the end
func (l *Logger) AlertlnV(format string, a ...interface{}) {
	l.doLog(2, "yellow", "%s", fmt.Sprintln(a...))
}

// AlertD same as Alert(), but will only print when the debug options is set
func (l *Logger) AlertD(format string, a ...interface{}) {
	l.doLog(3, "yellow", format, a...)
}

// Alert works like a fmt.Printf however it adds, datetime, a prefix label and a return at the end
func (l *Logger) AlertlnD(format string, a ...interface{}) {
	l.doLog(3, "yellow", "%s", fmt.Sprintln(a...))
}

// Error works like a fmt.Printf however it adds, datetime, a prefix label and a return at the end
func (l *Logger) Error(format string, a ...interface{}) {
	l.doLog(1, "red", format, a...)
}

// Error works like a fmt.Printf however it adds, datetime, a prefix label and a return at the end
func (l *Logger) Errorln(format string, a ...interface{}) {
	l.doLog(1, "red", "%s", fmt.Sprintln(a...))
}

// ErrorV same as Error(), but will only print when verbose or debug options are set
func (l *Logger) ErrorV(format string, a ...interface{}) {
	l.doLog(2, "red", format, a...)
}

// Error works like a fmt.Printf however it adds, datetime, a prefix label and a return at the end
func (l *Logger) ErrorlnV(format string, a ...interface{}) {
	l.doLog(2, "red", "%s", fmt.Sprintln(a...))
}

// ErrorD same as Error(), but will only print when the debug options is set
func (l *Logger) ErrorD(format string, a ...interface{}) {
	l.doLog(3, "red", format, a...)
}

// Error works like a fmt.Printf however it adds, datetime, a prefix label and a return at the end
func (l *Logger) ErrorlnD(format string, a ...interface{}) {
	l.doLog(3, "red", "%s", fmt.Sprintln(a...))
}

// Notice works like a fmt.Printf however it adds, datetime, a prefix label and a return at the end
func (l *Logger) Notice(format string, a ...interface{}) {
	l.doLog(1, "blue", format, a...)
}

// Notice works like a fmt.Printf however it adds, datetime, a prefix label and a return at the end
func (l *Logger) Noticeln(format string, a ...interface{}) {
	l.doLog(1, "blue", "%s", fmt.Sprintln(a...))
}

// NoticeV same as Notice(), but will only print when verbose or debug options are set
func (l *Logger) NoticeV(format string, a ...interface{}) {
	l.doLog(2, "blue", format, a...)
}

// Notice works like a fmt.Printf however it adds, datetime, a prefix label and a return at the end
func (l *Logger) NoticelnV(format string, a ...interface{}) {
	l.doLog(2, "blue", "%s", fmt.Sprintln(a...))
}

// NoticeD same as Notice(), but will only print when the debug options is set
func (l *Logger) NoticeD(format string, a ...interface{}) {
	l.doLog(3, "blue", format, a...)
}

// Notice works like a fmt.Printf however it adds, datetime, a prefix label and a return at the end
func (l *Logger) NoticelnD(format string, a ...interface{}) {
	l.doLog(3, "blue", "%s", fmt.Sprintln(a...))
}

// Task works like a fmt.Printf however it adds, datetime, a prefix label and a return at the end.
func (l *Logger) Task(format string, a ...interface{}) {
	l.doLog(1, "green", format, a...)
}

// Task works like a fmt.Printf however it adds, datetime, a prefix label and a return at the end.
func (l *Logger) Taskln(format string, a ...interface{}) {
	l.doLog(1, "green", "%s", fmt.Sprintln(a...))
}

// TaskV same as Task(), but will only print when verbose or debug options are set
func (l *Logger) TaskV(format string, a ...interface{}) {
	l.doLog(2, "green", format, a...)
}

// Task works like a fmt.Printf however it adds, datetime, a prefix label and a return at the end.
func (l *Logger) TasklnV(format string, a ...interface{}) {
	l.doLog(2, "green", "%s", fmt.Sprintln(a...))
}

// TaskV same as Task(), but will only print when verbose or debug options are set
func (l *Logger) TaskD(format string, a ...interface{}) {
	l.doLog(3, "green", format, a...)
}

// Task works like a fmt.Printf however it adds, datetime, a prefix label and a return at the end.
func (l *Logger) TasklnD(format string, a ...interface{}) {
	l.doLog(3, "green", "%s", fmt.Sprintln(a...))
}

// Notice works like a fmt.Printf however it adds, datetime, a prefix label and a return at the end
func (l *Logger) Bench(format string, a ...interface{}) {
	l.doLog(1, "magenta", format, a...)
}

// Notice works like a fmt.Printf however it adds, datetime, a prefix label and a return at the end
func (l *Logger) Benchln(format string, a ...interface{}) {
	l.doLog(1, "magenta", "%s", fmt.Sprintln(a...))
}

// Notice works like a fmt.Printf however it adds, datetime, a prefix label and a return at the end
func (l *Logger) BenchV(format string, a ...interface{}) {
	l.doLog(2, "magenta", format, a...)
}

// Notice works like a fmt.Printf however it adds, datetime, a prefix label and a return at the end
func (l *Logger) BenchlnV(format string, a ...interface{}) {
	l.doLog(2, "magenta", "%s", fmt.Sprintln(a...))
}

// Notice works like a fmt.Printf however it adds, datetime, a prefix label and a return at the end
func (l *Logger) BenchD(format string, a ...interface{}) {
	l.doLog(3, "magenta", format, a...)
}

// Notice works like a fmt.Printf however it adds, datetime, a prefix label and a return at the end
func (l *Logger) BenchlnD(format string, a ...interface{}) {
	l.doLog(3, "magenta", "%s", fmt.Sprintln(a...))
}
