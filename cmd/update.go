package cmd

import (
	"github.com/astaxie/beego/logs"
	"github.com/spf13/cobra"
)

var update = &cobra.Command{
	Use:   "update",
	Short: "update token",
	Long:  "update user token",
}

func NewUpdateTokenCmd() *cobra.Command {
	var updateTokenCmd = &cobra.Command{
		Use:   "refresh-root-token",
		Short: "refresh root token",
		// Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			var token *Token

			var err error

			if token, err = NewToken(); err != nil {
				logs.Error("refresh root token failed, err: %s", err)
				return
			}

			err = token.ForceRefresh()
			if err != nil {
				logs.Error("%s\n", err)
			}
		},
	}
	return updateTokenCmd
}
