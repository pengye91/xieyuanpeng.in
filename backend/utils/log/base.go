package log

import (
	"os"
	"sort"

	"github.com/pengye91/xieyuanpeng.in/backend/configs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

var (
	DevCfg      = zap.NewDevelopmentConfig()
	ErrCfg      = zap.NewDevelopmentConfig()
	ProdCfg     = zap.NewProductionConfig()
	Logger      *zap.Logger
	LoggerSugar *zap.SugaredLogger
)

func init() {
	os.Setenv("GIN_ACCESS_LOG", os.Getenv("MYGOPATH")+"/xieyuanpeng.in/logs/gin_access.log")
	os.Setenv("GIN_ERROR_LOG", os.Getenv("MYGOPATH")+"/xieyuanpeng.in/logs/gin_error.log")
	var err error
	Logger, err = ProdLogger()
	if err != nil {
		panic(err)
	}
	defer Logger.Sync()
	LoggerSugar = Logger.Sugar()
}

func MyDevLogger() (*zap.Logger, error) {
	DevCfg.OutputPaths = []string{
		"stdout",
		os.Getenv("MYGOPATH") + "/xieyuanpeng.in/logs/dev.log",
	}
	// this is the logging system error output paths
	DevCfg.ErrorOutputPaths = []string{
		"stderr",
		os.Getenv("MYGOPATH") + "/xieyuanpeng.in/logs/sys_err.log",
	}
	return DevCfg.Build()
}

func MyErrLogger() (*zap.Logger, error) {
	dyn := zap.NewAtomicLevel()
	dyn.SetLevel(zap.ErrorLevel)
	ErrCfg.Level = dyn
	ErrCfg.OutputPaths = []string{
		"stderr",
		os.Getenv("MYGOPATH") + "/xieyuanpeng.in/logs/err.log",
	}
	// this is the logging system error output paths
	ErrCfg.ErrorOutputPaths = []string{
		"stderr",
		configs.ErrorLogPath,
	}
	return ErrCfg.Build()
}

func prodLogger() (*zap.Logger, error) {
	errorPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})
	accessPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.ErrorLevel
	})

	errorFile, _, errorErr := zap.Open([]string{configs.ErrorLogPath, "stderr"}...)
	if errorErr != nil {
		panic(errorErr)
	}
	accessConsole := zapcore.AddSync(os.Stdout)
	errorConsole := zapcore.AddSync(os.Stderr)

	accessFile, _, accessErr := zap.Open([]string{configs.AccessLogPath, "stdout"}...)
	if accessErr != nil {
		panic(accessErr)
	}

	errorEncoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}
	accessEncoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	errorEncoder := zapcore.NewJSONEncoder(errorEncoderConfig)
	accessEncoder := zapcore.NewConsoleEncoder(accessEncoderConfig)
	accessConsoleEncoder := zapcore.NewJSONEncoder(accessEncoderConfig)
	errorConsoleEncoder := zapcore.NewConsoleEncoder(errorEncoderConfig)

	core := zapcore.NewTee(
		zapcore.NewCore(errorEncoder, errorFile, errorPriority),
		zapcore.NewCore(errorConsoleEncoder, errorConsole, errorPriority),
		zapcore.NewCore(accessEncoder, accessFile, accessPriority),
		zapcore.NewCore(accessConsoleEncoder, accessConsole, accessPriority),
	)

	opts := []zap.Option{zap.AddCaller(), zap.AddStacktrace(errorPriority)}
	logger := zap.New(core, opts...)
	defer logger.Sync()
	logger.Info("constructed a logger")
	return logger, nil
}

func ProdLogger() (*zap.Logger, error) {
	return prodLogger()
}

func globalLogger(status int, elapsed string, ip string, path string, method string) (*zap.Logger, error) {
	InitialFields := map[string]interface{}{
		"status":    status,
		"elapsed":   elapsed,
		"remote ip": ip,
		"path":      path,
		"method":    method,
	}
	opts := []zap.Option{}
	fs := []zapcore.Field{}
	keys := []string{}
	for k := range InitialFields {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fs = append(fs, zap.Any(k, InitialFields[k]))
	}
	opts = append(opts, zap.Fields(fs...))

	errorPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return status >= 400
	})

	accessPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return status < 400
	})

	allToConsolePriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return true
	})
	allToConsoleFile := zapcore.AddSync(os.Stdout)

	errorFile, _, errorErr := zap.Open([]string{configs.ErrorLogPath}...)
	if errorErr != nil {
		panic(errorErr)
	}
	accessFile, _, accessErr := zap.Open([]string{configs.AccessLogPath}...)
	if accessErr != nil {
		panic(accessErr)
	}

	errorEncoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     readableTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}
	accessEncoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     readableTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	errorEncoder := zapcore.NewJSONEncoder(errorEncoderConfig)
	accessEncoder := zapcore.NewJSONEncoder(accessEncoderConfig)
	allToConsoleEncoder := zapcore.NewConsoleEncoder(accessEncoderConfig)

	core := zapcore.NewTee(
		zapcore.NewCore(errorEncoder, errorFile, errorPriority),
		zapcore.NewCore(accessEncoder, accessFile, accessPriority),
		zapcore.NewCore(allToConsoleEncoder, allToConsoleFile, allToConsolePriority),
	)

	opts = append(opts, []zap.Option{zap.AddCaller(), zap.AddStacktrace(errorPriority)}...)
	logger := zap.New(core, opts...)
	defer logger.Sync()
	return logger, nil
}
func readableTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000 Z0700"))
}

func GlobalLogger(status int, elapsed string, ip string, path string, method string) (*zap.Logger, error) {
	return globalLogger(status, elapsed, ip, path, method)
}
