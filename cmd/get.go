package cmd

import (
	"fmt"
	"log"

	"github.com/astaxie/beego/logs"

	"github.com/spf13/cobra"
)

var get = &cobra.Command{
	Use:   "get",
	Short: "get",
	Long:  "get Command.",
}

func NewGetTokenCmd() *cobra.Command {
	var getCmd = &cobra.Command{
		Use:   "token",
		Short: "get user token",
		// Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			var token *Token

			var err error

			if token, err = NewToken(); err != nil {
				logs.Error("refresh root token failed, err: %s", err)
				return
			}

			userToken, err := token.GetToken(userName)
			if err != nil {
				log.Fatalf("%s\n", err)
			}
			fmt.Printf("the < %s > token is: \n\t%s\n", userName, userToken[userName])
		},
	}

	// note：使用子命令形式，下列可在help中展开
	// 命令参数，保存的值，参数名，默认参数，说明
	getCmd.Flags().StringVarP(&userName, "user", "u", "", "set the test mode")
	//getCmd.Flags().StringVarP(&rootToken, "root-token", "t", "", "set the test mode")
	getCmd.MarkFlagRequired("user")
	//getCmd.MarkFlagRequired("root-token")

	return getCmd
}

func NewGetTokenNameListCmd() *cobra.Command {
	var getCmd = &cobra.Command{
		Use:   "users",
		Short: "get exist users",
		// Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			var token *Token

			var err error

			if token, err = NewToken(); err != nil {
				logs.Error("init tokenDB err: %s", err)
				return
			}

			var users []string
			userInfo, err := token.GetUsers()
			if err != nil {
				log.Fatalf("%s\n", err)
			}
			for user := range userInfo {
				users = append(users, user)
			}
			fmt.Printf("the users is: \n\t%s\n", users)
		},
	}

	return getCmd
}
