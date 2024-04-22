package server

import (
	"fmt"

	"github.com/dgethings/lsp-cisco-ios/lsp/textdocument"
	"github.com/tliron/commonlog"
	_ "github.com/tliron/commonlog/slog"
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
	"github.com/tliron/glsp/server"
)

const lsName = "ios-lsp"

var (
	version string = "0.0.1"
	handler protocol.Handler
	path    string           = "/tmp/lsp.log"
	logger  commonlog.Logger = commonlog.GetLogger("")
)

func New() {
	commonlog.Configure(2, &path)
	handler = protocol.Handler{
		Initialize:             initialize,
		Initialized:            initialized,
		Shutdown:               shutdown,
		SetTrace:               setTrace,
		TextDocumentCompletion: textdocument.Completion,
		TextDocumentDidOpen:    textdocument.DidOpen,
		TextDocumentDidChange:  textdocument.DidChange,
	}

	server := server.NewServer(&handler, lsName, false)
	server.RunStdio()
}

func initialize(ctx *glsp.Context, params *protocol.InitializeParams) (any, error) {
	capabilities := handler.CreateServerCapabilities()
	logger.Infof("InitializeParams: %+v", params)
	// logger.Infof("InitializeCapabilities: %+v", capabilities)
	logger.Info("InitializeCapabilities", "TextDocumentSync", fmt.Sprintf("%s", capabilities.TextDocumentSync))

	return protocol.InitializeResult{
		Capabilities: capabilities,
		ServerInfo: &protocol.InitializeResultServerInfo{
			Name:    lsName,
			Version: &version,
		},
	}, nil
}

func initialized(ctx *glsp.Context, params *protocol.InitializedParams) error {
	logger.Info("Session initialized")
	return nil
}

func shutdown(ctx *glsp.Context) error {
	logger.Error("ios-lsp crashed")
	protocol.SetTraceValue(protocol.TraceValueOff)
	return nil
}

func setTrace(ctx *glsp.Context, params *protocol.SetTraceParams) error {
	protocol.SetTraceValue(params.Value)
	return nil
}
