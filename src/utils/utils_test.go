package utils

import (
	"fmt"
	"log"
	"os"
	"testing"

	"io/ioutil"
	"path/filepath"

	"github.com/pkg/errors"
)

func TestRunCloc(t *testing.T) {
	// Create temporary directory
	dir, err := ioutil.TempDir("", "tmpDir")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(dir)

	// Create temporary file
	tmpfn := filepath.Join(dir, "tmpFile")
	if err := ioutil.WriteFile(tmpfn, []byte("tmpFile's content"), 0666); err != nil {
		log.Fatal(err)
	}

	table := []struct {
		in     string
		gitObj string

		want     string
		expected error
	}{
		{in: dir, gitObj: "", expected: nil},
		{in: "fakeDir", gitObj: "", expected: errors.New("Source does not exist")},
		{in: dir, gitObj: "NOTHEAD", expected: errors.New("Not a valid git object")},
	}

	for i, test := range table {
		name := fmt.Sprintf("Test %d - RunCloc", i+1)
		t.Run(name, func(t *testing.T) {
			// Not testing output since we depend on cloc running accurately
			_, err := RunCloc(test.in, test.gitObj)
			if test.expected == nil && err != nil || test.expected != nil && err == nil {
				t.Errorf("%s: Expected %v, received %v", name, test.expected, err)
			}
		})
	}
}
