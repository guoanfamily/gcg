package log

import (
	std_log "log"
	"strings"
)

var _ = std_log.Print

// Level Level
type Level int

const (
	// DEBUG debug level
	DEBUG Level = iota
	// TRACE trace level
	TRACE
	// INFO info level
	INFO
	// WARN warn level
	WARN
	// ERROR error level
	ERROR
	// FATAL fatal level
	FATAL
)

var levelColor = map[Level]string{
	TRACE: "36",
	DEBUG: "32",
	INFO:  "33",
	WARN:  "31",
	ERROR: "31",
	FATAL: "35",
}

// NewLevel new level
func NewLevel(s string) Level {
	s = strings.ToLower(s)
	switch s {
	case "trace":
		return TRACE
	case "debug":
		return DEBUG
	case "info":
		return INFO
	case "warn":
		return WARN
	case "error":
		return ERROR
	case "fatal":
		return FATAL
	default:
		std_log.Printf("invalid log level '%s'", s)
		return TRACE
	}
}

// Color level color
func (l Level) Color() string {
	return levelColor[l]
}
