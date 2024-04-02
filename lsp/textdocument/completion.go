package textdocument

import (
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func Completion(ctx *glsp.Context, params *protocol.CompletionParams) (interface{}, error) {
	var items = []protocol.CompletionItem{

		protocol.CompletionItem{
			Label: "no activation-character",
		},
		protocol.CompletionItem{
			Label: "clear archive log config force persistent",
		},
		protocol.CompletionItem{
			Label: "no databits",
		},
		protocol.CompletionItem{
			Label: "no file privilege level {{ level }}",
		},
		protocol.CompletionItem{
			Label: "no length",
		},
		protocol.CompletionItem{
			Label: "monitor event-trace disable dump enable size stacktrace {{ component }}",
		},
		protocol.CompletionItem{
			Label: "no refuse-message",
		},
		protocol.CompletionItem{
			Label: "no slave auto-sync config",
		},
		protocol.CompletionItem{
			Label: "test flash",
		},
  }
  return items, nil
}

