package main

import (
	"cloud_storage/cmd/client/cmd"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{Use: "ccs"}

	rootCmd.AddCommand(cmd.RootCmd, cmd.VersionCmd, cmd.RmCmd, cmd.AddCmd, cmd.RegisterCmd, cmd.LoginCmd, cmd.TestCmd)

	rootCmd.Execute()
}
