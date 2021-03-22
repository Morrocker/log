package log

import "fmt"

var oneLogger *Logger

func init() {
	oneLogger = New()
}

// Info works like a fmt.Printf however it adds, datetime, a prefix label and a return at the end.
func Info(format string, a ...interface{}) {
	oneLogger.Info(format, a...)
}

// Info works like a fmt.Printf however it adds, datetime, a prefix label and a return at the end.
func Infoln(a ...interface{}) {
	oneLogger.Infoln(a...)
}

// InfoV same as Info(), but will only print when verbose or debug options are set
func InfoV(format string, a ...interface{}) {
	oneLogger.InfoV(format, a...)
}

// InfoV same as Info(), but will only print when verbose or debug options are set
func InfolnV(a ...interface{}) {
	oneLogger.InfolnV(a...)
}

// InfoD same as Info(), but will only print when the debug options is set
func InfoD(format string, a ...interface{}) {
	oneLogger.InfoD(format, a...)
}

// InfoD same as Info(), but will only print when the debug options is set
func InfolnD(a ...interface{}) {
	oneLogger.InfolnD(a...)
}

// Alert works like a fmt.Printf however it adds, datetime, a prefix label and a return at the end
func Alert(format string, a ...interface{}) {
	oneLogger.Alert(format, a...)
}

// Alert works like a fmt.Printf however it adds, datetime, a prefix label and a return at the end
func Alertln(a ...interface{}) {
	oneLogger.Alertln(a...)
}

// AlertV same as Alert(), but will only print when verbose or debug options are set
func AlertV(format string, a ...interface{}) {
	oneLogger.AlertV(format, a...)
}

// AlertV same as Alert(), but will only print when verbose or debug options are set
func AlertlnV(a ...interface{}) {
	oneLogger.AlertlnV(a...)
}

// AlertD same as Alert(), but will only print when the debug options is set
func AlertD(format string, a ...interface{}) {
	oneLogger.AlertD(format, a...)
}

// AlertD same as Alert(), but will only print when the debug options is set
func AlertlnD(a ...interface{}) {
	oneLogger.AlertlnD(a...)
}

// Error works like a fmt.Printf however it adds, datetime, a prefix label and a return at the end
func Error(format string, a ...interface{}) {
	oneLogger.Error(format, a...)
}

// Error works like a fmt.Printf however it adds, datetime, a prefix label and a return at the end
func Errorln(a ...interface{}) {
	oneLogger.Errorln(a...)
}

// ErrorV same as Error(), but will only print when verbose or debug options are set
func ErrorV(format string, a ...interface{}) {
	oneLogger.ErrorV(format, a...)
}

// ErrorV same as Error(), but will only print when verbose or debug options are set
func ErrorlnV(a ...interface{}) {
	oneLogger.ErrorlnV(a...)
}

// ErrorD same as Error(), but will only print when the debug options is set
func ErrorD(format string, a ...interface{}) {
	oneLogger.ErrorD(format, a...)
}

// ErrorD same as Error(), but will only print when the debug options is set
func ErrorlnD(a ...interface{}) {
	oneLogger.ErrorlnD(a...)
}

// Notice works like a fmt.Printf however it adds, datetime, a prefix label and a return at the end
func Notice(format string, a ...interface{}) {
	oneLogger.Notice(format, a...)
}

// Notice works like a fmt.Printf however it adds, datetime, a prefix label and a return at the end
func Noticeln(a ...interface{}) {
	oneLogger.Noticeln(a...)
}

// NoticeV same as Notice(), but will only print when verbose or debug options are set
func NoticeV(format string, a ...interface{}) {
	oneLogger.NoticeV(format, a...)
}

// NoticeV same as Notice(), but will only print when verbose or debug options are set
func NoticelnV(a ...interface{}) {
	oneLogger.NoticelnV(a...)
}

// NoticeD same as Notice(), but will only print when the debug options is set
func NoticeD(format string, a ...interface{}) {
	oneLogger.NoticeD(format, a...)
}

// NoticeD same as Notice(), but will only print when the debug options is set
func NoticelnD(a ...interface{}) {
	oneLogger.NoticelnD(a...)
}

// Task works like a fmt.Printf however it adds, datetime, a prefix label and a return at the end.
func Task(format string, a ...interface{}) {
	oneLogger.Task(format, a...)
}

// Task works like a fmt.Printf however it adds, datetime, a prefix label and a return at the end.
func Taskln(a ...interface{}) {
	oneLogger.Taskln(a...)
}

// TaskV same as Task(), but will only print when verbose or debug options are set
func TaskV(format string, a ...interface{}) {
	oneLogger.TaskV(format, a...)
}

// TaskV same as Task(), but will only print when verbose or debug options are set
func TasklnV(a ...interface{}) {
	oneLogger.TasklnV(a...)
}

// TaskD same as Task(), but will only print when the debug options is set
func TaskD(format string, a ...interface{}) {
	oneLogger.TaskD(format, a...)
}

// TaskD same as Task(), but will only print when the debug options is set
func TasklnD(a ...interface{}) {
	oneLogger.TasklnD(a...)
}

// Notice works like a fmt.Printf however it adds, datetime, a prefix label and a return at the end
func Bench(format string, a ...interface{}) {
	oneLogger.Bench(format, a...)
}

// Notice works like a fmt.Printf however it adds, datetime, a prefix label and a return at the end
func Benchln(a ...interface{}) {
	oneLogger.Benchln(a...)
}

// Notice works like a fmt.Printf however it adds, datetime, a prefix label and a return at the end
func BenchV(format string, a ...interface{}) {
	oneLogger.BenchV(format, a...)
}

// Notice works like a fmt.Printf however it adds, datetime, a prefix label and a return at the end
func BenchlnV(a ...interface{}) {
	oneLogger.BenchlnV(a...)
}

// Notice works like a fmt.Printf however it adds, datetime, a prefix label and a return at the end
func BenchD(format string, a ...interface{}) {
	oneLogger.BenchD(format, a...)
}

// Notice works like a fmt.Printf however it adds, datetime, a prefix label and a return at the end
func BenchlnD(a ...interface{}) {
	oneLogger.BenchlnD(a...)
}

// OutputToFile sets the log to be printed to a file instead of StdOut. What gets written can be tuned.
func OutputFile(filename string) {
	oneLogger.OutputFile(filename)
}

// SetModes sets the verbose and debug variables according to given parameters
func SetMode(mode string) {
	oneLogger.SetMode(mode)
}

// ToggleSilent enables/disables silent mode. No logs will be shown if enabled. Note that this does not prevent file logging.
func ToggleSilent() {
	oneLogger.ToggleSilent()
}

// ToggleTimestamp enables/disables timestamp on log
func ToggleTimestamp() {
	oneLogger.ToggleTimestamp()
}

// ToggleTimestamp enables/disables timestamp on log
func ToggleColor() {
	oneLogger.ToggleColor()
}

// TogglePreNote enables/disables timestamp on log
func TogglePreNote() {
	oneLogger.TogglePreNote()
}

// TogglePreNote enables/disables timestamp on log
func ToggleDualMode() {
	oneLogger.ToggleDualMode()
}

// StopLogToFile turns off the OutputToFile option and resets these preferences
func StopWriter() {
	oneLogger.StopWriter()
}

func StartWriter() {
	if oneLogger.outputFile == "" {
		fmt.Println("Writter cannot start without output file")
		return
	}
	oneLogger.StartWriter()
}

func SetScope(r, v, d bool) {
	oneLogger.SetScope(r, v, d)
}
