/*
Copyright © 2024 David Gethings
*/
package cmd

import (
	"fmt"
	"log/slog"
	"os"
	"regexp"
	"strings"
	"text/template"

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

		tmplFile := "keywords.tmpl"
		tmpl, err := template.New(tmplFile).ParseFiles(tmplFile)
		if err != nil {
			return err
		}
		err = tmpl.Execute(os.Stdout, kws)
		if err != nil {
			return err
		}
		return nil
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
	Command     string
	Description string
	Syntax      string
	Defaults    string
	Mode        string
	History     CommandHistory
	Usage       UsageGuideline
	Examples    Examples
}

type CommandHistory struct {
	Release      string
	Modification string
}

type UsageGuideline struct {
	Preamble string
	Note     string
}

type Examples struct {
	Preamble string
	Code     string
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

func setMode(s string) (string, error) {
	if strings.Contains(s, "Interface configuration Frame Relay DLCI configuration Template") {
		return "(config-template)", nil
	}
	if strings.Contains(s, "Global Configuration") {
		return "(config)", nil
	}
	if strings.Contains(s, "Global configuration") {
		return "(config)", nil
	}
	if strings.Contains(s, "Interface configuration") {
		return "(config-if)", nil
	}
	if strings.Contains(s, "Archive configuration") {
		return "(config-archive)", nil
	}
	if strings.Contains(s, "Archive config mode") {
		return "(config-archive-log-cfg)", nil
	}
	if strings.Contains(s, "Configuration change logger configuration") {
		return "(config-archive-log-config)", nil
	}
	if strings.Contains(s, "Line configuration") {
		return "(config-line)", nil
	}
	if strings.Contains(s, "MST configuration") {
		return "(config-mst)", nil
	}
	if strings.Contains(s, "Redundancy configuration") {
		return "(config-red)", nil
	}
	if strings.Contains(s, "Main CPU redundancy configuration") {
		return "(config-r-mc)", nil
	}
	if strings.Contains(s, "Time-range configuration") {
		return "(config-time-range)", nil
	}
	if strings.Contains(s, "log config (configuration-change logger) submode") {
		return "(config-archive-log-cfg)", nil
	}
	return "", fmt.Errorf("Unrecognised mode: %s", s)
}

func parseChapter(url string) []Keyword {
	var ks []Keyword
	c := colly.NewCollector()
	c.OnHTML(("div#chapterContent"), func(h *colly.HTMLElement) {
		h.ForEach("article.reference", func(_ int, e *colly.HTMLElement) {
			var k Keyword
			k.Command = trim(e.ChildText("h2.title"))
			descr := trim(e.ChildText("section.section > p.p"))
			slog.Debug("DESCR", "PREFORMAT", descr)
			k.Description = template.JSEscapeString(descr)
			slog.Debug("DESCR", "POSTFORMAT", k.Description)
			k.Syntax = trim(e.ChildText("section.refsyn"))
			k.Defaults = trim(e.ChildText("section.command_default > p"))
			mode, err := setMode(trim(e.ChildText("section.command_modes > p")))
			if err != nil {
				slog.Error("ModeParse", "UNKNOWN", k.Command)
			}
			k.Mode = mode
			k.History.Release = e.ChildText("section.command_history > td.entry :first-child")
			k.History.Modification = e.ChildText("section.command_history > td.entry :nth-child(2)")
			k.Usage.Preamble = trim(e.ChildText("section.usage_guidelines > :not(h3.sectiontitle) "))
			k.Usage.Note = e.ChildText("section.note__content")
			k.Examples.Preamble = trim(e.ChildText("section.command_examples > p"))
			k.Examples.Code = e.ChildText("section.command_examples > pre.codeblock")
			if strings.Contains(k.Mode, "config") {
				ks = append(ks, k)
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
	slog.Debug("SECTION KEYWORDS", "count", len(ks))
	if len(ks) > 0 {
		slog.Debug("SECTION KEYWORDS", "example", ks[0])
	}
	return ks
}
