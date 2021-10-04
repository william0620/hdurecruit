package config

import (
	"github.com/spf13/viper"
	"path"
)

type Service struct {
	Config
}

var s *Service

func init() {
	s = new(Service)

	viper.AddConfigPath(path.Join("config"))
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&s.Config)
	if err != nil {
		panic(err)
	}
}
