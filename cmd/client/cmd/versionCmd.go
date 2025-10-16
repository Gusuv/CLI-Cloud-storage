package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	version string = "unknown"
	commit  string = "unknown"
	build   string = "unknown"
)

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "version of client",
	Long:  `version of ccs client`,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Printf("Version: %s\nCommit: %s\nBuildDate: %s\n", version, commit, build)
	},
}
