package root

import (
	"github.com/sikalabs/redirect-server/version"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "redirect-server",
	Short: "redirect-server, " + version.Version,
}
