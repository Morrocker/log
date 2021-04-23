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
	timestamp  bool
	preNote    bool
	color      bool
	mode       uint8
	scope      Scope
	timeFormat string
	outputFile string
	writeLock  sync.Mutex
	preMsg     preMessages
}

type preMessages struct {
	red, yellow, green, blue, cyan, magenta string
}

type Scope struct {
	Regular      bool
	RegularWrite bool
	Verbose      bool
	VerboseWrite bool
	Debug        bool
	DebugWrite   bool
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
		scope: Scope{Regular: true, Verbose: true, Debug: true},
	}
	return newLogger
}

// Info works like a fmt.Printf however it adds, datetime, a prefix label and a return at the end.
func (l *Logger) Info(format string, a ...interface{}) {
	l.doLog(Regular, "cyan", format+"\n", a...)
}

// Info works like a fmt.Printf however it adds, datetime, a prefix label and a return at the end.
func (l *Logger) Infoln(a ...interface{}) {
	l.doLog(Regular, "cyan", "%s\n", fmt.Sprint(a...))
}

// InfoV same as Info(), but will only print when verbose or debug options are set
func (l *Logger) InfoV(format string, a ...interface{}) {
	l.doLog(Verbose, "cyan", format+"\n", a...)
}

// InfoV same as Info(), but will only print when verbose or debug options are set
func (l *Logger) InfolnV(a ...interface{}) {
	l.doLog(Verbose, "cyan", "%s\n", fmt.Sprint(a...))
}

// InfoD same as Info(), but will only print when the debug options is set
func (l *Logger) InfoD(format string, a ...interface{}) {
	l.doLog(Debug, "cyan", format+"\n", a...)
}

// InfoD same as Info(), but will only print when the debug options is set
func (l *Logger) InfolnD(a ...interface{}) {
	l.doLog(Debug, "cyan", "%s\n", fmt.Sprint(a...))
}

// Alert works like a fmt.Printf however it adds, datetime, a prefix label and a return at the end
func (l *Logger) Alert(format string, a ...interface{}) {
	l.doLog(Regular, "yellow", format+"\n", a...)
}

// Alert works like a fmt.Printf however it adds, datetime, a prefix label and a return at the end
func (l *Logger) Alertln(a ...interface{}) {
	l.doLog(Regular, "yellow", "%s\n", fmt.Sprint(a...))
}

// AlertV same as Alert(), but will only print when verbose or debug options are set
func (l *Logger) AlertV(format string, a ...interface{}) {
	l.doLog(Verbose, "yellow", format+"\n", a...)
}

// Alert works like a fmt.Printf however it adds, datetime, a prefix label and a return at the end
func (l *Logger) AlertlnV(a ...interface{}) {
	l.doLog(Verbose, "yellow", "%s\n", fmt.Sprint(a...))
}

// AlertD same as Alert(), but will only print when the debug options is set
func (l *Logger) AlertD(format string, a ...interface{}) {
	l.doLog(Debug, "yellow", format+"\n", a...)
}

// Alert works like a fmt.Printf however it adds, datetime, a prefix label and a return at the end
func (l *Logger) AlertlnD(a ...interface{}) {
	l.doLog(Debug, "yellow", "%s\n", fmt.Sprint(a...))
}

// Error works like a fmt.Printf however it adds, datetime, a prefix label and a return at the end
func (l *Logger) Error(format string, a ...interface{}) {
	l.doLog(Regular, "red", format+"\n", a...)
}

// Error works like a fmt.Printf however it adds, datetime, a prefix label and a return at the end
func (l *Logger) Errorln(a ...interface{}) {
	l.doLog(Regular, "red", "%s\n", fmt.Sprint(a...))
}

// ErrorV same as Error(), but will only print when verbose or debug options are set
func (l *Logger) ErrorV(format string, a ...interface{}) {
	l.doLog(Verbose, "red", format+"\n", a...)
}

// Error works like a fmt.Printf however it adds, datetime, a prefix label and a return at the end
func (l *Logger) ErrorlnV(a ...interface{}) {
	l.doLog(Verbose, "red", "%s\n", fmt.Sprint(a...))
}

// ErrorD same as Error(), but will only print when the debug options is set
func (l *Logger) ErrorD(format string, a ...interface{}) {
	l.doLog(Debug, "red", format+"\n", a...)
}

// Error works like a fmt.Printf however it adds, datetime, a prefix label and a return at the end
func (l *Logger) ErrorlnD(a ...interface{}) {
	l.doLog(Debug, "red", "%s\n", fmt.Sprint(a...))
}

// Notice works like a fmt.Printf however it adds, datetime, a prefix label and a return at the end
func (l *Logger) Notice(format string, a ...interface{}) {
	l.doLog(Regular, "blue", format+"\n", a...)
}

// Notice works like a fmt.Printf however it adds, datetime, a prefix label and a return at the end
func (l *Logger) Noticeln(a ...interface{}) {
	l.doLog(Regular, "blue", "%s\n", fmt.Sprint(a...))
}

// NoticeV same as Notice(), but will only print when verbose or debug options are set
func (l *Logger) NoticeV(format string, a ...interface{}) {
	l.doLog(Verbose, "blue", format+"\n", a...)
}

// Notice works like a fmt.Printf however it adds, datetime, a prefix label and a return at the end
func (l *Logger) NoticelnV(a ...interface{}) {
	l.doLog(Verbose, "blue", "%s\n", fmt.Sprint(a...))
}

// NoticeD same as Notice(), but will only print when the debug options is set
func (l *Logger) NoticeD(format string, a ...interface{}) {
	l.doLog(Debug, "blue", format+"\n", a...)
}

// Notice works like a fmt.Printf however it adds, datetime, a prefix label and a return at the end
func (l *Logger) NoticelnD(a ...interface{}) {
	l.doLog(Debug, "blue", "%s\n", fmt.Sprint(a...))
}

// Task works like a fmt.Printf however it adds, datetime, a prefix label and a return at the end.
func (l *Logger) Task(format string, a ...interface{}) {
	l.doLog(Regular, "green", format+"\n", a...)
}

// Task works like a fmt.Printf however it adds, datetime, a prefix label and a return at the end.
func (l *Logger) Taskln(a ...interface{}) {
	l.doLog(Regular, "green", "%s\n", fmt.Sprint(a...))
}

// TaskV same as Task(), but will only print when verbose or debug options are set
func (l *Logger) TaskV(format string, a ...interface{}) {
	l.doLog(Verbose, "green", format+"\n", a...)
}

// Task works like a fmt.Printf however it adds, datetime, a prefix label and a return at the end.
func (l *Logger) TasklnV(a ...interface{}) {
	l.doLog(Verbose, "green", "%s\n", fmt.Sprint(a...))
}

// TaskV same as Task(), but will only print when verbose or debug options are set
func (l *Logger) TaskD(format string, a ...interface{}) {
	l.doLog(Debug, "green", format+"\n", a...)
}

// Task works like a fmt.Printf however it adds, datetime, a prefix label and a return at the end.
func (l *Logger) TasklnD(a ...interface{}) {
	l.doLog(Debug, "green", "%s\n", fmt.Sprint(a...))
}

// Notice works like a fmt.Printf however it adds, datetime, a prefix label and a return at the end
func (l *Logger) Bench(format string, a ...interface{}) {
	l.doLog(Regular, "magenta", format+"\n", a...)
}

// Notice works like a fmt.Printf however it adds, datetime, a prefix label and a return at the end
func (l *Logger) Benchln(a ...interface{}) {
	l.doLog(Regular, "magenta", "%s\n", fmt.Sprint(a...))
}

// Notice works like a fmt.Printf however it adds, datetime, a prefix label and a return at the end
func (l *Logger) BenchV(format string, a ...interface{}) {
	l.doLog(Verbose, "magenta", format+"\n", a...)
}

// Notice works like a fmt.Printf however it adds, datetime, a prefix label and a return at the end
func (l *Logger) BenchlnV(a ...interface{}) {
	l.doLog(Verbose, "magenta", "%s\n", fmt.Sprint(a...))
}

// Notice works like a fmt.Printf however it adds, datetime, a prefix label and a return at the end
func (l *Logger) BenchD(format string, a ...interface{}) {
	l.doLog(Debug, "magenta", format+"\n", a...)
}

// Notice works like a fmt.Printf however it adds, datetime, a prefix label and a return at the end
func (l *Logger) BenchlnD(a ...interface{}) {
	l.doLog(Debug, "magenta", "%s\n", fmt.Sprint(a...))
}
