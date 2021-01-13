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
}

type File struct {
	Name     string `json:"name"`
	Color    RGB    `json:"color"`
	Value    int    `json:"size"`
	Language string `json:"language"`
}

func NewFile(name string, color RGB, value int, language string) *File {
	return &File{name, color, value, language}
}

func (f *File) String() string {
	return fmt.Sprintf("%v, %v, %v, %v", f.Name, f.Color, f.Value, f.Language)
}

type Directory struct {
	Name      string     `json:"name"`
	Color     RGB        `json:"color"`
	FileCount int        `json:"fileCount"`
	Children  []ChartObj `json:"children"`
}

func NewDirectory(name string, color RGB) *Directory {
	return &Directory{name, color, 0, make([]ChartObj, 0)}
}

func (d *Directory) String() string {
	return fmt.Sprintf("%v, %v, %v, %v", d.Name, d.Color, d.Children, d.FileCount)
}

func (d *Directory) Update(path []string, color RGB, value int, language string) {
	if len(path) == 0 {
		return
	}

	var child ChartObj
	_, child = isInSlice(path[0], d.Children)

	switch v := child.(type) {
	case *Directory:
		v.Update(path[1:], color, value, language)
	case *File:
		return
	default:
		d.addNewChild(path, color, value, language)
	}
	d.FileCount++
}

func (d *Directory) addNewChild(path []string, color RGB, value int, language string) {
	if len(path) == 1 {
		file := NewFile(path[0], color, value, language)
		d.Children = append(d.Children, file)
	} else {
		dir := NewDirectory(path[0], color)
		d.Children = append(d.Children, dir)
		dir.Update(path[1:], color, value, language)
	}
}

func (d *Directory) ToJSON() string {
	j, err := json.Marshal(d)
	if err != nil {
		log.Fatal(err)
	}
	return string(j)
}

func GetLinesByFile(data [][]string) *Directory {
	root := NewDirectory("root", NewRGB(0, 0, 0))

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
	root := NewDirectory("root", NewRGB(0, 0, 0))

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
