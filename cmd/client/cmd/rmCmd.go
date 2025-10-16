package cmd

import "github.com/spf13/cobra"

var RmCmd = &cobra.Command{
	Use:   "rm",
	Short: "delete a resource",
	Long:  `delete a resource in cloud_storage`,
	Run:   func(cmd *cobra.Command, args []string) {},
}
