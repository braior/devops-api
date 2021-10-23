package cmd

import (
	"os"

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
	cobra.OnInitialize(setConfig)
	DBPath = viper.GetString("database.dbPath")
	// 上传目录时候存在

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
	RootCmd.AddCommand(serverRunCmd)
}

// Execute executes the root command.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}


