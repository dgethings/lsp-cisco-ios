package ios

import (
	"encoding/json"
	"os"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

var Keywords []Keyword

type Keyword struct {
	Keyword       string    `json:"keyword"`
	Documentation string    `json:"documentation"`	
	Insert        string    `json:"insert"`
	Mode          string    `json:"mode"`
	MinVersion    string    `json:"min_version"`
	MaxVersion    string    `json:"max_version"`
	DeviceTypes   []string  `json:"device_types"`
	Children      []Keyword `json:"children"`
}

func LoadKeywords(path string) error {
	file, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	return json.Unmarshal(file, &Keywords)
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

func FindKeyword(name string) *Keyword {
	for i := range Keywords {
		if Keywords[i].Keyword == name {
			return &Keywords[i]
		}
	}
	return nil
}

func GetChildrenCompletions(k *Keyword) []protocol.CompletionItem {
	var c []protocol.CompletionItem
	for _, child := range k.Children {
		c = append(c, child.Completion())
	}
	return c
}
