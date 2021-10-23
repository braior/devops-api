package cmd

import (
	"log"

	"github.com/braior/devops-api/common"
	"github.com/spf13/cobra"
)

var delete = &cobra.Command{
	Use:   "delete",
	Short: "delete",
	Long:  "delete Command.",
}



func NewDeleteTokenCmd() *cobra.Command {
	var deleteTokenCmd = &cobra.Command{
		Use:   "token",
		Short: "token",
		Long:  "create Command.",

		Run: func(cmd *cobra.Command, args []string) {
			var token *common.Token

			err := token.DeleteToken(rootToken, userName)
			if err != nil {
				log.Fatalf("%s\n", err)
			}
		},
	}

	// note：使用子命令形式，下列可在help中展开
	// 命令参数，保存的值，参数名，默认参数，说明
	deleteTokenCmd.Flags().StringVarP(&userName, "user", "u", "", "set the test mode")
	deleteTokenCmd.Flags().StringVarP(&rootToken, "root-token", "t", "", "set the test mode")
	deleteTokenCmd.MarkFlagRequired("user")
	deleteTokenCmd.MarkFlagRequired("root-token")
	return deleteTokenCmd
}
