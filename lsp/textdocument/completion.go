package textdocument

import (
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func Completion(ctx *glsp.Context, params *protocol.CompletionParams) (interface{}, error) {
	items := []protocol.CompletionItem{
		enablePassword,
		enableSecret,
	}
	return items, nil
}

func new(label string, doc string, insert string) protocol.CompletionItem {
	kind := protocol.CompletionItemKindKeyword
	format := protocol.InsertTextFormatSnippet
	c := protocol.CompletionItem{
		Label: label,
		Documentation: protocol.MarkupContent{
			Kind:  protocol.MarkupKindMarkdown,
			Value: doc,
		},
	}
	c.InsertText = &insert
	c.Kind = &kind
	c.InsertTextFormat = &format
	return c
}

var (
	enablePassword = new("enable password", "To set a local password to control access to various privilege levels, use the enable password command in global configuration mode. To remove the password requirement, use the no form of this command.\n\nenable password _level_ {password | _encryption-type_ encrypted-password}\nno enable password _level_\n", "enable password ${1:15} ${2} ${3:5} ${4}")
	enableSecret   = new("enable secret", "To specify an additional layer of security over the enable password command, use the enable secret command in global configuration mode. To turn off the enable secret function, use the no form of this command.\n\nenable secret _level_ { _0_ unencrypted-password | encryption-type encrypted-password}\nno enable secret _level_ _encryption-type encrypted-password_", "enable secret ${1:15} ${2:0} ${3}")
)
