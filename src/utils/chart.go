package utils

import (
	"encoding/json"
	"fmt"
	"log"
)

type ChartObj struct {
	Name     string      `json:"name"`
	Color    string      `json:"color"`
	Value    int         `json:"value"`
	Children []*ChartObj `json:"children"`
}

func NewChartObj(name string, color string, value int) *ChartObj {
	return &ChartObj{name, color, value, make([]*ChartObj, 0)}
}

func (c *ChartObj) Update(path []string, color string, value int) {
	if len(path) == 0 {
		return
	}

	var child *ChartObj
	curr := path[0]

	isPresent, child := isInSlice(curr, c.Children)
	if !isPresent {
		if len(path) == 1 {
			child = NewChartObj(curr, color, value)
		} else {
			child = NewChartObj(curr, color, 0)
		}
		c.Children = append(c.Children, child)
	}
	child.Update(path[1:], color, value)
}

func (c *ChartObj) ToJSON() string {
	json, err := json.Marshal(c)
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

func (c *ChartObj) String() string {
	return fmt.Sprintf("%v, %v, %v, %v", c.Name, c.Color, c.Value, c.Children)
}
