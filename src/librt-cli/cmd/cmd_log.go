package cmd

import (
	"github.com/spf13/cobra"
	// "github.com/mixflowtech/go-librt/logger"
)

func init() {
	// execCmd.Flags().StringVarP(&output, "output", "o", "", "output to file or directory (default is stdout)")
	logCmd.AddCommand(logSendCmd)
	logCmd.AddCommand(logServeCmd)
}

var logCmd = &cobra.Command{
	Use:   "log [SubCommand]",
	Short: "Log subsystem",
	Long:  ``,
	// Args:  cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {
		/*
		name := args[0]

		var path string
		if len(args) == 2 {
			path = args[1]
		}
		*/
		cmd.Help()
	},
}