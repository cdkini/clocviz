package utils

import (
	"encoding/json"
	"fmt"
	"log"
)

type ChartObj interface {
	ToJSON() string
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

func (f *File) ToJSON() string {
	json, err := json.Marshal(f)
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
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

func (d *Directory) ToJSON() string {
	json, err := json.Marshal(d)
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

func (d *Directory) String() string {
	return fmt.Sprintf("%v, %v, %v", d.Name, d.Color, d.Children)
}

func (d *Directory) Update(path []string, color string, value int, language string) {
	var child ChartObj
	curr := path[0]

	isPresent, child := isInSlice(curr, d.Children)
	if !isPresent {
		if len(path) == 1 {
			child = NewFile(curr, color, value, language)
		} else {
			child = NewDirectory(curr, color)
		}
		d.Children = append(d.Children, child)
	}

	switch v := child.(type) {
	case *Directory:
		v.Update(path[1:], color, value, language)
	case *File:
		return
	}
}
