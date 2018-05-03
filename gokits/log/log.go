package log

import (
	"fmt"
	"io"
	std_log "log"
	"os"
	"path"
	"runtime"
)

var _ = std_log.Print

// DefaultLogger logger
var DefaultLogger *Logger

func init() {
	DefaultLogger = NewLogger(os.Stdout)
}

// Logger logger
type Logger struct {
	Level     Level
	Colorful  bool
	ShowLine  bool
	Prefix    string
	StackSkip int

	l *std_log.Logger
}

// NewLogger new logger
func NewLogger(w io.Writer) *Logger {
	return &Logger{
		Level:     TRACE,
		Colorful:  true,
		ShowLine:  true,
		StackSkip: 4,
		l:         std_log.New(w, "", std_log.LstdFlags),
	}
}

// NewLoggerBySkip new logger by skip
func NewLoggerBySkip(w io.Writer, skip int) *Logger {
	return &Logger{
		Level:     TRACE,
		Colorful:  true,
		ShowLine:  true,
		StackSkip: skip,
		l:         std_log.New(w, "", std_log.LstdFlags),
	}
}

func color(col, s string) string {
	if col == "" {
		return s
	}
	return "\x1b[0;" + col + "m" + s + "\x1b[0m"
}

// SetPrefix set prefix
func (l *Logger) SetPrefix(prefix string) *Logger {
	l.Prefix = prefix
	return l
}

// Skip skip
func (l *Logger) Skip(skip int) *Logger {
	nl := *l
	nl.StackSkip = skip
	return &nl
}

func (l *Logger) getPrefix(level Level) string {
	s := level.String()
	if l.Colorful {
		s = color(level.Color(), s)
	}
	s = "[" + s + "]"
	if l.Prefix != "" {
		s = s + " " + l.Prefix
	}
	return s
}

func (l *Logger) getPosix() string {
	if !l.ShowLine {
		return ""
	}
	_, file, line, ok := runtime.Caller(l.StackSkip)

	if !ok {
		return ""
	}
	file = path.Base(file)
	return fmt.Sprintf("[%s:%d]", file, line)
}

// Output output log
func (l *Logger) Output(level Level, format string, a ...interface{}) {
	if level < l.Level {
		return
	}
	var s string
	if len(a) != 0 {
		s = fmt.Sprintf(format, a...)
	} else {
		s = fmt.Sprint(format)
	}
	content := l.getPrefix(level) + " " + s + " " + l.getPosix()
	l.l.Println(content)
}

// Tracef trace formated
func (l *Logger) Tracef(format string, a ...interface{}) {
	l.Output(TRACE, format, a...)
}

// Debugf debug formated
func (l *Logger) Debugf(format string, a ...interface{}) {
	l.Output(DEBUG, format, a...)
}

// Printf trace formated
func (l *Logger) Printf(format string, a ...interface{}) {
	l.Output(TRACE, format, a...)
}

// Infof info formated
func (l *Logger) Infof(format string, a ...interface{}) {
	l.Output(INFO, format, a...)
}

// Warnf warn formated
func (l *Logger) Warnf(format string, a ...interface{}) {
	l.Output(WARN, format, a...)
}

// Fatalf fatal formated
func (l *Logger) Fatalf(format string, a ...interface{}) {
	l.Output(FATAL, format, a...)
	os.Exit(1)
}

// Errorf error formated
func (l *Logger) Errorf(format string, a ...interface{}) {
	l.Output(ERROR, format, a...)
}

// Tracef trace formated
func Tracef(format string, a ...interface{}) {
	DefaultLogger.Tracef(format, a...)
}

// Printf trace formated
func Printf(format string, a ...interface{}) {
	DefaultLogger.Printf(format, a...)
}

// Debugf debug formated
func Debugf(format string, a ...interface{}) {
	DefaultLogger.Debugf(format, a...)
}

// Infof info formated
func Infof(format string, a ...interface{}) {
	DefaultLogger.Infof(format, a...)
}

// Warnf warn formated
func Warnf(format string, a ...interface{}) {
	DefaultLogger.Warnf(format, a...)
}

// Errorf error formated
func Errorf(format string, a ...interface{}) {
	DefaultLogger.Errorf(format, a...)
}

// Fatalf fatal formated
func Fatalf(format string, a ...interface{}) {
	DefaultLogger.Fatalf(format, a...)
}
