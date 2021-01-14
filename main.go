// Package main deals with the CLI interface and overall program control flow.
package main

import (
	"log"
	"os"

	"github.com/cdkini/clocviz/src/utils"
	"github.com/cdkini/clocviz/src/visuals"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("clocviz: Usage 'clocviz [src] [optional: git hash/branch]'")
	}

	// Run cloc to generate file system and related stats
	in := os.Args[1]
	raw, err := utils.RunCloc(in)
	if err != nil {
		log.Fatal(err)
	}

	// Parse data and separate into two JSON objects
	data := utils.ParseResults(raw)
	byLang := utils.GetLinesByLang(data)
	byFile := utils.GetLinesByFile(data)

	// Feed data into HTML/CSS/JS and render to browser
	content := visuals.NewContent("Test", byLang, byFile)
	visuals.GenerateHTML(content)
	utils.OpenBrowser("out.html")

	os.Exit(0)
}
