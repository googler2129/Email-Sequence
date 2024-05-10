package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	zapLogger *zap.Logger
	level     *zap.AtomicLevel
}

type config struct {
	zap.Config
	callerSkip int
}

var (
	logger *Logger
)

func NewLogger() error {
	conf := &config{zap.NewProductionConfig(), 1} // default skip of 1

	zapLogger, err := conf.Build(
		zap.AddStacktrace(zapcore.DPanicLevel), // Add stack trace above this level automatically
		zap.AddCallerSkip(conf.callerSkip),     // skip configured level to print the correct caller
	)
	if err != nil {
		return err
	}

	logger = &Logger{
		zapLogger: zapLogger,
		level:     &conf.Level,
	}

	return nil
}

// WithError WithContext adds an error as single field to the Entry.
func WithError(err error) *Logger {
	return &Logger{
		zapLogger: logger.zapLogger.With(zap.Error(err)),
		level:     logger.level,
	}
}

// WithField creates an entry from the standard logger and adds a field to
// it.
//
// Note that it doesn't log until you call Debug, Print, Info, Warn, Fatal
// or Panic on the Entry it returns.
func WithField(key string, value interface{}) *Logger {
	return &Logger{
		zapLogger: logger.zapLogger.Sugar().With(key, value).Desugar(),
		level:     logger.level,
	}
}

// GetLevel returns the current logging level
// Logging levels can be - "debug", "info", "warning", "error", "fatal"
func GetLevel() string {
	return logger.level.String()
}

// Debug log a message at DEBUG level using the provided logger
func Debug(args ...interface{}) {
	logger.zapLogger.Sugar().Debug(args...)
}

// Debugf log a message at DEBUG level using the provided logger
func Debugf(format string, args ...interface{}) {
	logger.zapLogger.Sugar().Debugf(format, args...)
}

// Debugln log a message at DEBUG level using the provided logger
func Debugln(args ...interface{}) {
	logger.zapLogger.Sugar().Debug(args...)
}

// Print log a message at INFO level using the provided logger.
func Print(args ...interface{}) {
	logger.zapLogger.Sugar().Info(args...)
}

// Printf log a message at INFO level using the provided logger.
func Printf(format string, args ...interface{}) {
	logger.zapLogger.Sugar().Infof(format, args...)
}

// Println log a message at INFO level using the provided logger
func Println(args ...interface{}) {
	logger.zapLogger.Sugar().Info(args...)
}

// Info log a message at INFO level using the provided logger
func Info(args ...interface{}) {
	logger.zapLogger.Sugar().Info(args...)
}

// Infof log a message at INFO level using the provided logger
func Infof(format string, args ...interface{}) {
	logger.zapLogger.Sugar().Infof(format, args...)
}

// Infoln log a message at INFO level using the provided logger
func Infoln(args ...interface{}) {
	logger.zapLogger.Sugar().Info(args...)
}

// Warn log a message at WARNING level using the provided logger
func Warn(args ...interface{}) {
	logger.zapLogger.Sugar().Warn(args...)
}

// Warnf log a message at WARNING level using the provided logger
func Warnf(format string, args ...interface{}) {
	logger.zapLogger.Sugar().Warnf(format, args...)
}

// Warnln log a message at WARNING level using the provided logger
func Warnln(args ...interface{}) {
	logger.zapLogger.Sugar().Warn(args...)
}

// Error log a message at ERROR level using the provided logger
func Error(args ...interface{}) {
	logger.zapLogger.Sugar().Error(args...)
}

// Errorf log a message at ERROR level using the provided logger
func Errorf(format string, args ...interface{}) {
	logger.zapLogger.Sugar().Errorf(format, args...)
}

// Errorln log a message at ERROR level using the provided logger
func Errorln(args ...interface{}) {
	logger.zapLogger.Sugar().Error(args...)
}

// Panic log a message at PANIC level using the provided logger
func Panic(args ...interface{}) {
	logger.zapLogger.Sugar().Panic(args...)
}

// Panicf log a message at PANIC level using the given logger
func Panicf(format string, args ...interface{}) {
	logger.zapLogger.Sugar().Panicf(format, args...)
}

// Panicln log a message at PANIC level using the given logger
func Panicln(args ...interface{}) {
	logger.zapLogger.Sugar().Panic(args...)
}
