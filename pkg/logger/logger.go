package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

func Initialize() {
	var err error

	encoderConfig := zapcore.EncoderConfig{
		TimeKey: "time",
		LevelKey: "level",
		NameKey: "logger",
		CallerKey: "caller",
		MessageKey: "msg",
		StacktraceKey: "stacktrace",
		LineEnding: zapcore.DefaultLineEnding,
		EncodeLevel: zapcore.CapitalLevelEncoder,
		EncodeTime: zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller: zapcore.ShortCallerEncoder,
	}

	logFile, err := os.OpenFile("/var/log/api/api.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	fileWriter := zapcore.AddSync(logFile)

	loglevel := zapcore.DebugLevel

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		fileWriter,
		loglevel,
	)

	Logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))
	
}

// Sync flushes any buffered log entries
func Sync() {
	_ = Logger.Sync()
}