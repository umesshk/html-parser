package main

import (
	"flag"
	"log"
	"os"

	"github.com/umesshk/html-parser/internal"
)

func main() {
	htmlFileName := flag.String("file", "ex1.html", "Provide with the html File ")

	flag.Parse()

	file, err := os.Open("web/" + *htmlFileName)

	if err != nil {
		log.Fatal("Error Opening File ", err)
		os.Exit(-1)
	}

	internal.ParseHtml(file)

}
