package logs

import (
	"fmt"
	"sync"

	"github.com/data-market/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var mLogger *zap.SugaredLogger
var mLogLevel zap.AtomicLevel
var lk sync.Mutex

func Logger(name string) *zap.SugaredLogger {
	lk.Lock()
	defer lk.Unlock()
	return mLogger.Named(name)
}

func init() {
	mLogLevel = zap.NewAtomicLevel()

	outputs := []string{"stdout"}
	debugWriter, _, err := zap.Open(outputs...)
	if err != nil {
		panic(fmt.Sprintf("unable to open logging output: %v", err))
	}

	encoder := getEncoder()

	core := zapcore.NewCore(encoder, debugWriter, mLogLevel)

	logger := zap.New(core, zap.AddCaller())

	mLogger = logger.Sugar()

	l := getLogLevel(config.Cfg.LogLevel)

	mLogLevel.SetLevel(l)
}

func getLogLevel(level string) zapcore.Level {
	l := zapcore.InfoLevel
	switch level {
	case "debug", "DEBUG":
		l = zapcore.DebugLevel
	case "info", "INFO", "": // make the zero value useful
		l = zapcore.InfoLevel
	case "warn", "WARN":
		l = zapcore.WarnLevel
	case "error", "ERROR":
		l = zapcore.ErrorLevel
	case "dpanic", "DPANIC":
		l = zapcore.DPanicLevel
	case "panic", "PANIC":
		l = zapcore.PanicLevel
	case "fatal", "FATAL":
		l = zapcore.FatalLevel

	}

	return l
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "sub",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.RFC3339TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	return zapcore.NewJSONEncoder(encoderConfig)
}
