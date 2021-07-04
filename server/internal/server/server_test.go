package server

import (
	"context"
	"io"
	"log"
	"os"
	"testing"

	"github.com/creachadair/jrpc2"
	"github.com/creachadair/jrpc2/channel"
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
	var msg string
	if err := rsp.UnmarshalResult(&msg); err != nil {
		log.Fatalf("Decoding result: %v", err)
	}

	want := "initialize"
	if msg != want {
		t.Errorf("server() = %v, want %v", msg, want)
	}
}

func testLogger(w io.Writer, prefix string) *log.Logger {
	return log.New(w, prefix, log.LstdFlags|log.Lshortfile)
}
