package txconfigs

import (
	jww "github.com/spf13/jwalterweatherman"
	"github.com/spf13/viper"
)

type Configs struct {
	*viper.Viper
}

var cfgs *Configs

func GetConfigs() *Configs {
	return cfgs
}

// the options will be invorked before read config
type Option func(v *viper.Viper)

var AUTOMATIC_ENV_OPTION Option = func(v *viper.Viper) {
	v.AutomaticEnv()
}

var WATCH_CONFIG_OPTION Option = func(v *viper.Viper) {
	v.WatchConfig()
}

func SetConfigLocation(absFile string, ops ...Option) error {

	jww.SetStdoutThreshold(jww.LevelDebug)

	var vp = viper.New()

	vp.SetConfigFile(absFile)

	for _, op := range ops {
		op(vp)
	}

	if err := vp.ReadInConfig(); err != nil {
		return err
	}

	jww.INFO.Println("Using config file:", viper.ConfigFileUsed())

	cfgs = &Configs{
		vp,
	}
	return nil
}
