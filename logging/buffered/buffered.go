// Package buffered implements logging.Logger backed by an in-memory buffer,
// for use in tests.
package buffered

import (
	"bytes"
	"fmt"
	"strings"
	"sync"

	"github.com/saremox/signalilo/logging"
)

const (
	levelError = "ERROR "
	levelInfo  = "INFO "
	levelV     = "V[%d] "
)

// New creates a new buffered Logger. verbosity is the highest V() level
// that still produces output.
func New(verbosity int) *Logger {
	return &Logger{verbosity: verbosity, buf: &bytes.Buffer{}, mu: &sync.Mutex{}}
}

// Logger is a logging.Logger that writes to an in-memory buffer.
type Logger struct {
	level     int
	verbosity int
	buf       *bytes.Buffer
	mu        *sync.Mutex
}

// Info implements logging.Logger.
func (l *Logger) Info(args ...interface{}) {
	if l.Enabled() {
		l.writeLine(l.levelString(), fmt.Sprint(args...))
	}
}

// Infof implements logging.Logger.
func (l *Logger) Infof(format string, args ...interface{}) {
	if l.Enabled() {
		l.writeLine(l.levelString(), fmt.Sprintf(format, args...))
	}
}

// Enabled implements logging.Logger.
func (l *Logger) Enabled() bool {
	return l.level <= l.verbosity
}

// Error implements logging.Logger.
func (l *Logger) Error(args ...interface{}) {
	l.writeLine(levelError, fmt.Sprint(args...))
}

// Errorf implements logging.Logger.
func (l *Logger) Errorf(format string, args ...interface{}) {
	l.writeLine(levelError, fmt.Sprintf(format, args...))
}

// V implements logging.Logger.
func (l *Logger) V(level int) logging.InfoLogger {
	return &Logger{level: level, verbosity: l.verbosity, buf: l.buf, mu: l.mu}
}

// Buf returns the internal buffer holding all lines written so far.
func (l *Logger) Buf() *bytes.Buffer {
	return l.buf
}

func (l *Logger) levelString() string {
	if l.level > 0 {
		return fmt.Sprintf(levelV, l.level)
	}
	return levelInfo
}

func (l *Logger) writeLine(level, line string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.buf.WriteString(level)
	l.buf.WriteString(line)
	if !strings.HasSuffix(line, "\n") {
		l.buf.WriteRune('\n')
	}
}
