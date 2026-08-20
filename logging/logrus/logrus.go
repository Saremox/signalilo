// Package logrus implements logging.Logger backed by a *logrus.Logger.
package logrus

import (
	"github.com/saremox/signalilo/logging"
	"github.com/sirupsen/logrus"
)

// New creates a logging.Logger backed by l. verbosity is the highest V()
// level that still produces output; V() levels above zero log at Debug
// instead of Info.
func New(verbosity int, l *logrus.Logger) logging.Logger {
	return &logger{verbosity: verbosity, logger: l}
}

type logger struct {
	level     int
	verbosity int
	logger    *logrus.Logger
}

func (l *logger) Info(args ...interface{}) {
	if l.Enabled() {
		if l.level > 0 {
			l.logger.Debug(args...)
		} else {
			l.logger.Info(args...)
		}
	}
}

func (l *logger) Infof(format string, args ...interface{}) {
	if l.Enabled() {
		if l.level > 0 {
			l.logger.Debugf(format, args...)
		} else {
			l.logger.Infof(format, args...)
		}
	}
}

func (l *logger) Enabled() bool {
	return l.level <= l.verbosity
}

func (l *logger) Error(args ...interface{}) {
	l.logger.Error(args...)
}

func (l *logger) Errorf(format string, args ...interface{}) {
	l.logger.Errorf(format, args...)
}

func (l *logger) V(level int) logging.InfoLogger {
	return &logger{level: level, verbosity: l.verbosity, logger: l.logger}
}
