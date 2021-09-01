package cmd

import (
	"fmt"
	"os"

	"github.com/astaxie/beego"
	"github.com/braior/devops-api/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	cobra.OnInitialize(initConfig)
}

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
		beego.BeeLogger.Info("Using config file: %s", viper.ConfigFileUsed())
	}

	fmt.Println(viper.GetString("app.runMode"))
	utils.LogInit()
}
