package utils

import (
	"fmt"
	"os"
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

func isInSlice(target string, slice []ChartObj) (bool, ChartObj) {
	for _, obj := range slice {
		switch v := obj.(type) {
		case *Directory:
			if target == v.Name {
				return true, v
			}
		case *File:
			continue
		}
	}
	return false, nil
}
