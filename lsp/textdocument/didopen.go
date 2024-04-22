package textdocument

import (
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func DidOpen(ctx *glsp.Context, params *protocol.DidOpenTextDocumentParams) error {
	logger.Info("DidOpen", "params", params)
	State[params.TextDocument.URI] = params.TextDocument.Text
	return nil
}
