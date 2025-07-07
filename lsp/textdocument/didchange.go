package textdocument

import (
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func DidChange(ctx *glsp.Context, params *protocol.DidChangeTextDocumentParams) error {
	logger.Info("DidChange", "params", params)
	for _, change := range params.ContentChanges {
		State[params.TextDocument.URI] = change.(protocol.TextDocumentContentChangeEvent).Text
	}

	// Send diagnostics
	diagnostics := Diagnose(State[params.TextDocument.URI])
	ctx.Notify(protocol.ServerTextDocumentPublishDiagnostics, protocol.PublishDiagnosticsParams{
		URI:         params.TextDocument.URI,
		Diagnostics: diagnostics,
	})

	return nil
}
