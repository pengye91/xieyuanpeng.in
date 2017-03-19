package utils

import (
	"go.uber.org/zap"
)

var (
	DevCfg = zap.NewDevelopmentConfig()
	ErrCfg = zap.NewDevelopmentConfig()
	//ProdCfg = zap.NewProductionConfig()
)

func MyDevLogger() (*zap.Logger, error) {
	DevCfg.OutputPaths = []string{
		"stdout",
		"/root/go/src/github.com/pengye91/xieyuanpeng.in/logs/dev.log",
	}
	// this is the logging system error output paths
	DevCfg.ErrorOutputPaths = []string{
		"stderr",
		"/root/go/src/github.com/pengye91/xieyuanpeng.in/logs/sys_err.log",
	}
	return DevCfg.Build()
}

func MyErrLogger() (*zap.Logger, error) {
	dyn := zap.NewAtomicLevel()
	dyn.SetLevel(zap.ErrorLevel)
	ErrCfg.Level = dyn
	ErrCfg.OutputPaths = []string{
		"stderr",
		"/root/go/src/github.com/pengye91/xieyuanpeng.in/logs/err.log",
	}
	// this is the logging system error output paths
	ErrCfg.ErrorOutputPaths = []string{
		"stderr",
		"/root/go/src/github.com/pengye91/xieyuanpeng.in/logs/sys_err.log",
	}
	return ErrCfg.Build()
}
