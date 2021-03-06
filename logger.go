package log

import (
	"fmt"
	"log"
	"os"
	"time"
)

// ToggleTimestamp enables/disables timestamp on log
func (l *logger) ToggleTimestamp() {
	l.timestamp = !l.timestamp
}

// ToggleTimestamp enables/disables timestamp on log
func (l *logger) ToggleColor() {
	l.color = !l.color
}

// TogglePreNote enables/disables timestamp on log
func (l *logger) TogglePreNote() {
	l.preNote = !l.preNote
}

func (l *logger) logFormat(format string) string {
	var ret string
	if l.timestamp && l.preNote {
		ret = "%s\t%s: "
	} else if l.timestamp {
		ret = "%s: "
	} else if l.preNote {
		ret = "%s\t"
	}
	return fmt.Sprintf("%s%s", ret, format)
}

func (l *logger) printLog(color, format string, a ...interface{}) {
	writter := os.Stdout
	format = l.logFormat(format)
	a = l.coalesce(l.getPreMsg(color, l.color), a...)
	if color == "red" {
		writter = os.Stderr
	}

	fmt.Fprintf(writter, format, a...)
}

func (l *logger) writeLog(n int, color string, format string, a ...interface{}) {
	l.writeLock.Lock()
	defer l.writeLock.Unlock()
	format = l.logFormat(format)
	a = l.coalesce(l.getPreMsg(color, false), a...)

	if l.outputFile == "" {
		l.outputFile = "./default.log"
	}
	f, err := os.OpenFile(l.outputFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	txt := fmt.Sprintf(format, a...)
	if _, err := f.WriteString(txt); err != nil {
		log.Println(err)
	}
}

func (l *logger) doLog(n int, color string, format string, a ...interface{}) {
	switch {
	case l.scope.Regular && n == 1:
		l.printLog(color, format, a...)
	case l.scope.Verbose && n == 2:
		l.printLog(color, format, a...)
	case l.scope.Debug && n == 3:
		l.printLog(color, format, a...)
	}
	switch {
	case l.scope.RegularWrite && n == 1:
		l.writeLog(n, color, format, a...)
	case l.scope.VerboseWrite && n == 2:
		l.writeLog(n, color, format, a...)
	case l.scope.DebugWrite && n == 3:
		l.writeLog(n, color, format, a...)
	}
}

func (l *logger) coalesce(header string, a ...interface{}) []interface{} {
	d := l.getDate()
	if l.timestamp && l.preNote {
		return append([]interface{}{header, d}, a...)
	} else if l.timestamp {
		return append([]interface{}{d}, a...)
	} else if l.preNote {
		return append([]interface{}{header}, a...)
	}
	return nil
}

func (l *logger) getPreMsg(color string, colorize bool) string {
	var msg string
	if colorize {
		switch color {
		case "blue":
			msg = blue(l.preMsg.blue)
		case "red":
			msg = red(l.preMsg.red)
		case "yellow":
			msg = yellow(l.preMsg.yellow)
		case "cyan":
			msg = cyan(l.preMsg.cyan)
		case "green":
			msg = green(l.preMsg.green)
		case "magenta":
			msg = magenta(l.preMsg.magenta)
		}
	} else {
		switch color {
		case "blue":
			msg = l.preMsg.blue
		case "red":
			msg = l.preMsg.red
		case "yellow":
			msg = l.preMsg.yellow
		case "cyan":
			msg = l.preMsg.cyan
		case "green":
			msg = l.preMsg.green
		case "magenta":
			msg = l.preMsg.magenta
		}
	}
	return msg
}

func (l *logger) getDate() string {
	return time.Now().Format(l.timeFormat)
}

// OutputToFile sets the log to be printed to a file instead of StdOut. What gets written can be tuned.
func (l *logger) OutputFile(filename string) {
	l.outputFile = filename
}

func (l *logger) SetScope(r, rw, v, vw, d, dw bool) {
	l.scope = Scope{r, rw, v, vw, d, dw}
}

func (l *logger) SetRegularScope(r, w bool) {
	l.scope.Regular = r
	l.scope.RegularWrite = w
}
func (l *logger) SetVerboseScope(r, w bool) {
	l.scope.Verbose = r
	l.scope.VerboseWrite = w
}
func (l *logger) SetDebugScope(r, w bool) {
	l.scope.Debug = r
	l.scope.DebugWrite = w
}
