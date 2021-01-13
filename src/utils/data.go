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
func RunCloc(in string) (string, error) {
	if _, err := os.Stat(in); os.IsNotExist(err) {
		return runClocOnGitRepo(in)
	}
	return runClocOnLocalDir(in)
}

func runClocOnLocalDir(in string) (string, error) {
	if _, err := os.Stat(in); os.IsNotExist(err) {
		return "", errors.Wrap(err, fmt.Sprintf("clocviz: Invalid path '%s' passed", in))
	}
	cloc := exec.Command("cloc", in, "--csv", "--by-file")
	if out, err := cloc.Output(); err != nil {
		return "", errors.Wrap(err, fmt.Sprintf("clocviz: Could not write data to %s", out))
	} else {
		return string(out), nil
	}
}

func runClocOnGitRepo(in string) (string, error) {
	repo := fmt.Sprintf("git://github.com/%s.git", in)

	clone := exec.Command("git", "clone", "--depth", "1", repo)
	if _, err := clone.Output(); err != nil {
		return "", errors.Wrap(err, fmt.Sprintf("clocviz: Could not find git repo '%s'", in))
	}

	dir := strings.Split(in, "/")[1]
	defer func() {
		os.RemoveAll(dir)
	}()

	return runClocOnLocalDir(dir)
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
