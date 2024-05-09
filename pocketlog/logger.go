package pocketlog

// Logger is used to log information.
type Logger struct {
	threshold Level
}

// New returns you a logger, ready to log at the required threshold.
func New(threshold Level) *Logger {
	return &Logger{
		threshold: threshold,
	}
}
