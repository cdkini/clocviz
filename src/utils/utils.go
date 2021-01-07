package utils

import (
	"fmt"
	"os"

	"os/exec"

	"github.com/pkg/errors"
)

// CheckDependencies ensures that the user has all necessary software pkgs before performing visualization.
//func CheckDependencies(dependencies []string) error {
//	for _, dep := range dependencies {
//		cmd := exec.Command("command", "-v", dep)
//		if err := cmd.Run(); err != nil {
//			return errors.Wrap(err, fmt.Sprintf("clocviz: Missing '%s' dependency required to run script", dep))
//		}
//	}
//	return nil
//}

// RunCloc runs the 'cloc' tool on a target dir and writes STDOUT to a temporary file.
// Takes an optional git hash or branch; defaults to HEAD state if not provided.
func RunCloc(in string, gitObj string) (string, error) {
	if _, err := os.Stat(in); os.IsNotExist(err) {
		return "", errors.Wrap(err, fmt.Sprintf("clocviz: Invalid path '%s' passed", in))
	}
	if len(gitObj) == 0 {
		gitObj = "."
	}
	cloc := exec.Command("cloc", in, "--csv", "--by-file-by-lang", "--git", gitObj)
	if out, err := cloc.Output(); err != nil {
		return "", errors.Wrap(err, fmt.Sprintf("clocviz: Could not write data to %s", out))
	} else {
		return string(out), nil
	}
}

// CreateTempFile creates a CSV to store the results of cloc.
//func CreateTempFile() *os.File {
//	tmp, err := ioutil.TempFile("", "tmp.csv")
//	if err != nil {
//		log.Fatal(err)
//	}
//	return tmp // Must be closed => defer os.Remove(tmp.Name())
//}

// Visualize invokes a JS script to perform the visualization on the results of cloc.
//func Visualize(file *os.File) error {
//	viz := exec.Command("node", "visualize.js")
//	if err := viz.Run(); err != nil {
//		return errors.Wrap(err, "clocviz: Could not visualize provided data; are you sure your input is valid?")
//	}
//	return nil
//}
