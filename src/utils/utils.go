package utils

import (
	"fmt"
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
	for i := 5; i < len(rows)-2; i++ {
		row := strings.Split(rows[i], ",")
		out = append(out, row)
	}
	return out
}

func GetLinesByFileJSON(data [][]string) {

}

func GetLinesByLangJSON(data [][]string) {
	root := NewChartObj("Languages", "", 0, make([]ChartObj, 0))

	for _, row := range data {
		lang := row[0]
		name := row[1]
		value, _ := strconv.Atoi(row[4])
		color := getLangColor(lang)

		if contains(lang, root.Children) {

		} else {
			child := NewChartObj(lang, color, 0, make([]ChartObj, 0))
			root.Children = append(root.Children, child)

		}
	}
}

func contains(needle string, haystack []ChartObj) bool {
	for _, h := range haystack {
		if needle == h.Name {
			return true
		}
	}
	return false
}
