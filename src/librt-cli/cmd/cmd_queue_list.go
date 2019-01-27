package cmd

import (
	"fmt"
	queue "github.com/mixflowtech/go-librt/libqueue"
	"github.com/mixflowtech/go-librt/logger"
	"github.com/spf13/cobra"
	"io"
	"os"
	"runtime"
	"sync/atomic"
)


type TestWorker struct {
	id queue.WorkerID
}

type TestWorkerFactory struct {
	id queue.WorkerID
}

func (n *TestWorkerFactory) CreateWorker() (queue.Worker, error) {
	return &TestWorker{
		id: queue.WorkerID(atomic.AddUint64((*uint64)(&n.id), 1) - 1),
	}, nil
}

func (n *TestWorkerFactory) CanCreateWorkers() bool {
	return true
}

func (n *TestWorkerFactory) NeedTimeoutProcessing() bool {
	return false
}

func (n *TestWorkerFactory) Close() {
}

func (n *TestWorker) ProcessMessage(msg *queue.QueueItem) int {
	start, _ := msg.Stream.Seek(0, io.SeekCurrent)
	size, _ := msg.Stream.Seek(0, io.SeekEnd)
	size -= start
	msg.Stream.Seek(start, io.SeekStart)
	buf := make([]byte, size)
	msg.Stream.Read(buf)
	if string(buf) != "error" {
		return queue.ProcessedSuccessful
	}
	return queue.ProcessedWithError
}

func (n *TestWorker) ProcessTimeout() int {
	return queue.ProcessedSuccessful
}

func (n *TestWorker) GetID() queue.WorkerID {
	return n.id
}

func (n *TestWorker) Close() {
}


// TODO: move global flags to ...

func init() {
	// logSendCmd.Flags().BoolVarP(&g_flagDebug, "debug", "d", false, "debug mode (default is off)")
	// logSendCmd.Flags().StringVarP(&g_logLevel, "level", "l", "info", "log level of the message, in { debug, info, notice, warn, error, critical, fatal }, (default is info)")
	// logSendCmd.Flags().StringVarP(&g_logTo, "output", "o", "", "log to file (defaults to stderr)")
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
		q, err := queue.CreateQueue("Test", TestFolder, nil, &TestWorkerFactory{}, &opt)
		if err != nil {
			log.Fatalf("Cannot create storage: %s", err)
		}
		fmt.Printf("count %d\n", q.Count())

		//q.Insert([]byte("dkljfsdalkfjsd;aj"))
	},
}
