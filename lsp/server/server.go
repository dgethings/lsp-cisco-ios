package server

import (
	"github.com/dgethings/lsp-cisco-ios/lsp/textdocument"
	"github.com/tliron/commonlog"
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
	"github.com/tliron/glsp/server"
)

const lsName = "ios-lsp"

var (
	version string = "0.0.1"
	handler protocol.Handler
	path    string           = "lsp.log"
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
	}

	server := server.NewServer(&handler, lsName, false)

	logger.Info("starting server")
	server.RunStdio()
}

func initialize(ctx *glsp.Context, params *protocol.InitializeParams) (any, error) {
	capabilities := handler.CreateServerCapabilities()
	logger.Infof("InitializeParams: %+v", params)
	logger.Infof("InitializeCapabilities: %+v", capabilities)

	return protocol.InitializeResult{
		Capabilities: capabilities,
		ServerInfo: &protocol.InitializeResultServerInfo{
			Name:    lsName,
			Version: &version,
		},
	}, nil
}

func initialized(ctx *glsp.Context, params *protocol.InitializedParams) error {
	logger.Infof("InitializedParams: %+v", params)
	return nil
}

func shutdown(ctx *glsp.Context) error {
	protocol.SetTraceValue(protocol.TraceValueOff)
	return nil
}

func setTrace(ctx *glsp.Context, params *protocol.SetTraceParams) error {
	protocol.SetTraceValue(params.Value)
	return nil
}
