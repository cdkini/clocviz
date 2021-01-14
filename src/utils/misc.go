package utils

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"
)

// OpenBrowser determines the user's OS and opens an output file with the appropriate command.
func OpenBrowser(path string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", path).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", path).Start()
	case "darwin":
		err = exec.Command("open", path).Start()
	default:
		err = fmt.Errorf("clocviz: Script run on unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}
}

func isInSlice(target string, slice []ChartObj) (bool, ChartObj) {
	for _, obj := range slice {
		switch v := obj.(type) {
		case *Directory:
			if target == v.Name {
				return true, v
			}
		case *File:
			continue
		}
	}
	return false, nil
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
