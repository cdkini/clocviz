package utils

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"encoding/json"
)

// ChartObj is the interface that represents a node in a file tree.
type ChartObj interface {
	String() string
	ToJSON() string
}

// File represents a file in a given directory and implements the ChartObj interface.
type File struct {
	Name     string `json:"name"`
	Color    RGB    `json:"color"`
	Size     int    `json:"size"`
	Language string `json:"language"`
}

func NewFile(name string, color RGB, size int, language string) *File {
	return &File{name, color, size, language}
}

func (f *File) String() string {
	return fmt.Sprintf("%v, %v, %v, %v", f.Name, f.Color, f.Size, f.Language)
}

func (f *File) ToJSON() string {
	j, err := json.Marshal(f)
	if err != nil {
		log.Fatal(err)
	}
	return string(j)
}

// Directory represents a directory in a file tree and implements the ChartObj interface.
// Due to the recursive nature of directories, an instance of Directory can contain other ChartObj children.
type Directory struct {
	Name     string     `json:"name"`
	Color    RGB        `json:"color"`
	Size     int        `json:"size"`
	Children []ChartObj `json:"children"`
}

func NewDirectory(name string, color RGB, size int) *Directory {
	return &Directory{name, color, size, make([]ChartObj, 0)}
}

func (d *Directory) String() string {
	return fmt.Sprintf("%v, %v, %v, %v", d.Name, d.Color, d.Size, d.Children)
}

func (d *Directory) ToJSON() string {
	j, err := json.Marshal(d)
	if err != nil {
		log.Fatal(err)
	}
	return string(j)
}

// Update traverses a given path and recursively builds out a file tree from the given Directory receiver.
// Depending on the length of the path, either a Directory or File is instantiated.
func (d *Directory) Update(path []string, color RGB, size int, language string) {
	if len(path) == 0 {
		return
	}

	var child ChartObj
	_, child = isInSlice(path[0], d.Children)

	switch v := child.(type) {
	case *Directory: // Already in file tree so we move to the next segment of the path
		v.Update(path[1:], color, size, language)
	case *File:
		return
	default: // Instantiate new ChartObj based on what part of the path we are in
		if len(path) == 1 {
			file := NewFile(path[0], color, size, language)
			d.Children = append(d.Children, file)
		} else {
			dir := NewDirectory(path[0], color, 0)
			d.Children = append(d.Children, dir)
			dir.Update(path[1:], color, size, language)
		}
	}
	AverageRGB(&d.Color, d.Size, &color, size)
	d.Size += size
}

// GetLinesByFile instantiates a root Directory object and populates it with the contents of a file system.
// Nodes are populated in a depth-first manner, using Directory.Update() to traverse down to individual leaf nodes.
// Please see GetLinesByLang for an alternative method for creating our root.
func GetLinesByFile(data [][]string) *Directory {
	root := NewDirectory("root", NewRGB(0, 0, 0), 0)

	for _, row := range data {
		lang := row[0]
		path := strings.Split(row[1], "/")[1:]
		color := GetLangColor(lang)
		size, err := strconv.Atoi(row[4])
		if err != nil {
			log.Fatal(err)
		}

		root.Update(path, color, size, lang)
	}

	root.Size = 0 // Set to clean up visualization
	return root
}

// GetLinesByLang instantiates a root Directory object, sets its immediate children to the different languages used
// in a given codebase, and populates them with the contents of a file system. This approach will duplicate directories
// if they contain more than one language.
// Nodes are populated in a depth-first manner, using Directory.Update() to traverse down to individual leaf nodes.
// Please see GetLinesByFile for an alternative method for creating our root.
func GetLinesByLang(data [][]string) *Directory {
	root := NewDirectory("root", NewRGB(0, 0, 0), 0)

	for _, row := range data {
		lang := row[0]
		path := []string{lang}
		for _, str := range strings.Split(row[1], "/")[1:] {
			path = append(path, str) // Set file's language at index 0 of path slice
		}
		color := GetLangColor(row[0])
		size, err := strconv.Atoi(row[4])
		if err != nil {
			log.Fatal(err)
		}

		root.Update(path, color, size, lang)
	}

	root.Size = 0 // Set to clean up visualization
	return root
}
