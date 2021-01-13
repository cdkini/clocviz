package utils

import (
	"fmt"
	"testing"
)

func TestValidGetLangColor(t *testing.T) {
	table := []struct {
		in   string
		want RGB
	}{
		{in: "Go", want: NewRGB(55, 94, 171)},
		{in: "C", want: NewRGB(85, 85, 85)},
		{in: "Groovy", want: NewRGB(230, 159, 86)},
		{in: "HTML", want: NewRGB(228, 75, 35)},
		{in: "Java", want: NewRGB(176, 114, 25)},
		{in: "JavaScript", want: NewRGB(241, 224, 90)},
		{in: "Julia", want: NewRGB(162, 112, 186)},
	}

	for i, test := range table {
		name := fmt.Sprintf("Test %d - GetLangColor", i+1)
		t.Run(name, func(t *testing.T) {
			out := GetLangColor(test.in)
			if out != test.want {
				t.Errorf("%s: Expected %v, received %v (Lang: %s)", name, test.want, out, test.in)
			}
		})
	}
}
