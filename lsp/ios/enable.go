package ios

import protocol "github.com/tliron/glsp/protocol_3_16"

type Keyword struct {
	Keyword       string
	Documentation string
	Insert        string
}

func NewKeyword(k string, d string, i string) Keyword {
	return Keyword{Keyword: k, Documentation: d, Insert: i}
}

func (k *Keyword) Completion() protocol.CompletionItem {
	kind := protocol.CompletionItemKindKeyword
	format := protocol.InsertTextFormatSnippet
	c := protocol.CompletionItem{
		Label: k.Keyword,
		Documentation: protocol.MarkupContent{
			Kind:  protocol.MarkupKindMarkdown,
			Value: k.Documentation,
		},
	}
	c.InsertText = &k.Insert
	c.Kind = &kind
	c.InsertTextFormat = &format
	return c
}

var (
	EnablePassword = NewKeyword("enable password", "To set a local password to control access to various privilege levels, use the enable password command in global configuration mode. To remove the password requirement, use the no form of this command.\n\nenable password _level_ {password | _encryption-type_ encrypted-password}\nno enable password _level_\n", "enable password ${1:15} ${2} ${3:5} ${4}")
	EnableSecret   = NewKeyword("enable secret", "To specify an additional layer of security over the enable password command, use the enable secret command in global configuration mode. To turn off the enable secret function, use the no form of this command.\n\nenable secret _level_ { 0 _unencrypted-password_ | _encryption-type_ _encrypted-password_}\nno enable secret _level_ _encryption-type encrypted-password_", "enable secret ${1:15} ${2:0} ${3}")
)

func Completions() []protocol.CompletionItem {
	return []protocol.CompletionItem{
		EnablePassword.Completion(),
		EnableSecret.Completion(),
	}
}

func Keywords() map[string]Keyword {
	k := make(map[string]Keyword)
	k[EnablePassword.Keyword] = EnablePassword
	k[EnableSecret.Keyword] = EnableSecret
	return k
}
