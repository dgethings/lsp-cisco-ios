package textdocument

import (
	"fmt"
	"log/slog"
	"strings"

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
func selectedWord(f string, l uint, c uint) string {
	line := strings.Split(f, "\n")
	word := line[l]
	// if the char is a space then return empty string
	if string(line[l][c]) == " " {
		return ""
	}
	// find the " " closest to the char index, that should be the start of the word
	start := strings.LastIndex(line[l][:c], " ")
	slog.Info("start", "index", start)
	end := strings.Index(line[l][c+1:], " ")
	slog.Info("end", "end", int(c)+end+1)
	if start > 0 && end > 0 {
		word = line[l][start : int(c)+end+1]
	} else if start == -1 && end > 0 {
		word = line[l][0 : int(c)+end+1]
	} else if start > 0 && end == -1 {
		word = line[l][start:]
	}
	slog.Info("found word", "word", word)
	return word
}
