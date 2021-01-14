package visuals

import (
	"log"
	"os"

	"html/template"

	"github.com/cdkini/clocviz/src/utils"
)

// Content is a wrapper around all data to be fed to the frontend template.
type Content struct {
	Title  string
	ByLang *utils.Directory
	ByFile *utils.Directory
}

func NewContent(title string, rootByLang *utils.Directory, rootByFile *utils.Directory) Content {
	return Content{title, rootByLang, rootByFile}
}

// GenerateHTML takes in a Content object, provides the stored data to a template, and renders an output file.
func GenerateHTML(content Content) {
	t, err := template.ParseFiles("src/visuals/out.gohtml")
	if err != nil {
		log.Fatal(err)
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
