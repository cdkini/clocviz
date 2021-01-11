package main

import (
	"log"
	"os"

	"github.com/cdkini/clocviz/src/utils"
	"github.com/cdkini/clocviz/src/visuals"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("clocviz: Usage 'clocviz [src] [optional: git hash/branch]'")
	}

	in := os.Args[1]
	var gitObj string
	if len(os.Args) == 3 {
		gitObj = os.Args[2]
	}

	raw, err := utils.RunCloc(in, gitObj)
	if err != nil {
		log.Fatal(err)
	}

	clean := utils.ParseResults(raw)
	data := utils.GetLinesByLang(clean)
	visuals.GenerateHTML("Test", data)
	utils.OpenBrowser("out.html")

	os.Exit(0)
}
