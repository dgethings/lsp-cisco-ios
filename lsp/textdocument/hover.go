package textdocument

import (
	"fmt"
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
		Contents: fmt.Sprintf("you hovered on the word %s", word),
	}
	return &h, nil
}

// Finds the word surrounding the given line and char number
func selectedWord(contents string, lineNum uint, colNum uint) string {
	slog.Info("in the func")
	line := strings.Split(contents, "\n")[lineNum]
	word := line
	slog.Info("setting default match", "word", word)
	// if the char is a space then return empty string
	if string(line[colNum]) == " " {
		return ""
	}
	// find the " " closest to the char index, that should be the start of the word
	start := strings.LastIndex(line[:colNum], " ")
	slog.Info("start", "index", start)
	end := strings.Index(line[colNum+1:], " ")
	slog.Info("end", "end", int(colNum)+end+1)
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

// return the matching keyword
func getKeyword(contents string, lineNum uint, colNum uint) ios.Keyword {
	word := getSelectedWord(contents, lineNum, colNum)
	var keywords []ios.Keyword
	for _, kw := range ios.Keywords() {
		if strings.Contains(kw.Keyword, word) {
			keywords = append(keywords, kw)
		}
	}
	if len(keywords) == 1 {
		return keywords[0]
	} else {
		return keywords[0]
		// return uniqueKeyword()
	}
}

// from the given file, line number and column number, find the surrounding word
func getSelectedWord(contents string, lineNum uint, colNum uint) string {
	return "foo"
}
