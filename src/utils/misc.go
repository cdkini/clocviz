package utils

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"
)

func OpenBrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		fmt.Println("Linux")
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("clocviz: Script run on unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Err:", err)
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
