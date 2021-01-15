package web

import (
	"fmt"
	"log"
	"runtime"

	"html/template"
	"net/http"
	"os/exec"

	"github.com/cdkini/clocviz/src/utils"

	"github.com/GeertJohan/go.rice"
	"github.com/gorilla/mux"
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

// Serve starts a new HTTP server and renders a given template to a given localhost port.
func Serve(content Content, port int) {
	tmplMessage := executeTemplate()

	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := tmplMessage.ExecuteTemplate(w, "out", content)
		if err != nil {
			log.Fatal(err)
		}
	})

	r.PathPrefix("/src/static/").Handler(http.StripPrefix("/src/static/", http.FileServer(http.Dir("./src/static/"))))
	http.Handle("/", r)

	url := fmt.Sprintf("http://localhost:%v", port)
	fmt.Printf("Running on %s\n", url)
	fmt.Println("Press 'Ctrl + C' to exit")
	openBrowser(url)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func executeTemplate() *template.Template {
	// Find a rice.Box
	tmplBox, err := rice.FindBox("../static")
	if err != nil {
		log.Fatal(err)
	}

	// Get file contents as string
	tmplString, err := tmplBox.String("index.tmpl")
	if err != nil {
		log.Fatal(err)
	}

	// Parse and execute the template
	tmplMessage, err := template.New("out").Parse(tmplString)
	if err != nil {
		log.Fatal(err)
	}

	return tmplMessage
}

func openBrowser(path string) {
	switch runtime.GOOS {
	case "linux":
		exec.Command("xdg-open", path).Start()
	case "windows":
		exec.Command("rundll32", "url.dll,FileProtocolHandler", path).Start()
	case "darwin":
		exec.Command("open", path).Start()
	}
}
