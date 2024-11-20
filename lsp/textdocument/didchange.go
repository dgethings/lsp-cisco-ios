package textdocument

import (
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func DidChange(ctx *glsp.Context, params *protocol.DidChangeTextDocumentParams) error {
	logger.Debug("DidChange", "params", params)
	logger.Debug("DidChange", "file", params.TextDocument.URI, "contents", State[params.TextDocument.URI])
	for _, change := range params.ContentChanges {
		State[params.TextDocument.URI] = change.(protocol.TextDocumentContentChangeEventWhole).Text
	}
	logger.Debug("DidChange", "file", params.TextDocument.URI, "contents", State[params.TextDocument.URI])
	return nil
}
