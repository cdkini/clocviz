package utils

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"encoding/json"
)

type ChartObj interface {
	String() string
	ToJSON() string
}

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

func (d *Directory) Update(path []string, color RGB, size int, language string) {
	if len(path) == 0 {
		return
	}

	var child ChartObj
	_, child = isInSlice(path[0], d.Children)

	switch v := child.(type) {
	case *Directory:
		v.Update(path[1:], color, size, language)
	case *File:
		return
	default:
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

func (d *Directory) ToJSON() string {
	j, err := json.Marshal(d)
	if err != nil {
		log.Fatal(err)
	}
	return string(j)
}

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

	return root
}

func GetLinesByLang(data [][]string) *Directory {
	root := NewDirectory("root", NewRGB(0, 0, 0), 0)

	for _, row := range data {
		lang := row[0]
		path := []string{lang}
		for _, str := range strings.Split(row[1], "/")[1:] {
			path = append(path, str)
		}
		color := GetLangColor(row[0])
		size, err := strconv.Atoi(row[4])
		if err != nil {
			log.Fatal(err)
		}

		root.Update(path, color, size, lang)
	}

	return root
}
