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
		return &h, nil // Return nil hover, not an error
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

	lineParts := strings.Fields(line)
	var bestMatch ios.Keyword
	maxMatchLength := 0

	// Iterate through all keywords to find the best match
	for _, keyword := range ios.Keywords {
		keywordParts := strings.Fields(keyword.Keyword)
		currentMatchLength := 0

		// Check for a match from the beginning of the line
		for i := 0; i < len(lineParts) && i < len(keywordParts); i++ {
			if lineParts[i] == keywordParts[i] {
				currentMatchLength++
			} else {
				break
			}
		}

		// If we found a longer match, or a match of the same length that is more specific
		if currentMatchLength > maxMatchLength {
			// Ensure the cursor is within the matched part of the line
			matchedText := strings.Join(lineParts[:currentMatchLength], " ")
			if colNum <= len(matchedText) {
				maxMatchLength = currentMatchLength
				bestMatch = keyword
			}
		}
	}

	if maxMatchLength > 0 {
		logger.Debugf("Matching-Keyword %s", bestMatch.Keyword)
		return bestMatch, nil
	}

	return ios.Keyword{}, fmt.Errorf("no keyword found at position")
}

func contentAtLine(contents string, lineNum int) string {
	lines := strings.Split(contents, "\n")
	if lineNum >= len(lines) {
		return ""
	}
	logger.Debugf("Matching-Line %s", lines[lineNum])
	return lines[lineNum]
}
