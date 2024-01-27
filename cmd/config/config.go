package config

import (
	"os"
	"path"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/amirsalarsafaei/Gitlab-Tele-Bot/internal/clients/telegram"
	"github.com/amirsalarsafaei/Gitlab-Tele-Bot/internal/git"
	"github.com/amirsalarsafaei/Gitlab-Tele-Bot/internal/merge"
)

type Config struct {
	Telegram *telegram.Config `mapstructure:"telegram"`

	Merge merge.MergeConfig `mapstructure:"merge"`

	Git git.GitConfig `mapstructure:"git"`

	Secret string `mapstructure:"secret"`
}

func LoadConfig() *Config {
	viper.SetDefault("telegram.token", "sometoken")
	viper.SetDefault("telegram.chat_id", 0)
	viper.SetDefault("telegram.thread_id", 0)

	viper.SetDefault("merge.diff.enabled", false)
	viper.SetDefault("git.apikey", "apikey")
	viper.SetDefault("git.gitlabhost", "example.gitlab.com")

	viper.SetDefault("secret", "")

	home, err := os.UserHomeDir()
	cobra.CheckErr(err)
	viper.AddConfigPath(home)
	viper.AddConfigPath(".")
	viper.SetConfigName("notifier-config")

	if fileExists("./notifier-config.yaml") || fileExists(path.Join(home, "notifier-config.yaml")) {
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
