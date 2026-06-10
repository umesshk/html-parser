package internal

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func Printnodes(node *html.Node) {

	/* Traversing by recursion

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
		if n.Type == html.ElementNode && n.DataAtom == atom.A {
			for _, a := range n.Attr {

				if a.Key == "href" {
					fmt.Println(a.Val)
				}
			}
		}
	}

}

func ParseHtml(file *os.File) {

	node, err := html.Parse(file)

	if err != nil {
		log.Fatal("error occured", err)
		return
	}
	Printnodes(node)

}
