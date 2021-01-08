package utils

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"os/exec"

	"github.com/pkg/errors"
)

// RunCloc runs the 'cloc' tool on a target dir and writes STDOUT as a string.
// Takes an optional git hash or branch; defaults to HEAD state if not provided.
func RunCloc(in string, gitObj string) (string, error) {
	if _, err := os.Stat(in); os.IsNotExist(err) {
		return "", errors.Wrap(err, fmt.Sprintf("clocviz: Invalid path '%s' passed", in))
	}
	if len(gitObj) == 0 {
		gitObj = "."
	}
	cloc := exec.Command("cloc", in, "--csv", "--by-file", "--git", gitObj)
	if out, err := cloc.Output(); err != nil {
		return "", errors.Wrap(err, fmt.Sprintf("clocviz: Could not write data to %s", out))
	} else {
		return string(out), nil
	}
}

// ParseResults evaluates the cloc output string and converts it an easier-to-use format.
// Unnecessary rows are deleted and the result is saved as a [][]string.
func ParseResults(data string) [][]string {
	var out [][]string
	rows := strings.Split(data, "\n")
	for i := 0; i < len(rows); i++ {
		row := strings.Split(rows[i], ",")
		if len(row) == 5 && row[0] != "SUM" {
			out = append(out, row)
		}
	}
	return out
}

func GetLinesByFile(data [][]string) string {
	root := NewChartObj("root", "#000000", 0)

	for _, row := range data {
		lang := row[0]
		path := strings.Split(row[1], "/")[1:]
		color := GetLangColor(lang)
		value, err := strconv.Atoi(row[4])
		if err != nil {
			log.Fatal(err)
		}

		root.Update(path, color, value)
	}

	return root.ToJSON()
}

func GetLinesByLang(data [][]string) string {
	root := NewChartObj("root", "#000000", 0)

	for _, row := range data {
		lang := row[0]
		path := []string{lang}
		for _, str := range strings.Split(row[1], "/")[1:] {
			path = append(path, str)
		}
		color := GetLangColor(row[0])
		value, err := strconv.Atoi(row[4])
		if err != nil {
			log.Fatal(err)
		}

		root.Update(path, color, value)
	}

	return root.ToJSON()
}

func isInSlice(target string, slice []*ChartObj) (bool, *ChartObj) {
	for _, elem := range slice {
		if target == elem.Name {
			return true, elem
		}
	}
	return false, nil
}
