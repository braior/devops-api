package cmd

import (
	"log"

	"github.com/astaxie/beego"
	"github.com/braior/devops-api/common"
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
	return updateTokenCmd
}
