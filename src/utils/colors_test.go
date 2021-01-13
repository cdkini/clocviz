package utils

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
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

func TestGradate(t *testing.T) {
	table := []struct {
		in    RGB
		ratio float32

		want RGB
	}{
		{in: RGB{100, 100, 100}, ratio: 1.2, want: RGB{120, 120, 120}},
		{in: RGB{100, 100, 100}, ratio: 1, want: RGB{100, 100, 100}},
		{in: RGB{100, 100, 100}, ratio: 0.8, want: RGB{80, 80, 80}},
		{in: RGB{255, 0, 0}, ratio: 0.5, want: RGB{127, 0, 0}},
		{in: RGB{0, 150, 0}, ratio: 2, want: RGB{0, 255, 0}},
		{in: RGB{0, 0, 0}, ratio: 0.5, want: RGB{0, 0, 0}},
	}

	for i, test := range table {
		name := fmt.Sprintf("Test %d - GetLangColor", i+1)
		t.Run(name, func(t *testing.T) {
			test.in.Gradate(test.ratio)
			if !cmp.Equal(test.in, test.want) {
				t.Errorf("%s: %s", name, cmp.Diff(test.in, test.want))
			}
		})
	}
}

func TestAverageColor(t *testing.T) {
	table := []struct {
		in    RGB
		color RGB
		count int

		want RGB
	}{
		{in: RGB{100, 100, 100}, color: RGB{100, 100, 100}, count: 0, want: RGB{100, 100, 100}},
		{in: RGB{100, 100, 100}, color: RGB{100, 100, 100}, count: 3, want: RGB{100, 100, 100}},
		{in: RGB{255, 255, 255}, color: RGB{100, 100, 100}, count: 3, want: RGB{216, 216, 216}},
		{in: RGB{160, 0, 0}, color: RGB{140, 150, 150}, count: 1, want: RGB{150, 75, 75}},
	}

	for i, test := range table {
		name := fmt.Sprintf("Test %d - GetLangColor", i+1)
		t.Run(name, func(t *testing.T) {
			test.in.AverageColor(test.color, test.count)
			if !cmp.Equal(test.in, test.want) {
				t.Errorf("%s: %s", name, cmp.Diff(test.in, test.want))
			}
		})
	}
}
