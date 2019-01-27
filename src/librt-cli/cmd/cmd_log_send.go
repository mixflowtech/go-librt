package cmd

import (
	"github.com/mixflowtech/go-librt/logger"
	"github.com/spf13/cobra"
)

// TODO: move global flags to ...
var g_flagDebug bool
var g_logLevel string
var g_logTo string

func init() {
	logSendCmd.Flags().BoolVarP(&g_flagDebug, "debug", "d", false, "debug mode (default is off)")
	logSendCmd.Flags().StringVarP(&g_logLevel, "level", "l", "info", "log level of the message, in { debug, info, notice, warn, error, critical, fatal }, (default is info)")
	logSendCmd.Flags().StringVarP(&g_logTo, "output", "o", "", "log to file (defaults to stderr)")
}

var logSendCmd = &cobra.Command{
	Use:   "send [COMMAND or JavaScript PATH or URL] [[Clout PATH or URL]]",
	Short: "Execute JavaScript in Clout",
	Long:  ``,
	Args:  cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {
		message := args[0]
		/*

		if len(args) == 2 {
			target_host = args[1]
		}
		*/
		if len(args) == 2 {
			// var target_host = args[1]
		} else {
			log := logger.New("librt-cli.log.send")
			log.Configure("", g_flagDebug, g_logTo)
			// debug, info, notice, warn, error, critical, fatal
			// TODO: apply g_logLevel flag
			switch g_logLevel {
				case "debug":
					log.Debug(message)
				case "info":
					log.Info(message)
				// ....
				default:
					panic("TODO: unsupported log level")
			}
		}
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