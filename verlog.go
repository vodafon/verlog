package verlog

import (
	"log"
	"os"
)

// A Logger represents an active logging object that generates lines of output to an StdOut.
type Logger struct {
	*log.Logger
	vl int
}

// New creates a new Logger with verbose level. Smaller level - less output.
func New(verboseLevel int) *Logger {
	return &Logger{
		Logger: log.New(os.Stdout, "", log.LstdFlags),
		vl:     verboseLevel,
	}
}

// Printf calls l.Output to print to the logger.
func (obj *Logger) Printf(vl int, format string, v ...interface{}) {
	if obj.vl < vl {
		return
	}
	obj.Logger.Printf(format, v...)
}

// Print calls l.Output to print to the logger.
func (obj *Logger) Print(vl int, v ...interface{}) {
	if obj.vl < vl {
		return
	}
	obj.Logger.Print(v...)
}

// Println calls l.Output to print to the logger.
func (obj *Logger) Println(vl int, v ...interface{}) {
	if obj.vl < vl {
		return
	}
	obj.Logger.Println(v...)
}
