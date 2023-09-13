package config

import (
	"os"
	"path"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Config struct {
	Token string `mapstructure:"token"`
}

func LoadConfig() *Config {
	viper.SetDefault("token", "")

	home, err := os.UserHomeDir()
	cobra.CheckErr(err)
	viper.AddConfigPath(home)
	viper.AddConfigPath(".")
	viper.SetConfigName("config")

	if fileExists("./config.yaml") || fileExists(path.Join(home, "config.yaml")) {
		err = viper.ReadInConfig()
		log.Infof("using config file: %s", viper.ConfigFileUsed())
		cobra.CheckErr(err)
	}

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	for _, key := range viper.AllKeys() {
		err = viper.BindEnv(key, strings.ToUpper(strings.ReplaceAll(key, ".", "_")))
		cobra.CheckErr(err)
	}

	conf := Config{}
	if err := viper.Unmarshal(&conf); err != nil {
		cobra.CheckErr(err)
	}

	return &conf
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
