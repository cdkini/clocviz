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
