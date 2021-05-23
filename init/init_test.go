package init

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/ldez/go-git-cmd-wrapper/v2/git"
)

func TestInit(t *testing.T) {
	dir, err := ioutil.TempDir("", "go-git-cmd-wrapper")
	if err != nil {
		t.Fatal(err)
	}

	// clean up
	t.Cleanup(func() {
		if errRm := os.RemoveAll(dir); errRm != nil {
			log.Println(errRm)
		}
	})

	if err = os.Chdir(dir); err != nil {
		t.Fatal(err)
	}

	msg, err := git.Init(Directory("test"))
	if err != nil {
		t.Fatal(msg, err)
	}

	if ff, err := os.Stat(filepath.Join(dir, "test")); os.IsNotExist(err) {
		t.Fatal("Repository not created.", ff)
	}
}
