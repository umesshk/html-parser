package internal

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func emitBytes(tagText []byte) {
	text := string(tagText)
	text = strings.TrimSpace(text)
	if text != "" {

		fmt.Printf("Text : %s ", text)
	}
}

func ParseHtml() {

	htmlFileName := flag.String("file", "ex1.html", "Provide with the html File ")

	flag.Parse()

	file, err := os.Open("web/" + *htmlFileName)

	if err != nil {
		log.Fatal("Error Opening File ", err)
		os.Exit(-1)
	}

	z := html.NewTokenizer(file)

	for {
		tt := z.Next()
		if tt == html.ErrorToken {
			fmt.Println("END")
			break
		}

		if tt == html.StartTagToken {

			tn, attr := z.TagName()
			fmt.Println(string(tn))
			if attr {
				attrN, attrV, n := z.TagAttr()
				fmt.Println("Atributes : ", string(attrN), string(attrV), n)
			}
		}
		if tt == html.TextToken {
			emitBytes(z.Text())
		}

	}
}
