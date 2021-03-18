package log

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

// SetModes sets the verbose and debug variables according to given parameters
func (l *Logger) SetMode(mode string) {
	if mode == "verbose" {
		l.mode = Verbose
	} else if mode == "debug" {
		l.mode = Debug
	} else {
		fmt.Println("logger mode set not valid")
	}
}

// ToggleSilent enables/disables silent mode. No logs will be shown if enabled. Note that this does not prevent file logging.
func (l *Logger) ToggleSilent() {
	l.silent = !l.silent
}

// ToggleTimestamp enables/disables timestamp on log
func (l *Logger) ToggleTimestamp() {
	l.timestamp = !l.timestamp
}

// ToggleTimestamp enables/disables timestamp on log
func (l *Logger) ToggleColor() {
	l.color = !l.color
}

// TogglePreNote enables/disables timestamp on log
func (l *Logger) TogglePreNote() {
	l.preNote = !l.preNote
}

// TogglePreNote enables/disables timestamp on log
func (l *Logger) ToggleDualMode() {
	l.dualMode = !l.dualMode
}

func (l *Logger) logFormat(format string) string {
	var ret string
	if l.timestamp && l.preNote {
		ret = "%s\t%s: "
	} else if l.timestamp {
		ret = "%s: "
	} else if l.preNote {
		ret = "%s\t"
	}
	return ret + format
}

func (l *Logger) printLog(color, format string, a ...interface{}) {
	var writter io.Writer = os.Stdout
	format = l.logFormat(format)
	a = l.coalesce(l.getPreMsg(color, l.color), a...)
	if color == "red" {
		writter = os.Stderr
	}

	fmt.Fprintf(writter, format, a...)
}

func (l *Logger) writeLog(n int, color string, format string, a ...interface{}) {
	l.writeLock.Lock()
	defer l.writeLock.Unlock()
	format = l.logFormat(format)
	a = l.coalesce(l.getPreMsg(color, false), a...)

	f, err := os.OpenFile(l.outputFile,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	txt := fmt.Sprintf(format, a...)
	if _, err := f.WriteString(txt); err != nil {
		log.Println(err)
	}
}

func (l *Logger) doLog(n int, color string, format string, a ...interface{}) {
	switch {
	case uint8(n) > l.mode:
		return

	case !l.write && l.silent:
		return

	case l.write && l.silent:
		write := false
		switch n {
		case Regular:
			if l.writeScope.Regular {
				write = true
			}
		case Verbose:
			if l.writeScope.Verbose {
				write = true
			}
		case Debug:
			if l.writeScope.Debug {
				write = true
			}
		}
		if write {
			l.writeLog(n, color, format, a...)
		}

	case !l.write && !l.silent:
		l.printLog(color, format, a...)

	default:

		write := false
		switch n {
		case Regular:
			if l.writeScope.Regular {
				write = true
			}
		case Verbose:
			if l.writeScope.Verbose {
				write = true
			}
		case Debug:
			if l.writeScope.Debug {
				write = true
			}
		}

		if write {
			l.writeLog(n, color, format, a...)
			if l.dualMode {
				l.printLog(color, format, a...)
			}
			return
		}
		l.printLog(color, format, a...)
	}
}

func (l *Logger) coalesce(header string, a ...interface{}) []interface{} {
	d := l.getDate()
	var ret []interface{}
	if l.timestamp && l.preNote {
		ret = []interface{}{header, d}
	} else if l.timestamp {
		ret = []interface{}{d}
	} else if l.preNote {
		ret = []interface{}{header}
	}
	ret = append(ret, a...)
	return ret
}

func (l *Logger) getPreMsg(color string, colorize bool) string {
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

func (l *Logger) getDate() string {
	return time.Now().Format(l.timeFormat)
}

// OutputToFile sets the log to be printed to a file instead of StdOut. What gets written can be tuned.
func (l *Logger) OutputFile(filename string) {
	l.outputFile = filename
}

// StopLogToFile turns off the OutputToFile option and resets these preferences
func (l *Logger) StopWriter() {
	l.write = false
}

func (l *Logger) StartWriter() {
	if l.outputFile == "" {
		fmt.Println("Writter cannot start without output file")
		return
	}
	l.write = true
}

func (l *Logger) SetScope(r, v, d bool) {
	l.writeScope = Scope{r, v, d}
}
