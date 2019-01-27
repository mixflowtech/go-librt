package cmd

import (
	context "golang.org/x/net/context"
	"log"
	"time"

	"github.com/keybase/go-framed-msgpack-rpc/rpc"
	"github.com/mixflowtech/go-librt/engine/protocol/mxengine1"
	"github.com/mixflowtech/go-librt/logger"
	"github.com/spf13/cobra"
)

type RPCConnectHandler struct {
}

func (c *RPCConnectHandler) OnConnect(ctx context.Context,
	conn *rpc.Connection, client rpc.GenericClient, _ *rpc.Server) error {
	log.Printf("OnConnect\n")
	return nil
}

func (c *RPCConnectHandler) OnDoCommandError(err error, nextTime time.Duration) {
	log.Printf("OnDoCommandError\n")
}

func (c *RPCConnectHandler) OnDisconnected(ctx context.Context, status rpc.DisconnectStatus) {
	log.Printf("OnDisconnected\n")
}

func (c *RPCConnectHandler) OnConnectError(err error, dur time.Duration) {
	log.Printf("OnConnectError\n")
}

func (c *RPCConnectHandler) ShouldRetry(name string, err error) bool {
	log.Printf("ShouldRetry\n")
	return true
}

func (c *RPCConnectHandler) ShouldRetryOnConnect(err error) bool {
	log.Printf("ShouldRetryOnConnect\n")
	return true
}

func (c *RPCConnectHandler) HandlerName() string {
	return "RPCConnectHandler"
}


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
	Use:   "send [message] [log server, fmprpc://]",
	Short: "Send log message to server",
	Long:  ``,
	Args:  cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {
		message := args[0]

		log := logger.New("librt.cli.log_send")
		log.Configure("", g_flagDebug, g_logTo)

		if len(args) == 2 {
			var targetHost = args[1]
			uri, err := rpc.ParseFMPURI(targetHost)
			if err != nil {
				log.Fatalf("%v", err)
			}

			opts := rpc.ConnectionOpts{
				DontConnectNow: true,
			}

			logOpts := rpc.NewStandardLogOptions("", nil)
			trans := rpc.NewConnectionTransport(uri,
				rpc.NewSimpleLogFactory(nil, logOpts),
				nil, rpc.DefaultMaxFrameLength)
			// handler := &RPCConnectHandler{}

			// FIXME: add NewTLSConnection
			// conn := rpc.NewTLSConnection(rpc.NewFixedRemote(*srvAddr), cert, nil,
			// 	handler, rpc.NewSimpleLogFactory(nil, logOpts), nullLogOutput{}, rpc.DefaultMaxFrameLength, opts)

			conn := rpc.NewConnectionWithTransport(&RPCConnectHandler{}, trans, nil,
				logger.LogOutputWithDepthAdder{Logger: logger.New("demo.client")}, opts)

			client := mxengine1.LogClient{Cli: conn.GetClient()}

			// call DoLog
			arg := mxengine1.DoLogArg{
				SessionID: 1001,
				Message: message,
			}
			_ = client.DoLog(context.Background(), arg)
			// check err

		} else {
			// local mode.
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