package ios

import (
	"bytes"
	"fmt"
	"log"

	"github.com/gocolly/colly/v2"
)

const url = "https://www.cisco.com/c/en/us/td/docs/ios/fundamentals/command/reference/cf_book/cf_a1.html"

func getKeywords() string {
	kws := make(map[string][]string)
	c := colly.NewCollector()

	c.OnHTML("p.pCE_CmdEnv", func(h *colly.HTMLElement) {
		cmd := h.ChildText("b.cCN_CmdName")
		kws[cmd] = append(kws[cmd], h.Text)
	})

	c.OnRequest(func(r *colly.Request) {
		log.Println("Visting", r.URL)
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Fatalf("%s caused: %s", r.Request.URL, err)
	})

	c.Visit(url)

	b := new(bytes.Buffer)
	for k, v := range kws {
		for _, c := range v {
			fmt.Fprintf(b, "%s: %s,", k, c)
		}
	}
	return string(b.String())
}

var Keywords = []string{
	"enable",
}
