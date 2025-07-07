package textdocument

import (
	"strings"

	"github.com/dgethings/lsp-cisco-ios/lsp/ios"
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

// targetDeviceType will be set from LSP client configuration
var targetDeviceType string = ""

func SetTargetDeviceType(deviceType string) {
	targetDeviceType = deviceType
}

func Completion(ctx *glsp.Context, params *protocol.CompletionParams) (interface{}, error) {
	var completions []protocol.CompletionItem
	text := State[params.TextDocument.URI]
	lines := strings.Split(text, "\n")
	currentLine := lines[params.Position.Line]
	parts := strings.Fields(strings.TrimSpace(currentLine))

	currentMode := determineMode(lines[:params.Position.Line])

	if len(parts) == 0 {
		// No input, suggest top-level commands or commands for the current mode
		for _, k := range ios.Keywords {
			if (k.Mode == "global" || k.Mode == currentMode) && (len(k.DeviceTypes) == 0 || contains(k.DeviceTypes, targetDeviceType)) {
				completions = append(completions, k.Completion())
			}
		}
	} else {
		// Find the most specific keyword that matches the current input
		var currentKeyword *ios.Keyword
		for i := len(parts); i > 0; i-- {
			prefix := strings.Join(parts[:i], " ")
			k := ios.FindKeyword(prefix)
			if k != nil {
				currentKeyword = k
				break
			}
		}

		if currentKeyword != nil {
			// Suggest children of the current keyword
			for _, child := range currentKeyword.Children {
				if len(child.DeviceTypes) == 0 || contains(child.DeviceTypes, targetDeviceType) {
					completions = append(completions, child.Completion())
				}
			}
		} else {
			// No specific keyword found, suggest top-level commands that match the prefix
			// and are relevant to the current mode and device type
			prefix := strings.Join(parts, " ")
			for _, k := range ios.Keywords {
				if (k.Mode == "global" || k.Mode == currentMode) && (len(k.DeviceTypes) == 0 || contains(k.DeviceTypes, targetDeviceType)) && strings.HasPrefix(k.Keyword, prefix) {
					completions = append(completions, k.Completion())
				}
			}
		}
	}

	return completions, nil
}

// determineMode attempts to figure out the current configuration mode
// by looking at the lines above the current position.
// This is a simplified implementation and can be made more robust.
func determineMode(lines []string) string {
	mode := "global" // Default mode

	for i := len(lines) - 1; i >= 0; i-- {
		line := strings.TrimSpace(lines[i])
		if strings.HasPrefix(line, "interface") {
			mode = "interface"
			break
		} else if strings.HasPrefix(line, "router bgp") {
			mode = "router bgp"
			break
		} else if strings.HasPrefix(line, "line") {
			mode = "line"
			break
		} else if strings.HasPrefix(line, "exit") || strings.HasPrefix(line, "end") {
			// Exiting a mode, so go back to global or previous mode (simplistic)
			mode = "global"
			break
		}
	}
	return mode
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

