package logger

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
)

// Level defines the degrees of severity
type Level int

const (
	DebugLevel Level = iota
	InfoLevel
	WarnLevel
	ErrorLevel
)

// LoggerRepository is the interface that is used everywhere in the code
type LoggerRepository interface {
	Debugf(format string, v ...any)
	Infof(format string, v ...any)
	Warnf(format string, v ...any)
	Errorf(format string, v ...any)
	SetLevel(level Level)
	SetOutput(w io.Writer)
	SetOutputFile(path string) error
}

// Logger is the standard implementation with std log.Logger
type Logger struct {
	mu    sync.RWMutex
	level Level
	Debug *log.Logger
	Info  *log.Logger
	Warn  *log.Logger
	Error *log.Logger
}

// NewLogger creates a logger with Name-Prefix, Level and Writer
func NewLogger(name string, level Level, out io.Writer) *Logger {
	flags := log.Ldate | log.Ltime | log.Lshortfile
	return &Logger{
		level: level,
		Debug: log.New(out, name+" DEBUG: ", flags),
		Info:  log.New(out, name+" INFO:  ", flags),
		Warn:  log.New(out, name+" WARN:  ", flags),
		Error: log.New(out, name+" ERROR: ", flags),
	}
}

// SetLevel changes the minimum level
func (l *Logger) SetLevel(level Level) {
	l.mu.Lock()
	l.level = level
	l.mu.Unlock()
}

// SetOutput sets the writer for all levels
func (l *Logger) SetOutput(w io.Writer) {
	l.Debug.SetOutput(w)
	l.Info.SetOutput(w)
	l.Warn.SetOutput(w)
	l.Error.SetOutput(w)
}

// SetOutputFile opens (or creates) the named file (append-modus)
// and uses it for all level loggers
func (l *Logger) SetOutputFile(path string) error {
	// Open file for writing (Append, Create, Write-Only)
	file, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0o644)
	if err != nil {
		return fmt.Errorf("could not open log file %q: %w", path, err)
	}
	l.SetOutput(file)
	return nil
}

// Internal auxiliary function with level filter and optional colors
func (l *Logger) logf(lvl Level, logger *log.Logger, format string, v ...any) {
	l.mu.RLock()
	defer l.mu.RUnlock()
	if lvl < l.level {
		return
	}
	// color edition
	var wrapped string
	switch lvl {
	case DebugLevel:
		wrapped = "\033[0;32m" + format + "\033[0m"
	case InfoLevel:
		wrapped = "\033[0;34m" + format + "\033[0m"
	case WarnLevel:
		wrapped = "\033[0;33m" + format + "\033[0m"
	case ErrorLevel:
		wrapped = "\033[0;31m" + format + "\033[0m"
	default:
		wrapped = format
	}
	logger.Printf(wrapped, v...)
}

func (l *Logger) Debugf(format string, v ...any) { l.logf(DebugLevel, l.Debug, format, v...) }
func (l *Logger) Infof(format string, v ...any)  { l.logf(InfoLevel, l.Info, format, v...) }
func (l *Logger) Warnf(format string, v ...any)  { l.logf(WarnLevel, l.Warn, format, v...) }
func (l *Logger) Errorf(format string, v ...any) { l.logf(ErrorLevel, l.Error, format, v...) }

// ------------------------------------------------
// Context functions to pass on loggers in context
// ------------------------------------------------
type ctxKey struct{}

func WithLogger(ctx context.Context, l LoggerRepository) context.Context {
	return context.WithValue(ctx, ctxKey{}, l)
}

func FromContext(ctx context.Context) LoggerRepository {
	if l, ok := ctx.Value(ctxKey{}).(LoggerRepository); ok {
		return l
	}
	// Fallback to a default info logger
	return NewLogger("gobpmn", InfoLevel, os.Stdout)
}
