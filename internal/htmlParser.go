package internal

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func ParseHtml(file *os.File) {

	node, err := html.Parse(file)

	if err != nil {
		log.Fatal("Error Occured", err)
		return
	}

	for n := range node.Descendants() {

		if n.Type == html.ElementNode {
			if n.DataAtom == atom.A {
				for _, a := range n.Attr {
					if a.Key == "href" {
						fmt.Println(a.Val)
					}
				}

			}
		}

	}

}
