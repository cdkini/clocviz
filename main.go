package main

import (
	"log"
	"os"

	"github.com/cdkini/clocviz/src/utils"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("clocviz: Usage 'clocviz [src] [optional: git hash/branch]'")
	}

	tmp := utils.CreateTempFile()
	defer os.Remove(tmp.Name())

	in := os.Args[1]
	out := tmp.Name()
	var gitObj string
	if len(os.Args) == 3 {
		gitObj = os.Args[2]
	}

	err := utils.RunCloc(in, out, gitObj)
	if err != nil {
		log.Fatal(err)
	}

	os.Exit(0)
}
