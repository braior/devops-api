package cmd

import (
	"fmt"
	"log"

	"github.com/astaxie/beego"
	"github.com/braior/devops-api/common"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	userName  string
	rootToken string
)

// 注册命令
func init() {
	DBPath = viper.GetString("database.dbPath")

	//tokenCmd.PersistentFlags().BoolVar(&InitRootToken, "init-root-token", false, "init root token")
	//tokenCmd.PersistentFlags().BoolVar(&RefreshRootToken, "refresh-root-token", false, "refresh root token")

	create.AddCommand(NewInitRootCmd())
	create.AddCommand(NewRefreshRootCmd())

	create.AddCommand(NewCreateCmd())
	//tokenCmd.AddCommand(NewGetCmd())
	delete.AddCommand(NewDeleteCmd())
	refresh.AddCommand(NewRefreshRootCmd())
	get.AddCommand(NewGetCmd())
	RootCmd.AddCommand(create)
	RootCmd.AddCommand(get)
	RootCmd.AddCommand(delete)
	RootCmd.AddCommand(refresh)

}

var create = &cobra.Command{
	Use:   "create",
	Short: "create",
	Long:  "create Command.",
}

var delete = &cobra.Command{
	Use:   "delete",
	Short: "delete",
	Long:  "delete Command.",
}

var  refresh = &cobra.Command{
	Use:   "refresh-root",
	Short: "refresh",
	Long:  "refresh Command.",
}

var  get = &cobra.Command{
	Use:   "get",
	Short: "get",
	Long:  "get Command.",

}

func NewInitRootCmd() *cobra.Command {
	var initRootCmd = &cobra.Command{
		Use:   "init-root",
		Short: "init",
		Long:  "init root token Command.",

		Run: func(cmd *cobra.Command, args []string) {
			var token *common.Token
			var err error
			if token, err = common.NewToken(); err != nil {
				beego.BeeLogger.Error("new root token failed, err: %s", err)
				return
			}
			err = token.AddRootToken()
			if err != nil {
				log.Fatalf("%s\n", err)
			}
		},
	}

	return initRootCmd
}

func NewDeleteCmd() *cobra.Command {
	var deleteCmd = &cobra.Command{
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
	deleteCmd.Flags().StringVarP(&userName, "user", "u", "", "set the test mode")
	deleteCmd.Flags().StringVarP(&rootToken, "root-token", "t", "", "set the test mode")
	deleteCmd.MarkFlagRequired("user")
	deleteCmd.MarkFlagRequired("root-token")
	return deleteCmd
}

func NewCreateCmd() *cobra.Command {
	var createCmd = &cobra.Command{
		Use:   "token",
		Short: "token",
		Long:  "create Command is generate a token for user.",
		Run: func(cmd *cobra.Command, args []string) {
			// 如果没有输入 name
			var token *common.Token
			var err error

			if token, err = common.NewToken(); err != nil {
				beego.BeeLogger.Error("new %s failed, err: %s", userName, err)
				return
			}
			err = token.AddToken(rootToken, userName)
			if err != nil {
				log.Fatalf("%s\n", err)
			}
		},
	}

	// note：使用子命令形式，下列可在help中展开
	// 命令参数，保存的值，参数名，默认参数，说明
	createCmd.Flags().StringVarP(&userName, "user", "u", "", "set the test mode")
	createCmd.Flags().StringVarP(&rootToken, "root-token", "t", "", "set the test mode")
	createCmd.MarkFlagRequired("user")
	createCmd.MarkFlagRequired("root-token")

	return createCmd

}

func NewRefreshRootCmd() *cobra.Command {
	var refreshRootCmd = &cobra.Command{
		Use:   "token",
		Short: "token",
		// Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			var token *common.Token

			var err error

			if token, err = common.NewToken(); err != nil {
				beego.BeeLogger.Error("refresh root token failed, err: %s", err)
				return
			}

			err = token.ForceRefresh()
			if err != nil {
				log.Fatalf("%s\n", err)
			}
		},
	}
	return refreshRootCmd
}

func NewGetCmd() *cobra.Command {
	var getCmd = &cobra.Command{
		Use:   "token",
		Short: "get user token",
		// Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			var token *common.Token

			var err error

			if token, err = common.NewToken(); err != nil {
				beego.BeeLogger.Error("refresh root token failed, err: %s", err)
				return
			}

			userToken,err := token.GetToken(userName)
			if err != nil {
				log.Fatalf("%s\n", err)
			}
			fmt.Printf("the < %s > token is: %s",userName,userToken[userName])
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
