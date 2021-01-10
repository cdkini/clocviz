package visuals

import (
	"log"
	"os"

	"html/template"

	"github.com/cdkini/clocviz/src/utils"
)

type Content struct {
	Title string
	Data  *utils.Directory
}

func GenerateHTML(title string, root *utils.Directory) {
	t, err := template.ParseFiles("src/visuals/out.gohtml")
	if err != nil {
		log.Fatal(err)
	}

	content := Content{
		Title: title,
		Data:  root,
	}

	f, err := os.Create("out.html")
	if err != nil {
		log.Fatal()
	}

	err = t.Execute(f, content)
	if err != nil {
		log.Fatal(err)
	}
}

// func openbrowser(url string) {
// 	var err error

// 	switch runtime.GOOS {
// 	case "linux":
// 		err = exec.Command("xdg-open", url).Start()
// 	case "windows":
// 		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
// 	case "darwin":
// 		err = exec.Command("open", url).Start()
// 	default:
// 		err = fmt.Errorf("unsupported platform")
// 	}
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// }
