package textdocument

import (
	"strings"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

func DocumentSymbol(text string) []protocol.DocumentSymbol {
	var symbols []protocol.DocumentSymbol
	lines := strings.Split(text, "\n")

	for i, line := range lines {
		trimmedLine := strings.TrimSpace(line)
		if strings.HasPrefix(trimmedLine, "interface") || strings.HasPrefix(trimmedLine, "router") || strings.HasPrefix(trimmedLine, "line") {
			symbol := protocol.DocumentSymbol{
				Name:           trimmedLine,
				Detail:         "",
				Kind:           protocol.SymbolKindClass, // Or appropriate symbol kind
				Range:          protocol.Range{Start: protocol.Position{Line: uint32(i), Character: 0}, End: protocol.Position{Line: uint32(i), Character: uint32(len(line))}},
				SelectionRange: protocol.Range{Start: protocol.Position{Line: uint32(i), Character: 0}, End: protocol.Position{Line: uint32(i), Character: uint32(len(line))}},
			}
			symbols = append(symbols, symbol)
		}
	}

	return symbols
}
