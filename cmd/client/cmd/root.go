package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "client",
	Short: "CLI client",
	Long:  `CLI client`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.HelpFunc()(cmd, args)
	},
}

func init() {

	VersionCmd.Flags().BoolP("version", "V", false, "version of client")
	LoginCmd.Flags().BoolP("username", "U", true, "username")
	TestCmd.Flags().StringP("username", "U", "", "username")

}
