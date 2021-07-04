package main

import (
	"context"
	"log"
	"os"

	"github.com/dbanck/jest-diff-parser/internal/server"
)

func main() {
	ctx := context.Background()
	logger := log.New(os.Stderr, "", log.LstdFlags|log.Lshortfile)

	srv := server.NewServer(ctx)
	srv.SetLogger(logger)

	srv.StartAndWait(os.Stdin, os.Stdout)
}
