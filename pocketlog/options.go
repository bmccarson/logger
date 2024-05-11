package pocketlog

import "io"

// Optioin defines a functional option to Logger.
type Option func(*Logger)

func WithOutput(output io.Writer) Option {
	return func(lgr *Logger) {
		lgr.output = output
	}
}
