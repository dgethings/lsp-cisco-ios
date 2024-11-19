package textdocument

import (
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func Hover(ctx *glsp.Context, params *protocol.HoverParams) (*protocol.Hover, error) {
	h := protocol.Hover{}
	word, err := selectedWord(
		State[params.TextDocument.URI],
		int(params.Position.Line),
		int(params.Position.Character),
	)
	logger.Debugf("Selected-Keyword %s", word.Keyword)
	if err != nil {
		logger.Debugf("Selected-Keyword-Error %v", err)
		return &h, err
	}
	h.Contents = word.Documentation
	logger.Debugf("Selected-Keyword-Documentation %s", h.Contents)
	return &h, nil
}
