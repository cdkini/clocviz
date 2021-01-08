package utils

import (
	"fmt"
	"reflect"
	"testing"
)

func TestUpdate(t *testing.T) {
	table := []struct {
		path  []string
		color string
		value int

		want *ChartObj
	}{
		// Test usage in GetLinesByFile
		{
			path:  []string{"main.go"},
			color: "#375eab",
			value: 5,
			want: &ChartObj{"root", "#000000", 0, []*ChartObj{
				{"main.go", "#375eab", 5, []*ChartObj{}}}}},
		{
			path:  []string{"src", "main.py"},
			color: "#3572A5",
			value: 232,
			want: &ChartObj{"root", "#000000", 0, []*ChartObj{
				{"src", "#3572A5", 0, []*ChartObj{
					{"main.py", "#3572A5", 232, []*ChartObj{}}}}}}},
		{
			path:  []string{"src", "main", "main.c"},
			color: "#555555",
			value: 165,
			want: &ChartObj{"root", "#000000", 0, []*ChartObj{
				{"src", "#555555", 0, []*ChartObj{
					{"main", "#555555", 0, []*ChartObj{
						{"main.c", "#555555", 165, []*ChartObj{}}}}}}}}},

		// Test usage in GetLinesByLang
		{
			path:  []string{"Python", "src", "main.py"},
			color: "#3572A5",
			value: 232,
			want: &ChartObj{"root", "#000000", 0, []*ChartObj{
				{"Python", "#3572A5", 0, []*ChartObj{
					{"src", "#3572A5", 0, []*ChartObj{
						{"main.py", "#3572A5", 232, []*ChartObj{}}}}}}}}},
		{
			path:  []string{"C", "src", "main", "main.c"},
			color: "#555555",
			value: 165,
			want: &ChartObj{"root", "#000000", 0, []*ChartObj{
				{"C", "#555555", 0, []*ChartObj{
					{"src", "#555555", 0, []*ChartObj{
						{"main", "#555555", 0, []*ChartObj{
							{"main.c", "#555555", 165, []*ChartObj{}}}}}}}}}}},
	}

	for i, test := range table {
		c := NewChartObj("root", "#000000", 0)
		name := fmt.Sprintf("Test %d - Update", i+1)
		t.Run(name, func(t *testing.T) {
			c.Update(test.path, test.color, test.value)
			if !reflect.DeepEqual(c, test.want) {
				t.Errorf("%s: expected %+v, received %+v", name, test.want, c)
			}
		})
	}
}

func TestPersistentUpdate(t *testing.T) {
	table := []struct {
		paths  [][]string
		colors []string
		values []int

		want *ChartObj
	}{
		// Test usage in GetLinesByFile
		{
			paths:  [][]string{{"src", "main", "a.py"}, {"src", "main", "b.py"}, {"src", "main", "c.py"}},
			colors: []string{"#375eab", "#375eab", "#375eab"},
			values: []int{10, 20, 30},
			want: &ChartObj{"root", "#000000", 0, []*ChartObj{
				{"src", "#375eab", 0, []*ChartObj{
					{"main", "#375eab", 0, []*ChartObj{
						{"a.py", "#375eab", 10, []*ChartObj{}},
						{"b.py", "#375eab", 20, []*ChartObj{}},
						{"c.py", "#375eab", 30, []*ChartObj{}}}}}}}}},
		{
			paths:  [][]string{{"src", "dirA", "a.py"}, {"src", "dirB", "b.py"}, {"src", "dirC", "c.py"}},
			colors: []string{"#375eab", "#375eab", "#375eab"},
			values: []int{10, 20, 30},
			want: &ChartObj{"root", "#000000", 0, []*ChartObj{
				{"src", "#375eab", 0, []*ChartObj{
					{"dirA", "#375eab", 0, []*ChartObj{
						{"a.py", "#375eab", 10, []*ChartObj{}}}},
					{"dirB", "#375eab", 0, []*ChartObj{
						{"b.py", "#375eab", 20, []*ChartObj{}}}},
					{"dirC", "#375eab", 0, []*ChartObj{
						{"c.py", "#375eab", 30, []*ChartObj{}}}}}}}}},
		{
			paths:  [][]string{{"Python", "src", "main", "a.py"}, {"Python", "src", "main", "b.py"}, {"Python", "src", "main", "c.py"}},
			colors: []string{"#375eab", "#375eab", "#375eab"},
			values: []int{10, 20, 30},
			want: &ChartObj{"root", "#000000", 0, []*ChartObj{
				{"Python", "#375eab", 0, []*ChartObj{
					{"src", "#375eab", 0, []*ChartObj{
						{"main", "#375eab", 0, []*ChartObj{
							{"a.py", "#375eab", 10, []*ChartObj{}},
							{"b.py", "#375eab", 20, []*ChartObj{}},
							{"c.py", "#375eab", 30, []*ChartObj{}}}}}}}}}}},
		{
			paths:  [][]string{{"Python", "src", "dirA", "a.py"}, {"Python", "src", "dirB", "b.py"}, {"Python", "src", "dirC", "c.py"}},
			colors: []string{"#375eab", "#375eab", "#375eab"},
			values: []int{10, 20, 30},
			want: &ChartObj{"root", "#000000", 0, []*ChartObj{
				{"Python", "#375eab", 0, []*ChartObj{
					{"src", "#375eab", 0, []*ChartObj{
						{"dirA", "#375eab", 0, []*ChartObj{
							{"a.py", "#375eab", 10, []*ChartObj{}}}},
						{"dirB", "#375eab", 0, []*ChartObj{
							{"b.py", "#375eab", 20, []*ChartObj{}}}},
						{"dirC", "#375eab", 0, []*ChartObj{
							{"c.py", "#375eab", 30, []*ChartObj{}}}}}}}}}}},
	}

	for i, test := range table {
		c := NewChartObj("root", "#000000", 0)
		name := fmt.Sprintf("Test %d - Update", i+1)
		t.Run(name, func(t *testing.T) {
			for i := 0; i < len(test.paths); i++ {
				c.Update(test.paths[i], test.colors[i], test.values[i])
			}
			if !reflect.DeepEqual(c, test.want) {
				t.Errorf("%s: expected %+v, received %+v", name, test.want, c)
			}
		})
	}
}
