package utils

import (
	"encoding/json"
	"log"
)

type ChartObj struct {
	Name     string     `json:"name"`
	Color    string     `json:"color"`
	Value    int        `json:"value"`
	Children []ChartObj `json:"children"`
}

func NewChartObj(name string, color string, value int, children []ChartObj) ChartObj {
	return ChartObj{name, color, value, children}
}

func (c ChartObj) ToJSON() string {
	json, err := json.Marshal(c)
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

/*
Lang
  - Go
    - ./
	  - main.go
	  - test.go
  - JSON
    - ./
	  - data/
	    - data.json
  - HTML
*/
