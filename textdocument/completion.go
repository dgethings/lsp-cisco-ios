package textdocument

import (
	"github.com/dgethings/lsp-cisco-ios/ios"
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func Completion(ctx *glsp.Context, params *protocol.CompletionParams) (interface{}, error) {
	var items []protocol.CompletionItem

	for _, token := range ios.Keywords {
		items = append(items, protocol.CompletionItem{
			Label: token,
		})
	}
	return items, nil
}
