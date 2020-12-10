package config

import (
	"flag"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// Config type
type Config struct {
	Session string
	Dayth   int
	Part    int
}

// Load func
func Load() (*Config, error) {

	flag.Int("day", 1, "day of puzzle")
	flag.Int("part", 1, "1 or 2")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	viper.SetConfigName("config")
	viper.AddConfigPath("./config")
	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	var config Config
	config.Session = viper.GetString("COOKIE_SESSION")
	config.Dayth = viper.GetInt("day")
	config.Part = viper.GetInt("part")
	return &config, nil
}
