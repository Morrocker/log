package log

import "github.com/fatih/color"

var (
	red     = color.New(color.FgHiRed).SprintFunc()
	cyan    = color.New(color.FgHiCyan).SprintFunc()
	green   = color.New(color.FgHiGreen).SprintFunc()
	yellow  = color.New(color.FgHiYellow).SprintFunc()
	blue    = color.New(color.FgHiBlue).SprintFunc()
	magenta = color.New(color.FgHiMagenta).SprintFunc()
)
