package utils

import (
	"fmt"
	// "log"
	"os"
	"strings"
	"testing"

	// "io/ioutil"
	// "path/filepath"

	"github.com/pkg/errors"
)

func TestRunCloc(t *testing.T) {
	table := []struct {
		in string

		isGitRepo bool
		expected  error
	}{
		{in: ".", isGitRepo: false, expected: nil},
		{in: "fakeDir", isGitRepo: false, expected: errors.New("Source does not exist")},
		{in: "cdkini/clocviz", isGitRepo: true, expected: nil},
		{in: "fakeUser/fakeRepo", isGitRepo: true, expected: errors.New("Source does not exist")},
	}

	for i, test := range table {
		name := fmt.Sprintf("Test %d - RunCloc", i+1)
		t.Run(name, func(t *testing.T) {
			// Not testing output since we depend on cloc running accurately
			_, err := RunCloc(test.in)
			if test.expected == nil && err != nil || test.expected != nil && err == nil {
				t.Errorf("%s: Expected %v, received %v", name, test.expected, err)
			}

			if test.isGitRepo {
				dir := strings.Split(test.in, "/")[1]
				if _, err := os.Stat(dir); os.IsExist(err) {
					t.Errorf("%s: Test did not delete '%s'", name, test.expected)
				}
			}
		})
	}
}
