package cmd

import (
	"github.com/astaxie/beego/logs"
	"github.com/spf13/cobra"
)

var create = &cobra.Command{
	Use:   "create",
	Short: "create",
	Long:  "create Command.",
}

func NewCreateTokenCmd() *cobra.Command {
	var createCmd = &cobra.Command{
		Use:   "token",
		Short: "token",
		Long:  "create Command is generate a token for user.",
		Run: func(cmd *cobra.Command, args []string) {
			// 如果没有输入 name
			var token *Token
			var err error

			if token, err = NewToken(); err != nil {
				logs.Error("new %s failed, err: %s", userName, err)
				return
			}

			if userName == "root" {
				err = token.AddRootToken()
				if err != nil {
					logs.Error("%s\n", err)
				}
			} else {
				userToken, err := token.GetToken(userName)
				if err != nil {
					logs.Error("err: %s", err)
					return
				}
				if userToken == nil {
					err = token.AddToken(rootToken, userName)
					if err != nil {
						logs.Error("%s\n", err)
					}
				} else {
					logs.Error("%s's token is already exist", userName)
				}
			}

		},
	}

	// note：使用子命令形式，下列可在help中展开
	// 命令参数，保存的值，参数名，默认参数，说明
	createCmd.Flags().StringVarP(&userName, "user", "u", "", "set the test mode")
	createCmd.Flags().StringVarP(&rootToken, "root-token", "t", "", "set the test mode")
	createCmd.MarkFlagRequired("user")
	// createCmd.MarkFlagRequired("root-token")

	return createCmd

}
