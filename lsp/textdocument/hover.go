package textdocument

import (
	"errors"
	"fmt"
	"strings"

	"github.com/dgethings/lsp-cisco-ios/lsp/ios"
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func Hover(ctx *glsp.Context, params *protocol.HoverParams) (*protocol.Hover, error) {
	h := protocol.Hover{}
	word, err := selectedWord(
		State[params.TextDocument.URI],
		uint(params.Position.Line),
		uint(params.Position.Character),
	)
	if err != nil {
		return &h, err
	}
	h.Contents = word.Documentation
	return &h, nil
}

// Finds the word surrounding the given line and char number
func selectedWord(contents string, lineNum uint, colNum uint) (ios.Keyword, error) {
	line := contentAtLine(contents, lineNum)
	if string(line[colNum]) == " " {
		return ios.Keyword{}, nil
	}
	word := getSelectedWord(line, colNum)
	keywords := getKeyword(word)
	count := 100
	for len(keywords) > 1 {
		count--
		if count == 0 {
			return ios.Keyword{}, errors.New("failed to find unique keyword after 100 tries")
		}
		newColNum := nextSpaceIndex(line, colNum) + 1
		if newColNum == colNum {
			msg := fmt.Sprintf("Cannot find unique keyword for %s", word)
			return ios.Keyword{}, errors.New(msg)
		}
		nextWord := getSelectedWord(line, newColNum)
		word = fmt.Sprintf("%s %s", word, nextWord)
		keywords = getKeyword(word)
	}
	return keywords[0], nil
}

func nextSpaceIndex(line string, colNum uint) uint {
	idx := strings.Index(line[colNum:], " ")
	if idx == -1 {
		return colNum
	}
	return colNum + uint(idx)
}

func contentAtLine(contents string, lineNum uint) string {
	return strings.Split(contents, "\n")[lineNum]
}

func getKeyword(word string) []ios.Keyword {
	var keywords []ios.Keyword
	for _, kw := range ios.Keywords() {
		if strings.Contains(kw.Keyword, word) || kw.Keyword == word {
			keywords = append(keywords, kw)
		}
	}
	return keywords
}

// from the given file, line number and column number, find the surrounding word
func getSelectedWord(line string, colNum uint) string {
	var word string
	start := strings.LastIndex(line[:colNum], " ")
	end := strings.Index(line[colNum+1:], " ")
	// there are no spaces in the line so return the line
	if start == -1 && end == -1 {
		word = line
	}
	// word is at the start of the line
	if start == -1 && end > -1 {
		word = line[0 : int(colNum)+end+1]
	}
	// word is in the middle of the line
	if start > 0 && end > 0 {
		word = line[int(start)+1 : int(colNum)+end+1]
	}
	// word is at the end of the line
	if start > 0 && end == -1 {
		word = line[int(start)+1:]
	}
	return word
}
