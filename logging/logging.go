// Package logging defines the minimal leveled-logging interfaces used
// throughout Signalilo, decoupling callers from any specific logging
// backend.
package logging

// InfoLogger can log non-error messages, optionally gated by a verbosity
// level.
type InfoLogger interface {
	// Info logs a non-error message.
	Info(args ...interface{})

	// Infof logs a formatted non-error message.
	Infof(format string, args ...interface{})

	// Enabled reports whether this InfoLogger is enabled at its level.
	Enabled() bool
}

// Logger logs both informational and error messages.
type Logger interface {
	InfoLogger

	// Error logs an error message.
	Error(args ...interface{})

	// Errorf logs a formatted error message.
	Errorf(format string, args ...interface{})

	// V returns an InfoLogger for the given verbosity level. Higher levels
	// are less important and are silenced once they exceed the logger's
	// configured verbosity.
	V(level int) InfoLogger
}
