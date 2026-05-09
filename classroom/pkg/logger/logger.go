package logger

import (
	"os"

	"github.com/fernoe1/WATEC/classroom/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type logger interface {
	InitLogger()
	Debug(args ...interface{})
	Debugf(template string, args ...interface{})
	Info(args ...interface{})
	Infof(template string, args ...interface{})
	Warn(args ...interface{})
	Warnf(template string, args ...interface{})
	Error(args ...interface{})
	Errorf(template string, args ...interface{})
	DPanic(args ...interface{})
	DPanicf(template string, args ...interface{})
	Fatal(args ...interface{})
	Fatalf(template string, args ...interface{})
	Printf(template string, args ...interface{})
}

var levels = map[string]zapcore.Level{
	"debug":  zapcore.DebugLevel,
	"info":   zapcore.InfoLevel,
	"warn":   zapcore.WarnLevel,
	"error":  zapcore.ErrorLevel,
	"dpanic": zapcore.DPanicLevel,
	"panic":  zapcore.PanicLevel,
	"fatal":  zapcore.FatalLevel,
}

type Logger struct {
	logger *zap.SugaredLogger
	cfg    *config.Config
}

func (l *Logger) getCfgs() (encoding string, level zapcore.Level, caller, stacktrace bool) {

	return l.cfg.Logger.Encoding,
		levels[l.cfg.Logger.Level],
		l.cfg.Logger.Caller,
		l.cfg.Logger.Stacktrace
}

func (l *Logger) getEncoder(encoding string) zapcore.Encoder {
	var encoderCfg zapcore.EncoderConfig
	if l.cfg.Server.Development {
		encoderCfg = zap.NewDevelopmentEncoderConfig()
	} else {
		encoderCfg = zap.NewProductionEncoderConfig()
	}
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	if encoding == "console" {

		return zapcore.NewConsoleEncoder(encoderCfg)
	} else {

		return zapcore.NewJSONEncoder(encoderCfg)
	}
}

func buildOptions(caller, stacktrace bool, level zapcore.Level) []zap.Option {
	var opts []zap.Option

	if caller {
		opts = append(opts, zap.AddCaller(), zap.AddCallerSkip(1))
	}
	if stacktrace {
		opts = append(opts, zap.AddStacktrace(level))
	}

	return opts
}

func (l *Logger) InitLogger() {
	encoding, level, caller, stacktrace := l.getCfgs()

	l.logger = zap.New(
		zapcore.NewCore(l.getEncoder(encoding),
			zapcore.AddSync(os.Stderr),
			zap.NewAtomicLevelAt(level),
		),
		buildOptions(caller, stacktrace, level)...).Sugar()
}

func (l *Logger) Debug(args ...interface{}) {
	l.logger.Debug(args...)
}

func (l *Logger) Debugf(template string, args ...interface{}) {
	l.logger.Debugf(template, args...)
}

func (l *Logger) Info(args ...interface{}) {
	l.logger.Info(args...)
}

func (l *Logger) Infof(template string, args ...interface{}) {
	l.logger.Infof(template, args...)
}

func (l *Logger) Warn(args ...interface{}) {
	l.logger.Warn(args...)
}

func (l *Logger) Warnf(template string, args ...interface{}) {
	l.logger.Warnf(template, args...)
}

func (l *Logger) Error(args ...interface{}) {
	l.logger.Error(args...)
}

func (l *Logger) Errorf(template string, args ...interface{}) {
	l.logger.Errorf(template, args...)
}

func (l *Logger) DPanic(args ...interface{}) {
	l.logger.DPanic(args...)
}

func (l *Logger) DPanicf(template string, args ...interface{}) {
	l.logger.DPanicf(template, args...)
}

func (l *Logger) Fatal(args ...interface{}) {
	l.logger.Fatal(args...)
}

func (l *Logger) Fatalf(template string, args ...interface{}) {
	l.logger.Fatalf(template, args...)
}

func (l *Logger) Printf(template string, args ...interface{}) {
	l.logger.Infof(template, args...)
}
