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
	if err != nil {
		return &h, err
	}
	h.Contents = word.Documentation
	return &h, nil
}

// Finds the word surrounding the given line and char number
func selectedWord(contents string, lineNum int, colNum int) (ios.Keyword, error) {
	logger.Debug("params", "line", lineNum, "column", colNum)
	line := contentAtLine(contents, lineNum)
	logger.Debugf("Line-Content %s", line)
	if colNum > len(line) {
		return ios.Keyword{}, nil
	}
	if string(line[colNum-1]) == " " {
		return ios.Keyword{}, nil
	}
	word := getSelectedWord(line, colNum)
	logger.Debugf("Word-Under-Cursor %s", word)
	keywords := getKeyword(word)
	logger.Noticef("Keyword-Count %d", len(keywords))
	if len(keywords) == 0 {
		return ios.Keyword{}, nil
	}
	if len(keywords) == 1 {
		return keywords[0], nil
	}

	// there is more than one match. Find longest unique match and return that
	return longestMatch(line, keywords), nil
	// count := 100
	// for len(keywords) > 1 {
	// 	count--
	// 	if count == 0 {
	// 		logger.Noticef("No-Keyword failed to find unique keyword after 100 tries")
	// 		return ios.Keyword{}, errors.New("failed to find unique keyword after 100 tries")
	// 	}
	// 	newColNum := nextSpaceIndex(line, colNum) + 1
	// 	if newColNum == colNum {
	// 		msg := fmt.Sprintf("Cannot find unique keyword for %s", word)
	// 		logger.Noticef("No-Keyword Cannot find unique keyword for %s", word)
	// 		return ios.Keyword{}, errors.New(msg)
	// 	}
	// 	nextWord := getSelectedWord(line, newColNum)
	// 	word = fmt.Sprintf("%s %s", word, nextWord)
	// 	keywords = getKeyword(word)
	// }
	// logger.Debugf("Remaining-Keyword-Count %d", len(keywords))
	// // return keywords[0], nil
	// return ios.Keyword{}, nil
}

func longestMatch(line string, keywords []ios.Keyword) ios.Keyword {
	matchIdx := 0
	maxCount := 0
	for i, k := range keywords {
		matchCount := 0
		for j := 0; j < len(k.Keyword); j++ {
			if j >= len(line) {
				break
			}
			if k.Keyword[j] != line[j] {
				break
			}
			matchCount++
		}
		if matchCount > maxCount {
			maxCount = matchCount
			matchIdx = i
		}
	}
	return keywords[matchIdx]
}

func nextSpaceIndex(line string, colNum int) int {
	idx := strings.Index(line[colNum:], " ")
	if idx == -1 {
		return colNum
	}
	return colNum + idx + 1
}

func contentAtLine(contents string, lineNum int) string {
	lines := strings.Split(contents, "\n")
	logger.Debug("contentAtLine", "lines", len(lines), "lineNum", lineNum)
	if lineNum > len(lines) {
		return ""
	}
	logger.Debugf("Matching-Line %s", lines[lineNum])
	return lines[lineNum-1]
}

func getKeyword(word string) []ios.Keyword {
	var keywords []ios.Keyword
	for kw, obj := range ios.Keywords() {
		if strings.Contains(kw, word) || kw == word {
			keywords = append(keywords, obj)
		}
	}
	return keywords
}

// from the given file, line number and column number, find the surrounding word
func getSelectedWord(line string, colNum int) string {
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
