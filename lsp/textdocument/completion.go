package textdocument

import (
	"strings"

	"github.com/dgethings/lsp-cisco-ios/lsp/ios"
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func Completion(ctx *glsp.Context, params *protocol.CompletionParams) (interface{}, error) {
	var completions []protocol.CompletionItem
	text := State[params.TextDocument.URI]
	lines := strings.Split(text, "\n")
	line := lines[params.Position.Line]
	parts := strings.Split(strings.TrimSpace(line), " ")

	switch len(parts) {
	case 1:
		completions = ios.Completions()
	default:
		keyword := ios.FindKeyword(parts[0])
		if keyword != nil {
			completions = ios.GetChildrenCompletions(keyword)
		}
	}

	return completions, nil
}
