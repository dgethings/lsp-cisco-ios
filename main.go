package main

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"

	"github.com/dgethings/lsp-cisco-ios/jsonrpc"
)

func main() {
	log := getLogger("/tmp/debug.log")
	log.Info("server started")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(jsonrpc.Split)

	for scanner.Scan() {
		msg := scanner.Text()
		handleMessage(msg)
	}
}

func handleMessage(_ any) {
}

func getLogger(filename string) *slog.Logger {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		msg := fmt.Sprintf("unable to write to %s log file %v", filename, err)
		panic(msg)
	}
	logger := slog.New(slog.NewTextHandler(f, nil))
	slog.SetDefault(logger)
	return logger
}
