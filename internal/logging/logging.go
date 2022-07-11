package logging

// Level represents a log level.
type Level int32

const (
	Trace  Level = 0
	Debug  Level = 1
	Info   Level = 2
	Warn   Level = 3
	Error  Level = 4
	Fatal  Level = 5
	Silent Level = 6
)

type Logger interface {
	Log(level Level, msg string)
}
