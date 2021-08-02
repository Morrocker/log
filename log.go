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

type Logger interface {
	ToggleTimestamp()
	ToggleColor()
	TogglePreNote()
	OutputFile(string)
	SetScope(bool, bool, bool, bool, bool, bool)
	SetRegularScope(bool, bool)
	SetVerboseScope(bool, bool)
	SetDebugScope(bool, bool)

	Info(string, ...interface{})
	InfoV(string, ...interface{})
	InfoD(string, ...interface{})
	Infoln(...interface{})
	InfolnV(...interface{})
	InfolnD(...interface{})

	Alert(string, ...interface{})
	AlertV(string, ...interface{})
	AlertD(string, ...interface{})
	Alertln(...interface{})
	AlertlnV(...interface{})
	AlertlnD(...interface{})

	Error(string, ...interface{})
	ErrorV(string, ...interface{})
	ErrorD(string, ...interface{})
	Errorln(...interface{})
	ErrorlnV(...interface{})
	ErrorlnD(...interface{})

	Notice(string, ...interface{})
	NoticeV(string, ...interface{})
	NoticeD(string, ...interface{})
	Noticeln(...interface{})
	NoticelnV(...interface{})
	NoticelnD(...interface{})

	Task(string, ...interface{})
	TaskV(string, ...interface{})
	TaskD(string, ...interface{})
	Taskln(...interface{})
	TasklnV(...interface{})
	TasklnD(...interface{})

	Bench(string, ...interface{})
	BenchV(string, ...interface{})
	BenchD(string, ...interface{})
	Benchln(...interface{})
	BenchlnV(...interface{})
	BenchlnD(...interface{})
}

type logger struct {
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

// New returns an initialized Logger interface
func New() Logger {
	newLogger := &logger{
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
func (l *logger) Info(format string, a ...interface{}) {
	l.doLog(Regular, "cyan", format+"\n", a...)
}

// InfoV same as Info(), but will only print when verbose or debug options are set
func (l *logger) InfoV(format string, a ...interface{}) {
	l.doLog(Verbose, "cyan", format+"\n", a...)
}

// InfoD same as Info(), but will only print when the debug options is set
func (l *logger) InfoD(format string, a ...interface{}) {
	l.doLog(Debug, "cyan", format+"\n", a...)
}

// Infoln works like a fmt.Printf however it adds, datetime, a prefix label and a return at the end.
func (l *logger) Infoln(a ...interface{}) {
	l.doLog(Regular, "cyan", "%s\n", fmt.Sprint(a...))
}

// InfoV same as Info(), but will only print when verbose or debug options are set
func (l *logger) InfolnV(a ...interface{}) {
	l.doLog(Verbose, "cyan", "%s\n", fmt.Sprint(a...))
}

// InfoD same as Info(), but will only print when the debug options is set
func (l *logger) InfolnD(a ...interface{}) {
	l.doLog(Debug, "cyan", "%s\n", fmt.Sprint(a...))
}

// Alert works like a fmt.Printf however it adds, datetime, a prefix label and a return at the end
func (l *logger) Alert(format string, a ...interface{}) {
	l.doLog(Regular, "yellow", format+"\n", a...)
}

// AlertV same as Alert(), but will only print when verbose or debug options are set
func (l *logger) AlertV(format string, a ...interface{}) {
	l.doLog(Verbose, "yellow", format+"\n", a...)
}

// AlertD same as Alert(), but will only print when the debug options is set
func (l *logger) AlertD(format string, a ...interface{}) {
	l.doLog(Debug, "yellow", format+"\n", a...)
}

// Alert works like a fmt.Printf however it adds, datetime, a prefix label and a return at the end
func (l *logger) Alertln(a ...interface{}) {
	l.doLog(Regular, "yellow", "%s\n", fmt.Sprint(a...))
}

// Alert works like a fmt.Printf however it adds, datetime, a prefix label and a return at the end
func (l *logger) AlertlnV(a ...interface{}) {
	l.doLog(Verbose, "yellow", "%s\n", fmt.Sprint(a...))
}

// Alert works like a fmt.Printf however it adds, datetime, a prefix label and a return at the end
func (l *logger) AlertlnD(a ...interface{}) {
	l.doLog(Debug, "yellow", "%s\n", fmt.Sprint(a...))
}

// Error works like a fmt.Printf however it adds, datetime, a prefix label and a return at the end
func (l *logger) Error(format string, a ...interface{}) {
	l.doLog(Regular, "red", format+"\n", a...)
}

// ErrorV same as Error(), but will only print when verbose or debug options are set
func (l *logger) ErrorV(format string, a ...interface{}) {
	l.doLog(Verbose, "red", format+"\n", a...)
}

// ErrorD same as Error(), but will only print when the debug options is set
func (l *logger) ErrorD(format string, a ...interface{}) {
	l.doLog(Debug, "red", format+"\n", a...)
}

// Error works like a fmt.Printf however it adds, datetime, a prefix label and a return at the end
func (l *logger) Errorln(a ...interface{}) {
	l.doLog(Regular, "red", "%s\n", fmt.Sprint(a...))
}

// Error works like a fmt.Printf however it adds, datetime, a prefix label and a return at the end
func (l *logger) ErrorlnV(a ...interface{}) {
	l.doLog(Verbose, "red", "%s\n", fmt.Sprint(a...))
}

// Error works like a fmt.Printf however it adds, datetime, a prefix label and a return at the end
func (l *logger) ErrorlnD(a ...interface{}) {
	l.doLog(Debug, "red", "%s\n", fmt.Sprint(a...))
}

// Notice works like a fmt.Printf however it adds, datetime, a prefix label and a return at the end
func (l *logger) Notice(format string, a ...interface{}) {
	l.doLog(Regular, "blue", format+"\n", a...)
}

// NoticeV same as Notice(), but will only print when verbose or debug options are set
func (l *logger) NoticeV(format string, a ...interface{}) {
	l.doLog(Verbose, "blue", format+"\n", a...)
}

// NoticeD same as Notice(), but will only print when the debug options is set
func (l *logger) NoticeD(format string, a ...interface{}) {
	l.doLog(Debug, "blue", format+"\n", a...)
}

// Notice works like a fmt.Printf however it adds, datetime, a prefix label and a return at the end
func (l *logger) Noticeln(a ...interface{}) {
	l.doLog(Regular, "blue", "%s\n", fmt.Sprint(a...))
}

// Notice works like a fmt.Printf however it adds, datetime, a prefix label and a return at the end
func (l *logger) NoticelnV(a ...interface{}) {
	l.doLog(Verbose, "blue", "%s\n", fmt.Sprint(a...))
}

// Notice works like a fmt.Printf however it adds, datetime, a prefix label and a return at the end
func (l *logger) NoticelnD(a ...interface{}) {
	l.doLog(Debug, "blue", "%s\n", fmt.Sprint(a...))
}

// Task works like a fmt.Printf however it adds, datetime, a prefix label and a return at the end.
func (l *logger) Task(format string, a ...interface{}) {
	l.doLog(Regular, "green", format+"\n", a...)
}

// TaskV same as Task(), but will only print when verbose or debug options are set
func (l *logger) TaskV(format string, a ...interface{}) {
	l.doLog(Verbose, "green", format+"\n", a...)
}

// TaskV same as Task(), but will only print when verbose or debug options are set
func (l *logger) TaskD(format string, a ...interface{}) {
	l.doLog(Debug, "green", format+"\n", a...)
}

// Task works like a fmt.Printf however it adds, datetime, a prefix label and a return at the end.
func (l *logger) Taskln(a ...interface{}) {
	l.doLog(Regular, "green", "%s\n", fmt.Sprint(a...))
}

// Task works like a fmt.Printf however it adds, datetime, a prefix label and a return at the end.
func (l *logger) TasklnV(a ...interface{}) {
	l.doLog(Verbose, "green", "%s\n", fmt.Sprint(a...))
}

// Task works like a fmt.Printf however it adds, datetime, a prefix label and a return at the end.
func (l *logger) TasklnD(a ...interface{}) {
	l.doLog(Debug, "green", "%s\n", fmt.Sprint(a...))
}

// Notice works like a fmt.Printf however it adds, datetime, a prefix label and a return at the end
func (l *logger) Bench(format string, a ...interface{}) {
	l.doLog(Regular, "magenta", format+"\n", a...)
}

// Notice works like a fmt.Printf however it adds, datetime, a prefix label and a return at the end
func (l *logger) BenchV(format string, a ...interface{}) {
	l.doLog(Verbose, "magenta", format+"\n", a...)
}

// Notice works like a fmt.Printf however it adds, datetime, a prefix label and a return at the end
func (l *logger) BenchD(format string, a ...interface{}) {
	l.doLog(Debug, "magenta", format+"\n", a...)
}

// Notice works like a fmt.Printf however it adds, datetime, a prefix label and a return at the end
func (l *logger) Benchln(a ...interface{}) {
	l.doLog(Regular, "magenta", "%s\n", fmt.Sprint(a...))
}

// Notice works like a fmt.Printf however it adds, datetime, a prefix label and a return at the end
func (l *logger) BenchlnV(a ...interface{}) {
	l.doLog(Verbose, "magenta", "%s\n", fmt.Sprint(a...))
}

// Notice works like a fmt.Printf however it adds, datetime, a prefix label and a return at the end
func (l *logger) BenchlnD(a ...interface{}) {
	l.doLog(Debug, "magenta", "%s\n", fmt.Sprint(a...))
}
