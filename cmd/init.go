package cmd

import (
	"fmt"
	"log"

	"github.com/braior/devops-api/common"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// 注册命令
func init() {
	DBPath = viper.GetString("database.dbPath")
	initCmd.PersistentFlags().BoolVar(&RefreshRootToken, "refresh-root-token", false, "refresh root token")
	RootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Short: "init",
	Use:   "init",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("run user update command")
		token, err := common.NewToken()
		if err != nil {
			log.Fatalf("%s\n", err)
		}

		if RefreshRootToken {
			err := token.AddRootToken(true)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			err := token.AddRootToken(false)
			if err != nil {
				fmt.Println(err)
			}
		}
	},
}
