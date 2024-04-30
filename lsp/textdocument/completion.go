package textdocument

import (
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func Completion(ctx *glsp.Context, params *protocol.CompletionParams) (interface{}, error) {
	snippet := protocol.CompletionItemKindSnippet
	items := []protocol.CompletionItem{
		{
			Label:         "enable password ${1:15}",
			Documentation: "To set a local password to control access to various privilege levels, use the enable password command in global configuration mode. To remove the password requirement, use the no form of this command.",
			Kind:          &snippet,
		},
	}
	return items, nil
}
