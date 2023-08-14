package config

import "github.com/spf13/viper"

func Load(fname string) {
	viper.SetConfigFile(fname)
	viper.SetConfigType("yaml")

}
