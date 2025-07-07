/*
Copyright © 2024 David Gethings
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"regexp"
	"strings"

	"github.com/gocolly/colly"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ios_scraper",
	Short: "Tool to scraper Cisco IOS docs to populate IOS LSP with data",
	Long:  `This tool creates go files used by the ios-lsp. It extracts data from Cisco website.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	RunE: func(cmd *cobra.Command, args []string) error {
		kws, err := getKeywords()
		if err != nil {
			return err
		}

		encoder := json.NewEncoder(os.Stdout)
		encoder.SetIndent("", "  ")
		return encoder.Encode(kws)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	slog.SetLogLoggerLevel(slog.LevelDebug)
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	slog.SetLogLoggerLevel(slog.LevelDebug)
}

const url = "https://www.cisco.com/c/en/us/td/docs/ios-xml/ios/fundamentals/command/cf_command_ref.html"

var keywords []Keyword

type Keyword struct {
	Command     string         `json:"keyword"`
	Description string         `json:"documentation"`
	Syntax      string         `json:"syntax"`
	Defaults    string         `json:"defaults"`
	Mode        string         `json:"mode"`
	MinVersion  string         `json:"min_version"`
	MaxVersion  string         `json:"max_version"`
	History     CommandHistory `json:"history"`
	Usage       UsageGuideline `json:"usage"`
	Examples    Examples       `json:"examples"`
}

type CommandHistory struct {
	Release      string `json:"release"`
	Modification string `json:"modification"`
}

type UsageGuideline struct {
	Preamble string `json:"preamble"`	
	Note     string `json:"note"`
}

type Examples struct {
	Preamble string `json:"preamble"`
	Code     string `json:"code"`
}

func getKeywords() ([]Keyword, error) {
	c := colly.NewCollector(
		colly.CacheDir("./cache"),
	)

	c.OnHTML("ul#bookToc", func(h *colly.HTMLElement) {
		// iterate through akk the "book" sections
		h.ForEach("li", func(_ int, li *colly.HTMLElement) {
			// ignore the Introduction section, that doesn't contain any commands
			if li.ChildText("a") != "Introduction" {
				url := fmt.Sprintf("%s%s", "https://www.cisco.com", li.ChildAttr("a", "href"))
				kw := parseChapter(url)
				if len(kw) > 0 {
					keywords = append(keywords, kw...)
				}
			}
		})
	})

	c.OnRequest(func(r *colly.Request) {
		slog.Debug("Vist", "URL", r.URL)
	})

	c.OnError(func(r *colly.Response, err error) {
		slog.Error("Failure", "URL", r.Request.URL, "Error", err)
	})

	c.Visit(url)

	return keywords, nil
}

func trim(s string) string {
	regex, err := regexp.Compile(`[[:space:]]+`)
	if err != nil {
		slog.Error(fmt.Sprintf("Could not create regex: %v", err))
	}
	return regex.ReplaceAllString(s, " ")
}

func parseChapter(url string) []Keyword {
	var ks []Keyword
	c := colly.NewCollector()
	c.OnHTML(("div#chapterContent"), func(h *colly.HTMLElement) {
		h.ForEach("article.reference", func(_ int, e *colly.HTMLElement) {
			var k Keyword
			k.Command = trim(e.ChildText("h2.title"))
			k.Description = trim(e.ChildText("section.section > p.p"))
			k.Syntax = trim(e.ChildText("section.refsyn"))
			k.Defaults = trim(e.ChildText("section.command_default > p"))
			k.Mode = trim(e.ChildText("section.command_modes > p"))
			
			// Extract version information from command history
			historyText := e.ChildText("section.command_history")
			k.MinVersion, k.MaxVersion = extractVersions(historyText)

			k.History.Release = e.ChildText("section.command_history > td.entry :first-child")
			k.History.Modification = e.ChildText("section.command_history > td.entry :nth-child(2)")
			k.Usage.Preamble = trim(e.ChildText("section.usage_guidelines > :not(h3.sectiontitle) "))
			k.Usage.Note = e.ChildText("section.note__content")
			k.Examples.Preamble = trim(e.ChildText("section.command_examples > p"))
			k.Examples.Code = e.ChildText("section.command_examples > pre.codeblock")
			k.DeviceTypes = extractDeviceTypes(e.ChildText("section.command_modes") + e.ChildText("section.usage_guidelines") + e.ChildText("section.command_examples"))
			ks = append(ks, k)
		})
	})
	c.OnRequest(func(r *colly.Request) {
		slog.Debug("Vist", "URL", r.URL)
	})
	c.OnError(func(r *colly.Response, err error) {
		slog.Error("Failure", "URL", r.Request.URL, "Error", err)
	})
	c.Visit(url)
	slog.Debug("SECTION KEYWORDS", "count", len(ks))
	if len(ks) > 0 {
		slog.Debug("SECTION KEYWORDS", "example", ks[0])
	}
	return ks
}

// extractVersions attempts to find min and max versions from a given text.
// This is a best-effort function due to inconsistent formatting in documentation.
func extractVersions(text string) (string, string) {
	minVersion := ""
	maxVersion := ""

	// Look for patterns like "Cisco IOS Release X.Y" or "Introduced in X.Y"
	// and also for deprecation notices.
	// This regex is a starting point and might need refinement.
	versionRegex := regexp.MustCompile(`(?:Release|Introduced in|from)\s+([\d\.]+[A-Z]?)`)
	deprecatedRegex := regexp.MustCompile(`Deprecated in\s+([\d\.]+[A-Z]?)`)

	matches := versionRegex.FindAllStringSubmatch(text, -1)
	if len(matches) > 0 {
		minVersion = matches[0][1]
		if len(matches) > 1 {
			// If multiple versions are found, assume the last one is the latest for now
			// This logic might need to be more sophisticated for ranges
			maxVersion = matches[len(matches)-1][1]
		}
	}

	deprecatedMatches := deprecatedRegex.FindStringSubmatch(text)
	if len(deprecatedMatches) > 1 {
		maxVersion = deprecatedMatches[1]
	}

	return minVersion, maxVersion
}

// extractDeviceTypes attempts to find device types from a given text.
// This is a best-effort function due to inconsistent formatting in documentation.
func extractDeviceTypes(text string) []string {
	var deviceTypes []string

	// Look for common device type patterns
	deviceRegex := regexp.MustCompile(`(?:router|switch|firewall|asr|catalyst|nexus|ios-xe|ios-xr|nx-os|cisco\s+\d{4}\s+series|cisco\s+asr\d{4}\s+series|cisco\s+catalyst\s+\d{4}\s+series|cisco\s+nexus\s+\d{4}\s+series)`)
	matches := deviceRegex.FindAllString(strings.ToLower(text), -1)

	seen := make(map[string]bool)
	for _, match := range matches {
		if !seen[match] {
			deviceTypes = append(deviceTypes, match)
			seen[match] = true
		}
	}

	return deviceTypes
}