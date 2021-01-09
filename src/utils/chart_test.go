package utils

import (
	"fmt"
	"reflect"
	"testing"
)

func TestUpdate(t *testing.T) {
	table := []struct {
		path     []string
		color    string
		value    int
		language string

		want ChartObj
	}{
		// Test usage in GetLinesByFile
		{
			path:     []string{"main.go"},
			color:    "#375eab",
			value:    5,
			language: "Go",
			want: &Directory{"root", "#000000", []ChartObj{
				&File{"main.go", "#375eab", 5, "Go"}}}},
		{
			path:     []string{"src", "main.py"},
			color:    "#3572A5",
			value:    232,
			language: "Python",
			want: &Directory{"root", "#000000", []ChartObj{
				&Directory{"src", "#3572A5", []ChartObj{
					&File{"main.py", "#3572A5", 232, "Python"}}}}}},
		{
			path:     []string{"src", "main", "main.c"},
			color:    "#555555",
			value:    165,
			language: "C",
			want: &Directory{"root", "#000000", []ChartObj{
				&Directory{"src", "#555555", []ChartObj{
					&Directory{"main", "#555555", []ChartObj{
						&File{"main.c", "#555555", 165, "C"}}}}}}}},

		// Test usage in GetLinesByLang
		{
			path:     []string{"Python", "src", "main.py"},
			color:    "#3572A5",
			value:    232,
			language: "Python",
			want: &Directory{"root", "#000000", []ChartObj{
				&Directory{"Python", "#3572A5", []ChartObj{
					&Directory{"src", "#3572A5", []ChartObj{
						&File{"main.py", "#3572A5", 232, "Python"}}}}}}}},
		{
			path:     []string{"C", "src", "main", "main.c"},
			color:    "#555555",
			value:    165,
			language: "C",
			want: &Directory{"root", "#000000", []ChartObj{
				&Directory{"C", "#555555", []ChartObj{
					&Directory{"src", "#555555", []ChartObj{
						&Directory{"main", "#555555", []ChartObj{
							&File{"main.c", "#555555", 165, "C"}}}}}}}}}},
	}

	for i, test := range table {
		c := NewDirectory("root", "#000000")
		name := fmt.Sprintf("Test %d - Update", i+1)
		t.Run(name, func(t *testing.T) {
			c.Update(test.path, test.color, test.value, test.language)
			if !reflect.DeepEqual(c, test.want) {
				t.Errorf("%s: expected %+v, received %+v", name, test.want, c)
			}
		})
	}
}

func TestPersistentUpdate(t *testing.T) {
	type MockInput struct {
		path     []string
		color    string
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
				{[]string{"src", "main", "a.py"}, "#375eab", 10, "Python"},
				{[]string{"src", "main", "b.py"}, "#375eab", 20, "Python"},
				{[]string{"src", "main", "c.py"}, "#375eab", 30, "Python"}},
			want: &Directory{"root", "#000000", []ChartObj{
				&Directory{"src", "#375eab", []ChartObj{
					&Directory{"main", "#375eab", []ChartObj{
						&File{"a.py", "#375eab", 10, "Python"},
						&File{"b.py", "#375eab", 20, "Python"},
						&File{"c.py", "#375eab", 30, "Python"}}}}}}}},
		{
			inputs: []MockInput{
				{[]string{"src", "dirA", "a.py"}, "#375eab", 10, "Python"},
				{[]string{"src", "dirB", "b.py"}, "#375eab", 20, "Python"},
				{[]string{"src", "dirC", "c.py"}, "#375eab", 30, "Python"}},
			want: &Directory{"root", "#000000", []ChartObj{
				&Directory{"src", "#375eab", []ChartObj{
					&Directory{"dirA", "#375eab", []ChartObj{
						&File{"a.py", "#375eab", 10, "Python"}}},
					&Directory{"dirB", "#375eab", []ChartObj{
						&File{"b.py", "#375eab", 20, "Python"}}},
					&Directory{"dirC", "#375eab", []ChartObj{
						&File{"c.py", "#375eab", 30, "Python"}}}}}}}},

		// Test usage in GetLinesByLang
		{
			inputs: []MockInput{
				{[]string{"Python", "src", "main", "a.py"}, "#375eab", 10, "Python"},
				{[]string{"Python", "src", "main", "b.py"}, "#375eab", 20, "Python"},
				{[]string{"Python", "src", "main", "c.py"}, "#375eab", 30, "Python"}},
			want: &Directory{"root", "#000000", []ChartObj{
				&Directory{"Python", "#375eab", []ChartObj{
					&Directory{"src", "#375eab", []ChartObj{
						&Directory{"main", "#375eab", []ChartObj{
							&File{"a.py", "#375eab", 10, "Python"},
							&File{"b.py", "#375eab", 20, "Python"},
							&File{"c.py", "#375eab", 30, "Python"}}}}}}}}}},
		{
			inputs: []MockInput{
				{[]string{"Python", "src", "dirA", "a.py"}, "#375eab", 10, "Python"},
				{[]string{"Python", "src", "dirB", "b.py"}, "#375eab", 20, "Python"},
				{[]string{"Python", "src", "dirC", "c.py"}, "#375eab", 30, "Python"}},
			want: &Directory{"root", "#000000", []ChartObj{
				&Directory{"Python", "#375eab", []ChartObj{
					&Directory{"src", "#375eab", []ChartObj{
						&Directory{"dirA", "#375eab", []ChartObj{
							&File{"a.py", "#375eab", 10, "Python"}}},
						&Directory{"dirB", "#375eab", []ChartObj{
							&File{"b.py", "#375eab", 20, "Python"}}},
						&Directory{"dirC", "#375eab", []ChartObj{
							&File{"c.py", "#375eab", 30, "Python"}}}}}}}}}},
	}

	for i, test := range table {
		c := NewDirectory("root", "#000000")
		name := fmt.Sprintf("Test %d - Update", i+1)
		t.Run(name, func(t *testing.T) {
			for _, input := range test.inputs {
				c.Update(input.path, input.color, input.value, input.language)
			}
			if !reflect.DeepEqual(c, test.want) {
				t.Errorf("%s: expected %+v, received %+v", name, test.want, c)
			}
		})
	}
}
