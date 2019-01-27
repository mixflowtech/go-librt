package cmd

import (
	"github.com/mixflowtech/go-librt/logger"
	queue "github.com/mixflowtech/go-librt/libqueue"
	"github.com/spf13/cobra"
	"os"
	"runtime"
)

// TODO: move global flags to ...
func init() {
	// logSendCmd.Flags().BoolVarP(&g_flagDebug, "debug", "d", false, "debug mode (default is off)")
	// logSendCmd.Flags().StringVarP(&g_logLevel, "level", "l", "info", "log level of the message, in { debug, info, notice, warn, error, critical, fatal }, (default is info)")
	// logSendCmd.Flags().StringVarP(&g_logTo, "output", "o", "", "log to file (defaults to stderr)")
}

func normalizeFilePath(Path string) string {
	a := []rune(Path)
	if a[len(a)-1] == rune(os.PathSeparator) {
		return Path
	}
	return Path + string(os.PathSeparator)
}

var quePushCmd = &cobra.Command{
	Use:   "push [message] [[Queue Name]] [[Data PATH]]",
	Short: "push data to disk-queue",
	Long:  ``,
	Args:  cobra.RangeArgs(1, 3),
	Run: func(cmd *cobra.Command, args []string) {
		message := args[0]

		log := logger.New("librt.cli.queue_push")
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
		q, err := queue.CreateQueue("Test", TestFolder, nil, nil, &opt)
		if err != nil {
			log.Fatalf("Cannot create storage: %s", err)
		}

		/*
		tmp := make([]byte, 50000)
		for i := 0; i < 2; i++ {
			saved := q.Insert(tmp)
			if !saved {
				log.Fatalf("Cannot insert date")
			}
		}
		*/
		saved := q.Insert([]byte(message))
		if !saved {
			log.Fatalf("Cannot insert date")
		}
		q.Close()
	},
}
