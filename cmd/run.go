package cmd

import (
	"fmt"
	"os"

	"github.com/astaxie/beego"
	"github.com/braior/brtool"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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

		DBPath = viper.GetString("database.dbPath")
		QrImageDir = viper.GetString("twoStepAuth.qrImageDir")
		if !brtool.IsExist(QrImageDir) {
			os.MkdirAll(QrImageDir, os.ModePerm)
		}
		EnableToken = viper.GetBool("security.enableToken")

		UploadPath = viper.GetString("app.uploadDir")
		if !brtool.IsExist(UploadPath) {
			os.MkdirAll(UploadPath, os.ModePerm)
		}
		if EnableToken {
			token, err := NewToken()
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
