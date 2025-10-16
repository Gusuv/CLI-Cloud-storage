package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var TestCmd = &cobra.Command{
	Use: "test",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Usage()
		fmt.Println(args[1])
	},
}
