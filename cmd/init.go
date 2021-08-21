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
		fmt.Printf("%v", *token.TokenDB)
		if err != nil {
			log.Fatalf("%s\n", err)
		}
		fmt.Println(RefreshRootToken)

		if RefreshRootToken {
			err := token.AddRootToken(true)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("hehr")
			err := token.AddRootToken(false)
			if err != nil {
				fmt.Println(err)
			}
		}
	},
}
