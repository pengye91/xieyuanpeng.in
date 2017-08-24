package log

import (
	"go.uber.org/zap"
	"os"
)

var (
	DevCfg = zap.NewDevelopmentConfig()
	ErrCfg = zap.NewDevelopmentConfig()
	//ProdCfg = zap.NewProductionConfig()
)

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
		os.Getenv("MYGOPATH") + "/xieyuanpeng.in/logs/sys_err.log",
	}
	return ErrCfg.Build()
}
