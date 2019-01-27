package impl

import (
	"fmt"
	"github.com/keybase/go-framed-msgpack-rpc/rpc"
	"golang.org/x/net/context"

	"github.com/mixflowtech/go-librt/engine/protocol/mxengine1"
)

type Log1Server struct {
	//s3store    *s3test.OfficialS3Store
	//randSource *rand.Rand
	//buf        []byte
}

func NewLog1Server() (*Log1Server, error) {
	/*
	s3store, err := s3test.NewOfficialS3Store("us-east-1", *bucketPrefix, true)
	if err != nil {
		return nil, err
	}
	*/

	return &Log1Server{
	}, nil
}

func (f *Log1Server) CreateServerAndRegister(rpcs *rpc.Server) {
	rpcs.Register(mxengine1.LogProtocol(f))
}

func (f *Log1Server) Shutdown() {
}

// func (f *Log1Server) Get(ctx context.Context, arg s3test.GetArg) (res s3test.GetRes, err error) {
func (f *Log1Server) DoLog(ctx context.Context, arg mxengine1.DoLogArg) (err error) {
	fmt.Printf("got message %s\n", arg.Message)
	/*
	if arg.Size > 0 && len(f.buf) < arg.Size {
		f.buf = make([]byte, arg.Size)
		f.randSource.Read(f.buf)
	}
	if arg.Size > 0 {
		res.Value = f.buf[0:arg.Size]
		return res, nil
	}

	buf, err := f.s3store.Get(arg.Key)
	if err != nil {
		log.Fatal(err)
	}
	res.Value = buf
	*/
	return nil
}
