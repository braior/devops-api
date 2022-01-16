package cmd

import (
	"fmt"
	"github.com/astaxie/beego/logs"

	"github.com/astaxie/beego"
	"github.com/braior/brtool"
	"github.com/spf13/cobra"
)

// 注册命令
// func init() {

// 	serverRunCmd.PersistentFlags().StringVar(&RunMode, "runmode", "dev", "author name for copyright attribution")
// 	serverRunCmd.PersistentFlags().BoolVar(&utils.Debug, "debug", false, "author name for copyright attribution")
// 	serverRunCmd.PersistentFlags().StringVar(&utils.LogPathFromCli, "log", "", "author name for copyright attribution")
// 	serverRunCmd.PersistentFlags().StringVar(&CfgFile, "config", "", "author name for copyright attribution")

// 	fmt.Println(utils.Debug)
// }

var run = &cobra.Command{
	Use:   "run",
	Short: "run",

	// Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		if EnableToken {
			token, err := NewToken()
			if err != nil {
				errLog := fmt.Sprintf("init token error: %v", err)
				logs.Error(errLog)
				return
			}
			r, _ := token.IsExistToken("root")
			if !r {
				logs.Error("root token not exist, please init")
				// return
			}
		}
		if RunMode != "" {
			if _, ok := brtool.InstringSlice([]string{"dev", "test", "prod"}, RunMode); !ok {
				logs.Error("get run mode input error, mode: dev|test|prod")
			}

		}
		beego.BConfig.RunMode = RunMode

		beego.SetStaticPath("/api/static/download/qr", "static/download/qr")
		beego.Run()
	},
}
