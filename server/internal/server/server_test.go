package server

import (
	"context"
	"io"
	"log"
	"os"
	"testing"

	"github.com/creachadair/jrpc2"
	"github.com/creachadair/jrpc2/channel"
	"github.com/dbanck/jest-diff-parser/internal/protocol"
	"github.com/google/go-cmp/cmp"
)

func Test_server_initialize(t *testing.T) {
	srvStdin, clientStdout := io.Pipe()
	clientStdin, srvStdout := io.Pipe()

	ctx := context.Background()
	srv := NewServer(ctx)
	if testing.Verbose() {
		srv.SetLogger(testLogger(os.Stdout, "[SERVER] "))
	}

	rpcSrv := srv.startServer(srvStdin, srvStdout)

	go func() {
		rpcSrv.Wait()
	}()

	clientCh := channel.LSP(clientStdin, clientStdout)
	opts := &jrpc2.ClientOptions{}
	if testing.Verbose() {
		opts.Logger = testLogger(os.Stdout, "[CLIENT] ")
	}
	client := jrpc2.NewClient(clientCh, opts)

	rsp, err := client.Call(ctx, "initialize", nil)
	if err != nil {
		log.Fatalf("Call: %v", err)
	}
	var msg protocol.InitializeResult
	if err := rsp.UnmarshalResult(&msg); err != nil {
		log.Fatalf("Decoding result: %v", err)
	}

	want := protocol.InitializeResult{
		Capabilities: protocol.ServerCapabilities{},
	}
	if diff := cmp.Diff(want, msg); diff != "" {
		t.Errorf("initialize msg mismatch (-want +got):\n%s", diff)
	}
}

func testLogger(w io.Writer, prefix string) *log.Logger {
	return log.New(w, prefix, log.LstdFlags|log.Lshortfile)
}
