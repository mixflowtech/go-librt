package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	// execCmd.Flags().StringVarP(&output, "output", "o", "", "output to file or directory (default is stdout)")
	queueCmd.AddCommand(quePushCmd)
	queueCmd.AddCommand(queListCmd)
}

var queueCmd = &cobra.Command{
	Use:   "queue",
	Short: "Disk queue subsystem",
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