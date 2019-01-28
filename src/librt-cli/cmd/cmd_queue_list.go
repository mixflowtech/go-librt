package cmd

import (
	"fmt"
	"github.com/mixflowtech/go-librt/libqueue"
	"github.com/mixflowtech/go-librt/logger"
	"github.com/spf13/cobra"
	"os"
	"runtime"
)

// TODO: move global flags to ...
var g_offset, g_count uint64

func init() {
	// logSendCmd.Flags().BoolVarP(&g_flagDebug, "debug", "d", false, "debug mode (default is off)")
	// logSendCmd.Flags().StringVarP(&g_logLevel, "level", "l", "info", "log level of the message, in { debug, info, notice, warn, error, critical, fatal }, (default is info)")
	// logSendCmd.Flags().StringVarP(&g_logTo, "output", "o", "", "log to file (defaults to stderr)")
	queListCmd.Flags().Uint64VarP(&g_offset, "offset", "l", 0, "Where to start")
	queListCmd.Flags().Uint64VarP(&g_count, "count", "n", 1000, "Number item(s) listed.( default 1000 )")
}

var queListCmd = &cobra.Command{
	Use:   "list [[Queue Name]] [[Data PATH]]",
	Short: "List Entry(s) in Queue",
	Long:  ``,
	Args:  cobra.RangeArgs(0, 2),
	Run: func(cmd *cobra.Command, args []string) {
		//message := args[0]
		log := logger.New("librt.cli.log_send")
		log.Configure("", true, "")

		// get folder

		var TestFolder string
		// var logFolder string
		tmp := normalizeFilePath(os.TempDir())
		if runtime.GOOS == "windows" {
			TestFolder = tmp + "queue\\"
			// logFolder = tmp + "queue\\log\\"
		} else {
			TestFolder = tmp + "queue/"
			// logFolder = tmp + "queue/log/"
		}
		_ = os.Mkdir(TestFolder, 766)
		// end get folder

		// begin write
		opt := queue.DefaultQueueOptions
		opt.InputTimeOut = 0
		// TODO: pass log to -> ... , when log porting finished.

		// Give a TestWorkerFactory{}
		q, err := queue.CreateQueue("Test", TestFolder+"bblot", log, &opt)
		if err != nil {
			log.Fatalf("Cannot create storage: %s", err)
		}
		fmt.Printf("total (history) count %d\n", q.Count())
		//q.Insert([]byte("dkljfsdalkfjsd;aj"))
		q.Fetch(g_offset, g_count, func(buf []byte) error {
			fmt.Printf("- %s.\n", buf)
			return nil
		});

		q.Close()
	},
}
