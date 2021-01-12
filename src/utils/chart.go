package utils

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type ChartObj interface {
	String() string
}

type File struct {
	Name     string `json:"name"`
	Color    string `json:"color"`
	Value    int    `json:"size"`
	Language string `json:"language"`
}

func NewFile(name string, color string, value int, language string) *File {
	return &File{name, color, value, language}
}

func (f *File) String() string {
	return fmt.Sprintf("%v, %v, %v", f.Name, f.Color, f.Value)
}

type Directory struct {
	Name     string     `json:"name"`
	Color    string     `json:"color"`
	Children []ChartObj `json:"children"`
}

func NewDirectory(name string, color string) *Directory {
	return &Directory{name, color, make([]ChartObj, 0)}
}

func (d *Directory) String() string {
	return fmt.Sprintf("%v, %v, %v", d.Name, d.Color, d.Children)
}

func (d *Directory) Update(path []string, color string, value int, language string) {
	if len(path) == 0 {
		return
	}

	var child ChartObj
	_, child = isInSlice(path[0], d.Children)

	switch v := child.(type) {
	// Already exists so we go down the path
	case *Directory:
		v.Update(path[1:], color, value, language)
	case *File:
		return

	// Not an existing child so a new node is created
	default:
		if len(path) == 1 {
			file := NewFile(path[0], color, value, language)
			d.Children = append(d.Children, file)
		} else {
			dir := NewDirectory(path[0], color)
			d.Children = append(d.Children, dir)
			dir.Update(path[1:], color, value, language)
		}
	}
}

func GetLinesByFile(data [][]string) *Directory {
	root := NewDirectory("root", "#000000")

	for _, row := range data {
		lang := row[0]
		path := strings.Split(row[1], "/")[1:]
		color := GetLangColor(lang)
		value, err := strconv.Atoi(row[4])
		if err != nil {
			log.Fatal(err)
		}

		root.Update(path, color, value, lang)
	}

	return root
}

func GetLinesByLang(data [][]string) *Directory {
	root := NewDirectory("root", "#000000")

	for _, row := range data {
		lang := row[0]
		path := []string{lang}
		for _, str := range strings.Split(row[1], "/")[1:] {
			path = append(path, str)
		}
		color := GetLangColor(row[0])
		value, err := strconv.Atoi(row[4])
		if err != nil {
			log.Fatal(err)
		}

		root.Update(path, color, value, lang)
	}

	return root
}
