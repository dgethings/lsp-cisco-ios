package ios

import (
	"log"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

const url = "https://www.cisco.com/c/en/us/td/docs/ios/fundamentals/command/reference/cf_book/cf_a1.html"

func getKeywords() string {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var cmdTag *html.Node
	var cmd []string
	var f func(*html.Node)
	f = func(h *html.Node) {
		if h.Type == html.ElementNode && h.Data == "h2" {
			if len(h.Attr) == 1 && h.Attr[0].Val == "pCRC_CmdRefCommand" {
				// log.Println("Found tag")
				// var buf bytes.Buffer
				// w := io.Writer(&buf)
				// html.Render(w, h)
				// log.Println(buf.String())
				cmdTag = h
			}
			if cmdTag != nil {
				log.Println("Looking for text")
				if h.FirstChild != nil && h.FirstChild.Type == html.TextNode {
					log.Println("Found text node")
					cmd = append(cmd, h.FirstChild.Data)
				}
			}
		}
		for e := h.FirstChild; e != nil; e = e.NextSibling {
			f(e)
		}
	}
	f(doc)

	return string(strings.Join(cmd, ", "))
}

var Keywords = []string{
	"enable",
}
