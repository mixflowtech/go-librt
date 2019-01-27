package cmd

import (
	"github.com/spf13/cobra"
	//"github.com/mixflowtech/go-librt/logger"
)

func init() {
	// execCmd.Flags().StringVarP(&output, "output", "o", "", "output to file or directory (default is stdout)")
}

var logServeCmd = &cobra.Command{
	Use:   "serve [COMMAND or JavaScript PATH or URL] [[Clout PATH or URL]]",
	Short: "Execute JavaScript in Clout",
	Long:  ``,
	Args:  cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {
		/*
		name := args[0]

		var path string
		if len(args) == 2 {
			path = args[1]
		}
		*/
	},
}
/*
func Exec(name string, sourceCode string, c *clout.Clout) error {
	program, err := js.GetProgram(name, sourceCode)
	if err != nil {
		return err
	}

	jsContext := js.NewContext(name, log, common.Quiet, ardFormat, output)
	_, runtime := jsContext.NewCloutContext(c)
	_, err = runtime.RunProgram(program)

	return js.UnwrapError(err)
}
*/