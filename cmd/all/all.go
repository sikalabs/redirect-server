package server

import (
	"log"

	"github.com/sikalabs/redirect-server/cmd/root"
	"github.com/sikalabs/redirect-server/pkg/server"
	"github.com/spf13/cobra"
)

var FlagTarget string
var PreservePath bool

var Cmd = &cobra.Command{
	Use:     "all",
	Short:   "Redirect all to ...",
	Aliases: []string{"s"},
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		err := server.Server(FlagTarget, PreservePath)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	root.Cmd.AddCommand(Cmd)
	Cmd.PersistentFlags().StringVarP(
		&FlagTarget,
		"target",
		"t",
		"",
		"Redirect target (eg.: https://google.com)",
	)
	Cmd.MarkPersistentFlagRequired("target")
	Cmd.PersistentFlags().BoolVarP(
		&PreservePath,
		"preserve-path",
		"p",
		false,
		"Preserve path",
	)
}
