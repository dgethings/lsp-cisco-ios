package textdocument

import (
	"bytes"
	"encoding/gob"
	"log"
	"os"

	"github.com/dgethings/lsp-cisco-ios/ios"
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func Completion(ctx *glsp.Context, params *protocol.CompletionParams) (interface{}, error) {
	var items []protocol.CompletionItem

	b, err := os.ReadFile("ios/keywords.buf")
	if err != nil {
		log.Fatal(err)
	}
	buf := bytes.NewBuffer(b)
	dec := gob.NewDecoder(buf)
	if err := dec.Decode(&ios.Keywords); err != nil {
		log.Fatal(err)
	}

	for kw, comps := range ios.Keywords {
		for _, comp := range comps {
			items = append(items, protocol.CompletionItem{
				Label:      kw,
				InsertText: &comp,
			})
		}
	}
	return items, nil
}
