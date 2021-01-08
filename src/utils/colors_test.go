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

func TestRandomGetLangColor(t *testing.T) {
	seen := make(map[string]bool, 0)
	for i := 0; i < 1000; i++ {
		t.Run("GetLangColor", func(t *testing.T) {
			out := GetLangColor("fakeLang")
			if _, ok := seen[out]; ok {
				t.Error("Failed to create random color")
			}
			seen[out] = true
		})
	}
}
