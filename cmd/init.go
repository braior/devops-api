package cmd

import (
	"log"
	"os"

	"github.com/astaxie/beego"
	"github.com/braior/brtool"
	"github.com/braior/devops-api/common"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (

	// UploadPath 上传目录
	UploadPath string

	// Used for flags.
	CfgFile string

	RefreshRootToken bool

	RunMode string

	// DBPath 数据库文件路径
	DBPath string

	InitRootToken bool
	Name          string
)

func init() {
	DBPath = viper.GetString("database.dbPath")
	// 上传目录时候存在
	UploadPath = viper.GetString("app.uploadDir")
	if !brtool.IsExist(UploadPath) {
		os.MkdirAll(UploadPath, os.ModePerm)
	}

	serverInit.AddCommand(NewInitRootCmd())
	RootCmd.AddCommand(serverInit)

}

var serverInit = &cobra.Command{
	Use:   "init",
	Short: "init root token",
	Long:  "init root token",
}

func NewInitRootCmd() *cobra.Command {
	var initRootCmd = &cobra.Command{
		Use:   "root-rooken",
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

// // 注册命令
// func init() {
// 	DBPath = viper.GetString("database.dbPath")
// 	initCmd.PersistentFlags().BoolVar(&RefreshRootToken, "refresh-root-token", false, "refresh root token")
// 	RootCmd.AddCommand(initCmd)
// }

// var initCmd = &cobra.Command{
// 	Short: "init",
// 	Use:   "init",
// 	Run: func(cmd *cobra.Command, args []string) {
// 		fmt.Println("run user update command")
// 		token, err := common.NewToken()
// 		if err != nil {
// 			log.Fatalf("%s\n", err)
// 		}

// 		if RefreshRootToken {
// 			err := token.AddRootToken(true)
// 			if err != nil {
// 				fmt.Println(err)
// 			}
// 		} else {
// 			err := token.AddRootToken(false)
// 			if err != nil {
// 				fmt.Println(err)
// 			}
// 		}
// 	},
// }
