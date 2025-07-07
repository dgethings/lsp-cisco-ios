package server

import (
	"fmt"
	"runtime/debug"

	"github.com/dgethings/lsp-cisco-ios/lsp/ios"
	"github.com/dgethings/lsp-cisco-ios/lsp/textdocument"
	"github.com/tliron/commonlog"
	"github.com/tliron/commonlog/slog"
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
	defer func() {
		if r := recover(); r != nil {
			logger.Errorf("recovered in f: %v", r)
			logger.Errorf("stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	commonlog.Configure(2, &path)

	// Load keywords from JSON file
	err := ios.LoadKeywords("ios/commands.json")
	if err != nil {
		logger.Errorf("Failed to load keywords: %v", err)
		return
	}

	handler = protocol.Handler{
		Initialize:             initialize,
		Initialized:            initialized,
		Shutdown:               shutdown,
		SetTrace:               setTrace,
		TextDocumentCompletion: textdocument.Completion,
		TextDocumentDidOpen:    textdocument.DidOpen,
		TextDocumentDidChange:  textdocument.DidChange,
		TextDocumentHover:      textdocument.Hover,
		TextDocumentFormatting: textdocument.Formatting,
		TextDocumentDocumentSymbol: textdocument.DocumentSymbol,
		WorkspaceDidChangeConfiguration: didChangeConfiguration,
	}

	server := server.NewServer(&handler, lsName, true)
	server.RunStdio()
}

func initialize(ctx *glsp.Context, params *protocol.InitializeParams) (any, error) {
	capabilities := handler.CreateServerCapabilities()
	capabilities.TextDocumentSync = protocol.TextDocumentSyncKindFull
	capabilities.PublishDiagnostics = protocol.PublishDiagnosticsOptions{ /* Add any specific options here if needed */ }
	capabilities.DocumentFormattingProvider = true
	capabilities.DocumentSymbolProvider = true
	capabilities.Workspace = &protocol.WorkspaceServerCapabilities{
		Configuration: true,
	}
	logger.Debugf("InitializeParams: %+v", params)
	logger.Debug("InitializeCapabilities", "TextDocumentSync", fmt.Sprintf("%v", capabilities.TextDocumentSync))

	return protocol.InitializeResult{
		Capabilities: capabilities,
		ServerInfo: &protocol.InitializeResultServerInfo{
			Name:    lsName,
			Version: &version,
		},
	}, nil
}

func initialized(ctx *glsp.Context, params *protocol.InitializedParams) error {
	logger.Debug("Session initialized")
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

func didChangeConfiguration(ctx *glsp.Context, params *protocol.DidChangeConfigurationParams) error {
	logger.Debugf("DidChangeConfiguration: %+v", params)

	// Extract the target IOS version from the configuration
	if settings, ok := params.Settings.(map[string]interface{}); ok {
		if iosLspSettings, ok := settings["ios-lsp"].(map[string]interface{}); ok {
			if targetVersion, ok := iosLspSettings["targetIOSVersion"].(string); ok {
				textdocument.SetTargetIOSVersion(targetVersion)
				logger.Debugf("Set target IOS version to: %s", targetVersion)
			}
			if targetDeviceType, ok := iosLspSettings["targetDeviceType"].(string); ok {
				textdocument.SetTargetDeviceType(targetDeviceType)
				logger.Debugf("Set target device type to: %s", targetDeviceType)
			}
		}
	}
	return nil
}
