package cmd

import (
	_ "github.com/sikalabs/redirect-server/cmd/all"
	"github.com/sikalabs/redirect-server/cmd/root"
	_ "github.com/sikalabs/redirect-server/cmd/version"
	"github.com/spf13/cobra"
)

func Execute() {
	cobra.CheckErr(root.Cmd.Execute())
}
