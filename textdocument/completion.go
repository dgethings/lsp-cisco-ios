package textdocument

import (
	"github.com/dgethings/lsp-cisco-ios/ios"
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func Completion(ctx *glsp.Context, params *protocol.CompletionParams) (interface{}, error) {
	var items []protocol.CompletionItem

	for token, nextToken := range ios.Enable {
		next := nextToken
		items = append(items, protocol.CompletionItem{
			Label:      token,
			Detail:     &next,
			InsertText: &next,
		})
	}
	return items, nil
}
