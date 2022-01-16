package cmd

import (
	"os"

	"github.com/astaxie/beego/logs"
	"github.com/braior/devops-api/utils"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func initConfig() {
	if CfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(CfgFile)
	} else {
		// Find home directory.
		// home, err := os.UserHomeDir()
		home, err := os.Getwd()
		cobra.CheckErr(err)

		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(home + "/conf")
		viper.SetConfigType("yaml")
		viper.SetConfigName("devops-api")
	}

	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err == nil {
		logs.Info("Using config file: %s", viper.ConfigFileUsed())
	} else {
		logs.Error("%s", err)
	}

	viper.OnConfigChange(func(e fsnotify.Event) {
		logs.Info("Config file changed:", e.Name)
	})
	viper.WatchConfig()

	utils.LogInit()
}
