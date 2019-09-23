package txlogger

import (
	"fmt"
	"github.com/heirko/go-contrib/logrusHelper"
	"github.com/sirupsen/logrus"
	"github.com/txvier/base/txconfigs"
	_ "github.com/txvier/base/txlogger/hook"
	"os"
)

///////////////////////////////////////////////////////////////////////////////////////////////////
type TxLogger struct {
	*logrus.Logger
}

var txlogger TxLogger

func GetLogger() TxLogger {
	return txlogger
}

///////////////////////////////////////////////////////////////////////////////////////////////////
// the options will be invorked before read config
type Option func(l *logrus.Logger)

var REPORT_CALLER_OPTION Option = func(l *logrus.Logger) {
	l.SetReportCaller(true)
}

///////////////////////////////////////////////////////////////////////////////////////////////////
func InitLogger(lgLocation string, ops ...Option) (err error) {
	if err := txconfigs.SetConfigLocation(lgLocation); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	lg := logrus.New()

	for _, op := range ops {
		op(lg)
	}

	cfg := txconfigs.GetConfigs()

	// Read configuration
	//mate.RegisterWriter("rotatelogs", NewRotatelogsWriter)
	var c = logrusHelper.UnmarshalConfiguration(cfg.Viper)                    // Unmarshal configuration from Viper
	if err = logrusHelper.SetConfig(logrus.StandardLogger(), c); err != nil { // for e.g. apply it to logrus default instance
		return err
	}
	txlogger = TxLogger{lg}
	return
	// ### End Read Configuration
}

///////////////////////////////////////////////////////////////////////////////////////////////////
