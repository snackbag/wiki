package projects

import (
	"os"

	"github.com/go-git/go-git/v5"
	cp "github.com/otiai10/copy"
	"github.com/snackbag/wiki/util"
)

// Download clones a git repository at the given url into ./wikidl/<id>
func Download(id string, url string) error {
	path := util.PathWikidl(id, "")
	if _, err := os.Stat(path); err == nil {
		os.RemoveAll(path)
	}

	if err := os.MkdirAll(path, 0700); err != nil {
		return err
	}

	_, err := git.PlainClone(path, false, &git.CloneOptions{
		URL:      url,
		Progress: os.Stdout,
	})
	if err != nil {
		return err
	}

	return nil
}

func CopyLocal(id string, localPath string) error {
	path := util.PathWikidl(id, "")
	if err := cp.Copy(localPath, path); err != nil {
		return err
	}

	return nil
}
