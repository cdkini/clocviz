package utils

import (
	"fmt"
	"math/rand"
	"time"
)

var COLORS = map[string]string{
	"Java":       "",
	"C":          "",
	"Python":     "",
	"C++":        "",
	"C#":         "",
	"JavaScript": "",
	"TypeScript": "",
	"PHP":        "",
	"SQL":        "",
	"Assembly":   "",
	"HTML":       "",
	"CSS":        "",
	"R":          "",
	"Swift":      "",
	"Ruby":       "",
	"Go":         "#0099cc",
}

func getLangColor(lang string) string {
	if color, ok := COLORS[lang]; ok {
		return color
	}
	return getRandomColorInHex()
}

func getRandomColorInHex() string {
	rand.Seed(time.Now().UnixNano())
	red := rand.Intn(255)
	green := rand.Intn(255)
	blue := rand.Intn(255)
	hex := "#" + getHex(red) + getHex(green) + getHex(blue)
	return hex
}

func getHex(num int) string {
	hex := fmt.Sprintf("%x", num)
	if len(hex) == 1 {
		hex = "0" + hex
	}
	return hex
}
