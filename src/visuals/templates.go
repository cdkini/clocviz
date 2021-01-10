package visuals

import (
	"github.com/cdkini/clocviz/src/utils"
	"html/template"
	"log"
)

func generateTemplate(root *utils.Directory) {
	template, err := template.ParseFiles("out.gohtml")
	if err != nil {
		log.Fatal(err)
	}

	err := t.Execute(os.Stdout, root)

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
