package main

import (
	"fmt"
	"os"

	"github.com/cdkini/clocviz/src/utils"
)

func main() {
	fmt.Println("vim-go")
	/*
		clocviz [dir] [git-obj]
	*/
	if len(os.Args) < 3 {
		fmt.Printf("clocviz: Usage 'clocviz [src] [optional: git hash/branch]'")
		os.Exit(1)
	}
	utils.CheckDependencies([]string{"a"})
}
