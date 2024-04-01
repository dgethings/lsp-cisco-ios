/*
Copyright © 2024 David Gethings
*/
package cmd

import (
	"bytes"
	"encoding/gob"
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
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ios_scraper.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

const url = "https://www.cisco.com/c/en/us/td/docs/ios-xml/ios/fundamentals/command/cf_command_ref.html"

type keywords map[string][]string

var Keywords = make(keywords)

func getKeywords() error {
	c := colly.NewCollector(colly.Async(true))

	c.OnHTML("ul#bookToc", func(h *colly.HTMLElement) {
		h.ForEach("li", func(_ int, li *colly.HTMLElement) {
			if li.ChildText("a") != "Introduction" {
				url := fmt.Sprintf("%s%s", "https://www.cisco.com", li.ChildAttr("a", "href"))
				parseSection(Keywords, url)
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
	c.Wait()

	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	enc.Encode(Keywords)
	f, err := os.Create("completions.go")
	if err != nil {
		return err
	}
	defer f.Close()
	f.WriteString("package ios\n\n")
	f.WriteString("import protocol \"github.com/tliron/glsp/protocol_3_16\"\n\n")
	fmt.Fprintf(f, "var Completions make([]protocol.CompletionItem, %d)\n\n", len(Keywords))
	block := `
Completions = append(Completions, protocol.CompletionItem{
	Label:      "%s",
	InsertText: "%s",
})
`
	for kw, comps := range Keywords {
		for _, comp := range comps {
			fmt.Fprintf(f, block, kw, comp)
		}
	}

	return nil
}

func parseSection(k keywords, url string) {
	c := colly.NewCollector()
	c.OnHTML("article.reference", func(h *colly.HTMLElement) {
		var comp []string
		kw := h.ChildText("h2.title")
		h.ForEach("p.figgroup", func(_ int, h *colly.HTMLElement) {
			var t []string
			h.ForEach("span.kwd", func(_ int, h *colly.HTMLElement) {
				t = append(t, h.Text)
			})
			kwd := strings.Join(t, " ")
			t = nil
			h.ForEach("var", func(_ int, h *colly.HTMLElement) {
				t = append(t, h.Text)
			})
			arg := strings.Join(t, " ")
			if kwd != "" {
				if arg != "" {
					comp = append(comp, fmt.Sprintf("%s %s", kwd, arg))
				} else {
					comp = append(comp, kwd)
				}
			}
			if kwd == "alias" {
				html, _ := h.DOM.Html()
				slog.Info("debug", "tag", h.Name, "html", html)
			}
		})
		if strings.HasPrefix(kw, "show") {
			return
		}
		k[kw] = comp
		slog.Debug("completion", kw, comp)
	})
	c.OnRequest(func(r *colly.Request) {
		slog.Debug("Vist", "URL", r.URL)
	})
	c.OnError(func(r *colly.Response, err error) {
		slog.Error("Failure", "URL", r.Request.URL, "Error", err)
	})
	c.Visit(url)
}
