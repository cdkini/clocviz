package utils

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestUpdate(t *testing.T) {
	table := []struct {
		path     []string
		color    RGB
		value    int
		language string

		want ChartObj
	}{
		// Test usage in GetLinesByFile
		{
			path:     []string{"main.go"},
			color:    RGB{55, 94, 171},
			value:    5,
			language: "Go",
			want: &Directory{"root", RGB{55, 94, 171}, 5, []ChartObj{
				&File{"main.go", RGB{55, 94, 171}, 5, "Go"}}}},
		{
			path:     []string{"src", "main.py"},
			color:    RGB{53, 114, 165},
			value:    232,
			language: "Python",
			want: &Directory{"root", RGB{53, 114, 165}, 232, []ChartObj{
				&Directory{"src", RGB{53, 114, 165}, 232, []ChartObj{
					&File{"main.py", RGB{53, 114, 165}, 232, "Python"}}}}}},
		{
			path:     []string{"src", "main", "main.c"},
			color:    RGB{85, 85, 85},
			value:    165,
			language: "C",
			want: &Directory{"root", RGB{85, 85, 85}, 165, []ChartObj{
				&Directory{"src", RGB{85, 85, 85}, 165, []ChartObj{
					&Directory{"main", RGB{85, 85, 85}, 165, []ChartObj{
						&File{"main.c", RGB{85, 85, 85}, 165, "C"}}}}}}}},

		// Test usage in GetLinesByLang
		{
			path:     []string{"Python", "src", "main.py"},
			color:    RGB{53, 114, 165},
			value:    232,
			language: "Python",
			want: &Directory{"root", RGB{53, 114, 165}, 232, []ChartObj{
				&Directory{"Python", RGB{53, 114, 165}, 232, []ChartObj{
					&Directory{"src", RGB{53, 114, 165}, 232, []ChartObj{
						&File{"main.py", RGB{53, 114, 165}, 232, "Python"}}}}}}}},
		{
			path:     []string{"C", "src", "main", "main.c"},
			color:    RGB{85, 85, 85},
			value:    165,
			language: "C",
			want: &Directory{"root", RGB{85, 85, 85}, 165, []ChartObj{
				&Directory{"C", RGB{85, 85, 85}, 165, []ChartObj{
					&Directory{"src", RGB{85, 85, 85}, 165, []ChartObj{
						&Directory{"main", RGB{85, 85, 85}, 165, []ChartObj{
							&File{"main.c", RGB{85, 85, 85}, 165, "C"}}}}}}}}}},
	}

	for i, test := range table {
		c := NewDirectory("root", RGB{0, 0, 0}, 0)
		name := fmt.Sprintf("Test %d - Update", i+1)
		t.Run(name, func(t *testing.T) {
			c.update(test.path, test.color, test.value, test.language)
			if !cmp.Equal(c, test.want) {
				t.Errorf("%s: %s", name, cmp.Diff(c, test.want))
			}
		})
	}
}

// Same as TestUpdate but aggregates multiple queries into individual tests to determine persistence.
func TestPersistentUpdate(t *testing.T) {
	type MockInput struct {
		path     []string
		color    RGB
		value    int
		language string
	}

	table := []struct {
		inputs []MockInput

		want ChartObj
	}{
		// Test usage in GetLinesByFile
		{
			inputs: []MockInput{
				{[]string{"src", "main", "a.py"}, RGB{55, 94, 171}, 10, "Python"},
				{[]string{"src", "main", "b.py"}, RGB{55, 94, 171}, 20, "Python"},
				{[]string{"src", "main", "c.py"}, RGB{55, 94, 171}, 30, "Python"}},
			want: &Directory{"root", RGB{55, 94, 171}, 60, []ChartObj{
				&Directory{"src", RGB{55, 94, 171}, 60, []ChartObj{
					&Directory{"main", RGB{55, 94, 171}, 60, []ChartObj{
						&File{"a.py", RGB{55, 94, 171}, 10, "Python"},
						&File{"b.py", RGB{55, 94, 171}, 20, "Python"},
						&File{"c.py", RGB{55, 94, 171}, 30, "Python"}}}}}}}},
		{
			inputs: []MockInput{
				{[]string{"src", "dirA", "a.py"}, RGB{55, 94, 171}, 10, "Python"},
				{[]string{"src", "dirB", "b.py"}, RGB{55, 94, 171}, 20, "Python"},
				{[]string{"src", "dirC", "c.py"}, RGB{55, 94, 171}, 30, "Python"}},
			want: &Directory{"root", RGB{55, 94, 171}, 60, []ChartObj{
				&Directory{"src", RGB{55, 94, 171}, 60, []ChartObj{
					&Directory{"dirA", RGB{55, 94, 171}, 10, []ChartObj{
						&File{"a.py", RGB{55, 94, 171}, 10, "Python"}}},
					&Directory{"dirB", RGB{55, 94, 171}, 20, []ChartObj{
						&File{"b.py", RGB{55, 94, 171}, 20, "Python"}}},
					&Directory{"dirC", RGB{55, 94, 171}, 30, []ChartObj{
						&File{"c.py", RGB{55, 94, 171}, 30, "Python"}}}}}}}},

		// Test usage in GetLinesByLang
		{
			inputs: []MockInput{
				{[]string{"Python", "src", "main", "a.py"}, RGB{55, 94, 171}, 10, "Python"},
				{[]string{"Python", "src", "main", "b.py"}, RGB{55, 94, 171}, 20, "Python"},
				{[]string{"Python", "src", "main", "c.py"}, RGB{55, 94, 171}, 30, "Python"}},
			want: &Directory{"root", RGB{55, 94, 171}, 60, []ChartObj{
				&Directory{"Python", RGB{55, 94, 171}, 60, []ChartObj{
					&Directory{"src", RGB{55, 94, 171}, 60, []ChartObj{
						&Directory{"main", RGB{55, 94, 171}, 60, []ChartObj{
							&File{"a.py", RGB{55, 94, 171}, 10, "Python"},
							&File{"b.py", RGB{55, 94, 171}, 20, "Python"},
							&File{"c.py", RGB{55, 94, 171}, 30, "Python"}}}}}}}}}},
		{
			inputs: []MockInput{
				{[]string{"Python", "src", "dirA", "a.py"}, RGB{55, 94, 171}, 10, "Python"},
				{[]string{"Python", "src", "dirB", "b.py"}, RGB{55, 94, 171}, 20, "Python"},
				{[]string{"Python", "src", "dirC", "c.py"}, RGB{55, 94, 171}, 30, "Python"}},
			want: &Directory{"root", RGB{55, 94, 171}, 60, []ChartObj{
				&Directory{"Python", RGB{55, 94, 171}, 60, []ChartObj{
					&Directory{"src", RGB{55, 94, 171}, 60, []ChartObj{
						&Directory{"dirA", RGB{55, 94, 171}, 10, []ChartObj{
							&File{"a.py", RGB{55, 94, 171}, 10, "Python"}}},
						&Directory{"dirB", RGB{55, 94, 171}, 20, []ChartObj{
							&File{"b.py", RGB{55, 94, 171}, 20, "Python"}}},
						&Directory{"dirC", RGB{55, 94, 171}, 30, []ChartObj{
							&File{"c.py", RGB{55, 94, 171}, 30, "Python"}}}}}}}}}},
	}

	for i, test := range table {
		c := NewDirectory("root", RGB{0, 0, 0}, 0)
		name := fmt.Sprintf("Test %d - Update", i+1)
		t.Run(name, func(t *testing.T) {
			for _, input := range test.inputs {
				c.update(input.path, input.color, input.value, input.language)
			}
			if !cmp.Equal(c, test.want) {
				t.Errorf("%s: %s", name, cmp.Diff(c, test.want))
			}
		})
	}
}
