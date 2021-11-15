package std

// Level enumerates available log levels.
type Level int

var levels = []string{
	"PANI", "ERRO", "WARN", "INFO", "DEBU", "TRAC",
}

// String implements Stringer interface.
func (l Level) String() string {
	if l > LevelTrace {
		return "UNKN"
	}

	return levels[l]
}

// Log levels.
const (
	LevelError Level = iota
	LevelWarning
	LevelInfo
	LevelDebug
	LevelTrace
)
