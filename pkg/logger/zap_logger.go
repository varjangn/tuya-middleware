package logger

import (
	"errors"
	"fmt"
	"os"
	"syscall"

	"github.com/varjangn/tuya-middleware/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger interface {
	InitLogger()
	Debug(args ...interface{})
	Debugf(tmpl string, args ...interface{})
	Debugw(msg string, args ...interface{})
	Info(args ...interface{})
	Infof(tmpl string, args ...interface{})
	Infow(msg string, args ...interface{})
	Warn(args ...interface{})
	Warnf(tmpl string, args ...interface{})
	Warnw(msg string, args ...interface{})
	Error(args ...interface{})
	Errorf(tmpl string, args ...interface{})
	Errorw(msg string, args ...interface{})
}

var loggerLevelMap = map[string]zapcore.Level{
	"debug":  zapcore.DebugLevel,
	"info":   zapcore.InfoLevel,
	"warn":   zapcore.WarnLevel,
	"error":  zapcore.ErrorLevel,
	"dpanic": zapcore.DPanicLevel,
	"panic":  zapcore.PanicLevel,
	"fatal":  zapcore.FatalLevel,
}

type AppLogger struct {
	cfg         *config.Config
	sugarLogger *zap.SugaredLogger
}

func NewAppLogger(cfg *config.Config) *AppLogger {
	return &AppLogger{cfg: cfg}
}

func (l *AppLogger) getLoggerLevel(cfg *config.Config) zapcore.Level {
	level, exist := loggerLevelMap[cfg.Logger.Level]
	if !exist {
		return zapcore.DebugLevel
	}
	return level
}

func (l *AppLogger) InitLogger() {
	logLevel := l.getLoggerLevel(l.cfg)

	logWriter := zapcore.AddSync(os.Stderr)

	var encoderCfg zapcore.EncoderConfig
	if l.cfg.Server.Mode == "Development" {
		encoderCfg = zap.NewDevelopmentEncoderConfig()
	} else {
		encoderCfg = zap.NewProductionEncoderConfig()
	}

	var encoder zapcore.Encoder
	encoderCfg.LevelKey = "LEVEL"
	encoderCfg.CallerKey = "CALLER"
	encoderCfg.TimeKey = "TIME"
	encoderCfg.NameKey = "NAME"
	encoderCfg.MessageKey = "MESSAGE"

	if l.cfg.Logger.Encoding == "console" {
		encoder = zapcore.NewConsoleEncoder(encoderCfg)
	} else {
		encoder = zapcore.NewJSONEncoder(encoderCfg)
	}

	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	en := zap.NewAtomicLevelAt(logLevel)
	core := zapcore.NewCore(encoder, logWriter, en)
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))

	l.sugarLogger = logger.Sugar()
	err := l.sugarLogger.Sync()
	if err != nil && !errors.Is(err, syscall.ENOTTY) {
		fmt.Println(err)
	}

}

// Logger methods

func (l *AppLogger) Debug(args ...interface{}) {
	l.sugarLogger.Debug(args...)
}

func (l *AppLogger) Debugf(template string, args ...interface{}) {
	l.sugarLogger.Debugf(template, args...)
}

func (l *AppLogger) Debugw(msg string, args ...interface{}) {
	l.sugarLogger.Debugw(msg, args...)
}

func (l *AppLogger) Info(args ...interface{}) {
	l.sugarLogger.Info(args...)
}

func (l *AppLogger) Infof(template string, args ...interface{}) {
	l.sugarLogger.Infof(template, args...)
}

func (l *AppLogger) Infow(msg string, args ...interface{}) {
	l.sugarLogger.Infow(msg, args...)
}

func (l *AppLogger) Warn(args ...interface{}) {
	l.sugarLogger.Warn(args...)
}

func (l *AppLogger) Warnf(template string, args ...interface{}) {
	l.sugarLogger.Warnf(template, args...)
}

func (l *AppLogger) Warnw(msg string, args ...interface{}) {
	l.sugarLogger.Warnw(msg, args...)
}

func (l *AppLogger) Error(args ...interface{}) {
	l.sugarLogger.Error(args...)
}

func (l *AppLogger) Errorf(template string, args ...interface{}) {
	l.sugarLogger.Errorf(template, args...)
}

func (l *AppLogger) Errorw(msg string, args ...interface{}) {
	l.sugarLogger.Errorw(msg, args...)
}

func (l *AppLogger) DPanic(args ...interface{}) {
	l.sugarLogger.DPanic(args...)
}

func (l *AppLogger) DPanicf(template string, args ...interface{}) {
	l.sugarLogger.DPanicf(template, args...)
}

func (l *AppLogger) DPanicw(msg string, args ...interface{}) {
	l.sugarLogger.DPanicw(msg, args...)
}

func (l *AppLogger) Panic(args ...interface{}) {
	l.sugarLogger.Panic(args...)
}

func (l *AppLogger) Panicf(template string, args ...interface{}) {
	l.sugarLogger.Panicf(template, args...)
}

func (l *AppLogger) Panicw(msg string, args ...interface{}) {
	l.sugarLogger.Panicw(msg, args...)
}

func (l *AppLogger) Fatal(args ...interface{}) {
	l.sugarLogger.Fatal(args...)
}

func (l *AppLogger) Fatalf(template string, args ...interface{}) {
	l.sugarLogger.Fatalf(template, args...)
}

func (l *AppLogger) Fatalw(msg string, args ...interface{}) {
	l.sugarLogger.Fatalw(msg, args...)
}
