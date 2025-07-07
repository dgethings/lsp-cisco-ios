package textdocument

import (
	"fmt"
	"strings"

	"github.com/dgethings/lsp-cisco-ios/lsp/ios"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

// targetIOSVersion will be set from LSP client configuration
var targetIOSVersion string = ""

func SetTargetIOSVersion(version string) {
	targetIOSVersion = version
}

func Diagnose(text string) []protocol.Diagnostic {
	var diagnostics []protocol.Diagnostic
	lines := strings.Split(text, "\n")

	for lineNum, line := range lines {
		words := strings.Fields(line)
		if len(words) > 0 {
			command := words[0]
			found := false
			var matchedKeyword *ios.Keyword

			for _, keyword := range ios.Keywords {
				if keyword.Keyword == command {
					found = true
					matchedKeyword = &keyword
					break
				}
			}

			if !found {
				diagnostics = append(diagnostics, protocol.Diagnostic{
					Range: protocol.Range{
						Start: protocol.Position{Line: uint32(lineNum), Character: 0},
						End:   protocol.Position{Line: uint32(lineNum), Character: uint32(len(line))},
					},
					Severity: protocol.DiagnosticSeverityError,
					Source:   "ios-lsp",
					Message:  fmt.Sprintf("Unknown or invalid Cisco IOS command: %s", command),
				})
			} else if targetIOSVersion != "" {
				// Basic version validation (can be expanded for more complex version comparisons)
				if matchedKeyword.MinVersion != "" && compareVersions(targetIOSVersion, matchedKeyword.MinVersion) < 0 {
					diagnostics = append(diagnostics, protocol.Diagnostic{
						Range: protocol.Range{
							Start: protocol.Position{Line: uint32(lineNum), Character: 0},
							End:   protocol.Position{Line: uint32(lineNum), Character: uint32(len(line))},
						},
						Severity: protocol.DiagnosticSeverityWarning,
						Source:   "ios-lsp",
						Message:  fmt.Sprintf("Command '%s' requires IOS version %s or later (current: %s)", command, matchedKeyword.MinVersion, targetIOSVersion),
					})
				} else if matchedKeyword.MaxVersion != "" && compareVersions(targetIOSVersion, matchedKeyword.MaxVersion) > 0 {
					diagnostics = append(diagnostics, protocol.Diagnostic{
						Range: protocol.Range{
							Start: protocol.Position{Line: uint32(lineNum), Character: 0},
							End:   protocol.Position{Line: uint32(lineNum), Character: uint32(len(line))},
						},
						Severity: protocol.DiagnosticSeverityWarning,
						Source:   "ios-lsp",
						Message:  fmt.Sprintf("Command '%s' is deprecated or not available in IOS version %s (last known: %s)", command, targetIOSVersion, matchedKeyword.MaxVersion),
					})
				}
			}
		}
	}

	return diagnostics
}

// compareVersions compares two version strings (e.g., "15.2", "12.4(T)").
// Returns -1 if v1 < v2, 0 if v1 == v2, 1 if v1 > v2.
// This is a simplistic comparison and might need a more robust library for production.
func compareVersions(v1, v2 string) int {
	// Remove any non-numeric/dot characters for basic comparison
	v1 = strings.Split(v1, "(")[0]
	v2 = strings.Split(v2, "(")[0]

	p1 := strings.Split(v1, ".")
	p2 := strings.Split(v2, ".")

	maxLength := len(p1)
	if len(p2) > maxLength {
		maxLength = len(p2)
	}

	for i := 0; i < maxLength; i++ {
		n1, n2 := 0, 0
		if i < len(p1) {
			fmt.Sscanf(p1[i], "%d", &n1)
		}
		if i < len(p2) {
			fmt.Sscanf(p2[i], "%d", &n2)
		}

		if n1 < n2 {
			return -1
		} else if n1 > n2 {
			return 1
		}
	}
	return 0
}
