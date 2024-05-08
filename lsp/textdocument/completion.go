package textdocument

import (
	"github.com/dgethings/lsp-cisco-ios/lsp/ios"
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func Completion(ctx *glsp.Context, params *protocol.CompletionParams) (interface{}, error) {
	return ios.Completions(), nil
}
