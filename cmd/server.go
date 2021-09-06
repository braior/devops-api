package cmd

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/braior/brtool"
	"github.com/braior/devops-api/common"
	"github.com/braior/devops-api/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// 注册命令
func init() {
	// serverCmd.PersistentFlags().StringP("server", "server", "server name", "oo")
	serverCmd.PersistentFlags().StringVar(&RunMode, "runmode", "dev", "author name for copyright attribution")
	serverCmd.PersistentFlags().BoolVar(&utils.Debug, "debug", false, "author name for copyright attribution")
	serverCmd.PersistentFlags().StringVar(&utils.LogPathFromCli, "log", "", "author name for copyright attribution")
	serverCmd.PersistentFlags().StringVar(&CfgFile, "config", "", "author name for copyright attribution")
	RootCmd.AddCommand(serverCmd)
}

var serverCmd = &cobra.Command{
	Short: "server",
	Use:   "server ",
	// Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		EnableToken := viper.GetBool("security.enableToken")
		if EnableToken {
			token, err := common.NewToken()
			if err != nil {
				errLog := fmt.Sprintf("init token error: %v", err)
				beego.BeeLogger.Error(errLog)
				return
			}
			r, _ := token.IsExistToken("root")
			if !r {
				beego.BeeLogger.Error("root token not exist, please init")
				// return
			}
		}
		if RunMode != "" {
			if _, ok := brtool.InstringSlice([]string{"dev", "test", "prod"}, RunMode); !ok {
				beego.BeeLogger.Error("get run mode input error, mode: dev|test|prod")
				//log.Fatalf("get run mode input error, mode: dev|test|prod")
				// log.Fatalln("get run mode input error, mode: dev|test|prod")
			}

		}
		beego.BConfig.RunMode = RunMode

		beego.SetStaticPath("/api/static/download/qr", "static/download/qr")
		beego.Run()
	},
}
