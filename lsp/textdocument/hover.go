package textdocument

import (
	"log/slog"
	"strings"

	"github.com/dgethings/lsp-cisco-ios/lsp/ios"
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func Hover(ctx *glsp.Context, params *protocol.HoverParams) (*protocol.Hover, error) {
	slog.Info("Hover", "params", params)
	word := selectedWord(State[params.TextDocument.URI], uint(params.Position.Line), uint(params.Position.Character))
	h := protocol.Hover{
		Contents: word.Documentation,
	}
	return &h, nil
}

// Finds the word surrounding the given line and char number
func selectedWord(contents string, lineNum uint, colNum uint) ios.Keyword {
	line := contentAtLine(contents, lineNum)
	if string(line[colNum]) == " " {
		return ios.Keyword{}
	}
	word := getSelectedWord(line, colNum)
	keywords := getKeyword(word)
	return keywords[0]
}

func contentAtLine(contents string, lineNum uint) string {
	return strings.Split(contents, "\n")[lineNum]
}

func getKeyword(word string) []ios.Keyword {
	var keywords []ios.Keyword
	for _, kw := range ios.Keywords() {
		if strings.Contains(kw.Keyword, word) {
			keywords = append(keywords, kw)
		}
	}
	return keywords
}

// from the given file, line number and column number, find the surrounding word
func getSelectedWord(line string, colNum uint) string {
	word := line
	start := strings.LastIndex(line[:colNum], " ")
	end := strings.Index(line[colNum+1:], " ")
	if start > 0 && end > 0 {
		word = line[start : int(colNum)+end+1]
	} else if start == -1 && end > 0 {
		word = line[0 : int(colNum)+end+1]
	} else if start > 0 && end == -1 {
		word = line[start:]
	}
	slog.Info("found word", "word", word)
	return word
}
