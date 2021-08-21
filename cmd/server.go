package cmd

import (
	"log"

	"github.com/astaxie/beego"
	"github.com/braior/brtool"
	"github.com/spf13/cobra"
)

// 注册命令
func init() {
	// serverCmd.PersistentFlags().StringP("server", "server", "server name", "oo")
	serverCmd.PersistentFlags().StringVar(&RunMode, "runmode", "dev", "author name for copyright attribution")
	serverCmd.PersistentFlags().BoolVar(&Debug, "debug", false, "author name for copyright attribution")
	serverCmd.PersistentFlags().StringVar(&LogPathFromCli, "log", "", "author name for copyright attribution")
	serverCmd.PersistentFlags().StringVar(&CfgFile, "config", "", "author name for copyright attribution")
	RootCmd.AddCommand(serverCmd)
}

var serverCmd = &cobra.Command{
	Short: "server",
	Use:   "server ",
	// Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if RunMode != "" {
			if _, ok := brtool.InstringSlice([]string{"dev", "test", "prod"}, RunMode); !ok {
				log.Fatalln("get run mode input error, mode: dev|test|prod")
			}
		}
		beego.BConfig.RunMode = RunMode

		beego.SetStaticPath("/api/static/download/qr", "static/download/qr")
		beego.Run()
	},
}
