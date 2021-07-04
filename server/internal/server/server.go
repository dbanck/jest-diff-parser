package server

import (
	"context"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/creachadair/jrpc2"
	"github.com/creachadair/jrpc2/channel"
	"github.com/creachadair/jrpc2/handler"
	"github.com/dbanck/jest-diff-parser/internal/protocol"
	"github.com/dbanck/jest-diff-parser/internal/transformer"
)

type server struct {
	srvCtx      context.Context
	logger      *log.Logger
	jrpcOptions *jrpc2.ServerOptions
}

func NewServer(srvCtx context.Context) *server {
	opts := &jrpc2.ServerOptions{
		AllowPush: true,
	}

	return &server{
		srvCtx:      srvCtx,
		logger:      log.New(ioutil.Discard, "", 0),
		jrpcOptions: opts,
	}
}

func (srv *server) SetLogger(logger *log.Logger) {
	srv.jrpcOptions.Logger = logger
	srv.logger = logger
}

func (srv *server) startServer(reader io.Reader, writer io.WriteCloser) *jrpc2.Server {
	rpcSrv := jrpc2.NewServer(handler.Map{
		"initialize": handler.New(func(ctx context.Context, req *jrpc2.Request) (interface{}, error) {
			return protocol.InitializeResult{
				Capabilities: protocol.ServerCapabilities{},
			}, nil
		}),
		"initialized": handler.New(func(ctx context.Context, req *jrpc2.Request) (interface{}, error) {
			return nil, nil
		}),
		"shutdown": handler.New(func(ctx context.Context, req *jrpc2.Request) (interface{}, error) {
			return nil, nil
		}),
		"textDocument/codeLens": handler.New(func(ctx context.Context, req *jrpc2.Request) (interface{}, error) {
			return []protocol.CodeLens{}, nil
		}),
		"textDocument/documentLink": handler.New(func(ctx context.Context, req *jrpc2.Request) (interface{}, error) {
			return nil, nil
		}),
		"workspace/executeCommand": handler.New(func(ctx context.Context, cmd protocol.ExecuteCommandParams) (interface{}, error) {
			// TODO: maybe implement workspace/applyEdit
			return nil, nil
		}),
		"$/formatJestDiff": handler.New(func(ctx context.Context, args []string) (string, error) {
			return transformer.Transform(args[0]), nil
		}),
	}, srv.jrpcOptions)

	rpcSrv.Start(channel.LSP(reader, writer))

	return rpcSrv
}

func (srv *server) StartAndWait(reader io.Reader, writer io.WriteCloser) {
	rpcSrv := srv.startServer(reader, writer)
	srv.logger.Printf("Starting server with pid %d", os.Getpid())

	// Wrap waiter with a context so that we can cancel it here
	// after the service is cancelled (and srv.Wait returns)
	ctx, cancelFunc := context.WithCancel(srv.srvCtx)
	go func() {
		rpcSrv.Wait()
		cancelFunc()
	}()

	select {
	case <-ctx.Done():
		srv.logger.Printf("Stopping server with pid %d", os.Getpid())
		rpcSrv.Stop()
	}

	srv.logger.Printf("Server with pid %d stopped.", os.Getpid())
}
