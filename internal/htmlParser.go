package internal

import (
	"fmt"
	"io"
	"log"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

type Link struct {
	Href string `json:"href"`
	Text string `json:"text"`
}

var linkTags []Link

func getText(n *html.Node) string {

	var text string

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.TextNode {
			data := strings.TrimSpace(c.Data)
			if data != "" {
				text += data
			}
		}
	}
	return text
}

func Printnodes(node *html.Node) {

	// Traversing by recursion/DFS

	/*

		if node.type == html.elementnode && node.dataatom == atom.a {
			for _, a := range node.attr {
				if a.key == "href" {
					fmt.println(a.val)
				}
			}
		}

		for c := node.firstchild; c != nil; c = c.nextsibling {
			printnodes(c)
		}
	*/

	for n := range node.Descendants() {
		var link string
		var text string
		if n.Type == html.ElementNode && n.DataAtom == atom.A {

			for _, a := range n.Attr {

				if a.Key == "href" {
					val := strings.TrimSpace(a.Val)
					if val != "" {
						link = val
					}

					break
				}
			}

			text = getText(n)

			linkTags = append(linkTags, Link{Href: link, Text: text})
		}

	}
}

func ParseHtml(r io.Reader) {

	node, err := html.Parse(r)

	if err != nil {
		log.Fatal("error occured", err)
		return
	}
	Printnodes(node)

	for _, l := range linkTags {
		fmt.Println("href : ", l.Href)
		fmt.Println("Text: ", l.Text)
	}

}
