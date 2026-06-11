package parser

import (
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

func BuildNodes(node *html.Node) []Link {

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
	var linkTags []Link
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
	return linkTags
}

func ParseHtml(r io.Reader) []Link {

	node, err := html.Parse(r)

	if err != nil {
		log.Fatal("error occured", err)
		return nil
	}
	links := BuildNodes(node)

	return links

}
