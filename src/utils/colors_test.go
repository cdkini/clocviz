package utils

import (
	"fmt"
	"testing"
)

func TestValidGetLangColor(t *testing.T) {
	table := []struct {
		in   string
		want string
	}{
		{in: "Go", want: "#375eab"},
		{in: "C", want: "#555555"},
		{in: "CSS", want: "#563d7c"},
		{in: "HTML", want: "#e44b23"},
		{in: "Java", want: "#b07219"},
		{in: "JavaScript", want: "#f1e05a"},
		{in: "PHP", want: "#4F5D95"},
		{in: "Python", want: "#3572A5"},
		{in: "Scheme", want: "#1e4aec"},
		{in: "VimL", want: "#199f4b"},
	}

	for i, test := range table {
		name := fmt.Sprintf("Test %d - GetLangColor", i+1)
		t.Run(name, func(t *testing.T) {
			out := GetLangColor(test.in)
			if out != test.want {
				t.Errorf("%s: Expected %s, received %s (Lang: %s)", name, test.want, out, test.in)
			}
		})
	}
}

func TestGradateHex(t *testing.T) {
	table := []struct {
		hex        string
		percentage float32

		want string
	}{
		{hex: "#375eab", percentage: 1.2, want: "#4270cd"},
		{hex: "#375eab", percentage: 1.0, want: "#375eab"},
		{hex: "#375eab", percentage: 0.8, want: "#2c4b88"},
	}
	for i, test := range table {
		name := fmt.Sprintf("Test %d - GradateHex", i+1)
		t.Run(name, func(t *testing.T) {
			out := GradateHex(test.hex, test.percentage)
			if out != test.want {
				t.Errorf("%s: Expected %s, received %s", name, test.want, out)
			}
		})
	}
}
