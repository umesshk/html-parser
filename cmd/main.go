package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/umesshk/html-parser/parser"
)

func main() {
	htmlFileName := flag.String("file", "ex1.html", "Provide with the html File ")

	flag.Parse()

	file, err := os.Open("web/" + *htmlFileName)

	if err != nil {
		log.Fatal("Error Opening File ", err)
		os.Exit(-1)
	}

	links := parser.ParseHtml(file)

	fmt.Println(links)

}
