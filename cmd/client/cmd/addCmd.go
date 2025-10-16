package cmd

import "github.com/spf13/cobra"

var AddCmd = &cobra.Command{
	Use:   "add",
	Short: "add a resource",
	Long:  `add a resource in cloud_storage`,
	RunE:  func(cmd *cobra.Command, args []string) error { return nil },
}
