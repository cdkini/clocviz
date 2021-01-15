// Package main deals with the CLI interface and overall program control flow.
package main

import (
	"log"
	"os"

	// "net/http"

	// "github.com/GeertJohan/go.rice"
	"github.com/cdkini/clocviz/src/utils"
	"github.com/cdkini/clocviz/src/web"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("clocviz: Usage 'clocviz [src]'")
	}

	// Run cloc to generate file system and related stats
	raw, err := utils.RunCloc(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	// Parse data and aggregate into object to be fed into template
	data := utils.ParseResults(raw)
	byLang := utils.GetLinesByLang(data)
	byFile := utils.GetLinesByFile(data)
	content := web.NewContent("Test", byLang, byFile)

	// Feed data into HTML/CSS/JS, start server, and render to browser
	web.Serve(content, 8080)

	os.Exit(0)
}
