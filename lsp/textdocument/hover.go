package textdocument

import (
	"strings"

	"github.com/dgethings/lsp-cisco-ios/lsp/ios"
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

// Finds the word surrounding the given line and char number
func selectedWord(contents string, lineNum int, colNum int) (ios.Keyword, error) {
	logger.Debug("params", "line", lineNum, "column", colNum)
	line := contentAtLine(contents, lineNum)
	logger.Debugf("Line-Content %s", line)

	lineWords := strings.Fields(line)
	var bestMatch ios.Keyword
	maxMatch := 0

	for _, keyword := range ios.Keywords() {
		keywordWords := strings.Fields(keyword.Keyword)
		matchCount := 0
		for i := 0; i < len(lineWords) && i < len(keywordWords); i++ {
			if lineWords[i] == keywordWords[i] {
				matchCount++
			} else {
				break
			}
		}

		if matchCount > 0 && matchCount == len(keywordWords) && matchCount > maxMatch {
			// Check if cursor is within the matched keyword part on the line
			matchedPart := strings.Join(lineWords[:matchCount], " ")
			if colNum <= len(matchedPart) {
				maxMatch = matchCount
				bestMatch = keyword
			}
		}
	}

	if maxMatch > 0 {
		logger.Debugf("Matching-Keyword %s", bestMatch.Keyword)
		return bestMatch, nil
	}

	return ios.Keyword{}, nil
}

func contentAtLine(contents string, lineNum int) string {
	lines := strings.Split(contents, "\n")
	if lineNum >= len(lines) {
		return ""
	}
	logger.Debugf("Matching-Line %s", lines[lineNum])
	return lines[lineNum]
}
