package textdocument

import (
	"github.com/dgethings/lsp-cisco-ios/lsp/ios"
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func Completion(ctx *glsp.Context, params *protocol.CompletionParams) (interface{}, error) {
	word, err := selectedWord(State[params.TextDocument.URI], int(params.Position.Line), int(params.Position.Character))
	if err != nil {
		logger.Error("Unknown keyword", "file", params.TextDocument.URI, "line", params.Position.Line, "char", params.Position.Character)
		return nil, err
	}
	return ios.Completions(word), nil
}
