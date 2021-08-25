package cmd

import (
	"log"

	"github.com/astaxie/beego"
	"github.com/braior/devops-api/common"
	"github.com/spf13/cobra"
)

// 注册命令
func init() {
	// tokenCmd.PersistentFlags().StringP("server", "server", "server name", "oo")
	tokenCmd.PersistentFlags().StringVar(&RootToken, "root-token", "", "Token Manage")

	tokenCmd.PersistentFlags().StringVar(&CreateToken, "create", "", "author name for copyright (attribution)")
	tokenCmd.MarkFlagRequired("create")
	tokenCmd.PersistentFlags().StringVar(&DeleteToken, "delete", "", "author name for copyright attribution")
	tokenCmd.MarkFlagRequired("delete")
	RootCmd.AddCommand(tokenCmd)
}

var tokenCmd = &cobra.Command{
	Short: "token",
	Use:   "token ",
	// Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var token *common.Token
		var err error
		if token, err = common.NewToken(); err != nil {
			beego.BeeLogger.Error("new  failed, err: %s", CreateToken, err)
		}

		if CreateToken != "" {
			err = token.AddToken(RootToken, CreateToken)
			if err != nil {
				log.Fatalf("%s\n", err)
			}
		}
		if DeleteToken != "" {
			err = token.DeleteToken(RootToken, DeleteToken)
			if err != nil {
				log.Fatalf("%s\n", err)
			}
		}
	},
}
