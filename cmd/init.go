package cmd

import (
	"os"

	"github.com/braior/devops-api/utils"
	"github.com/spf13/cobra"
)

var (

	// UploadPath 上传目录
	UploadPath string

	RunMode string

	// QrImageDir 二维码图片目录
	QrImageDir string

	// Used for flags.
	CfgFile string

	RefreshRootToken bool

	// DBPath 数据库文件路径
	DBPath string

	InitRootToken bool
	Name          string
	EnableToken   bool
	Debug         bool

	userName  string
	rootToken string
)

var RootCmd = &cobra.Command{
	Use:   "devops-api",
	Short: "A generator for Cobra based Applications",
	Long: `Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
}

func init() {
	cobra.OnInitialize(initConfig)

	// token action cmd
	create.AddCommand(NewCreateTokenCmd())
	delete.AddCommand(NewDeleteTokenCmd())
	update.AddCommand(NewUpdateTokenCmd())
	get.AddCommand(NewGetTokenCmd())

	// user action cmd
	get.AddCommand(NewGetTokenNameListCmd())

	RootCmd.AddCommand(create)
	RootCmd.AddCommand(delete)
	RootCmd.AddCommand(update)
	RootCmd.AddCommand(get)

	run.PersistentFlags().StringVar(&RunMode, "mode", "dev", "author name for copyright attribution")
	run.PersistentFlags().BoolVar(&utils.Debug, "debug", false, "author name for copyright attribution")
	run.PersistentFlags().StringVar(&utils.LogPathFromCli, "log", "", "author name for copyright attribution")
	run.PersistentFlags().StringVar(&CfgFile, "config", "", "author name for copyright attribution")

	RootCmd.AddCommand(run)
}

// Execute executes the root command.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
