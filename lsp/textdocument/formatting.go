package textdocument

import (
	"strings"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

func Format(text string) []protocol.TextEdit {
	var edits []protocol.TextEdit
	lines := strings.Split(text, "\n")
	var formattedLines []string

	indentationLevel := 0
	for _, line := range lines {
		trimmedLine := strings.TrimSpace(line)
		if trimmedLine == "" {
			formattedLines = append(formattedLines, "")
			continue
		}

		// Simple logic for indentation (can be expanded)
		if strings.HasPrefix(trimmedLine, "no ") || strings.HasPrefix(trimmedLine, "exit") || strings.HasPrefix(trimmedLine, "end") {
			if indentationLevel > 0 {
				indentationLevel--
			}
		}

		formattedLine := strings.Repeat("  ", indentationLevel) + trimmedLine
		formattedLines = append(formattedLines, formattedLine)

		// Increase indentation for certain commands (e.g., entering a sub-mode)
		if strings.HasSuffix(trimmedLine, "{") || strings.HasSuffix(trimmedLine, ")") || strings.HasSuffix(trimmedLine, "config") {
			// This is a very simplistic heuristic and needs to be improved
			indentationLevel++
		}
	}

	newText := strings.Join(formattedLines, "\n")

	// Create a single edit that replaces the entire document content
	edits = append(edits, protocol.TextEdit{
		Range: protocol.Range{
			Start: protocol.Position{Line: 0, Character: 0},
			End:   protocol.Position{Line: uint32(len(lines)), Character: 0},
		},
		NewText: newText,
	})

	return edits
}

