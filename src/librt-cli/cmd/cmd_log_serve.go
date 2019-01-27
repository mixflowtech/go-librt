package cmd

import (
	"github.com/spf13/cobra"
	"github.com/mixflowtech/go-librt/logger"
	"github.com/mixflowtech/go-librt/engine"
	"github.com/mixflowtech/go-librt/engine/impl"
)

var g_laddr string

func init() {
	// TODO: add flag `debug`...
	logServeCmd.Flags().StringVarP(&g_laddr, "listen", "l", ":13399", "Address listen on (default is `0.0.0.0:13399` )")
}

var logServeCmd = &cobra.Command{
	Use:   "serve",
	Short: "Execute JavaScript in Clout",
	Long:  ``,
	Args:  cobra.MaximumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		log := logger.New("librt.cli.log_serve")
		// TODO: reuse flag with log.send
		log.Configure("", true, "")
		srv, err := impl.NewLog1Server()
		if err != nil {
			log.Fatalf("%v", err)
		}
		d, err := engine.NewDaemonNoCertTCP(srv, g_laddr)
		if err != nil {
			log.Fatalf("%v", err)
		}
		log.Info("Log Server listening at %s", g_laddr)
		err = d.AcceptLoop()
		//wait for shutdown to be complete
		d.WaitForShutdown()
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