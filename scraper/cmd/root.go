/*
Copyright © 2024 David Gethings
*/
package cmd

import (
	"fmt"
	"log/slog"
	"os"
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
		err := getKeywords()
		if err != nil {
			return err
		}
		fmt.Println("wrote completions.go")
		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// slog.SetLogLoggerLevel(slog.LevelDebug)
}

const url = "https://www.cisco.com/c/en/us/td/docs/ios-xml/ios/fundamentals/command/cf_command_ref.html"

var keywords [][]string

func getKeywords() error {
	c := colly.NewCollector()

	c.OnHTML("ul#bookToc", func(h *colly.HTMLElement) {
		// iterate through akk the "book" sections
		h.ForEach("li", func(_ int, li *colly.HTMLElement) {
			// ignore the Introduction section, that doesn't contain any commands
			if li.ChildText("a") != "Introduction" {
				url := fmt.Sprintf("%s%s", "https://www.cisco.com", li.ChildAttr("a", "href"))
				kw := parseSection(url)
				if len(kw) > 0 {
					keywords = append(keywords, kw)
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

	fmt.Println(start)
	slog.Debug("COMPLETE LIST", "kw", keywords)
	for _, k := range keywords {
		fmt.Printf(block, k[0], k[1])
	}
	fmt.Println(end)
	return nil
}

func parseSection(url string) []string {
	var k []string
	c := colly.NewCollector()
	c.OnHTML("article.reference", func(h *colly.HTMLElement) {
		var comp []string
		// h2 is the config keyword
		// TODO: workout how better to handle this if its not a single keyword but a sentance
		kw := h.ChildText("h2.title")
		// TODO: right now we create a string of all the keywords and args/vars
		h.ForEach("p.figgroup", func(_ int, h *colly.HTMLElement) {
			var t []string
			h.ForEach("span.kwd", func(_ int, h *colly.HTMLElement) {
				t = append(t, h.Text)
			})
			kwd := strings.Join(t, " ")
			t = nil
			h.ForEach("var", func(_ int, h *colly.HTMLElement) {
				t = append(t, fmt.Sprintf("{{ %s }}", h.Text))
			})
			arg := strings.Join(t, " ")
			if kwd != "" {
				if arg != "" {
					comp = append(comp, fmt.Sprintf("%s %s", kwd, arg))
				} else {
					comp = append(comp, kwd)
				}
			}
		})
		// skip show commands
		if strings.HasPrefix(kw, "show") {
			slog.Debug("skipping show command", "keyword", kw)
			return
		}
		k = append(k, comp...)
		slog.Debug("completion", kw, comp)
	})
	c.OnRequest(func(r *colly.Request) {
		slog.Debug("Vist", "URL", r.URL)
	})
	c.OnError(func(r *colly.Response, err error) {
		slog.Error("Failure", "URL", r.Request.URL, "Error", err)
	})
	c.Visit(url)
	slog.Debug("ALL KEYWORDS", "kewords", k)
	return k
}

var start = `
package textdocument

import (
	"github.com/tliron/glsp"
	protocol "github.com/tliron/lsp/protocol_3_16"
)

func Completion(ctx *glsp.Context, params *protocol.CompletionParams) (interface{}, error) {
	var items []protocol.CompletionItem{
`

var block = `
		protocol.Completion{
			Label: "%s"
			InsertText: "%s"
		},
`

var end = `
  }
  return items, nil
}
`
