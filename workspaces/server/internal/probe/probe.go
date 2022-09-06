package probe

import (
	"io/fs"
	"io/ioutil"
	"os"
)

type LiveFile struct {
	Path     string
	BaseName string
	file     fs.FileInfo
}

func (lf *LiveFile) Create() (*os.File, error) {
	// check if the path provided exsts
	_, err := os.Stat(lf.Path)
	if err != nil {
		// attempt to create our directory with all non-existing parents when possible
		if err := os.MkdirAll(lf.Path, 0755); err != nil {
			return nil, err
		}
	}

	// attempt to create our temporarry file
	liveFile, err := ioutil.TempFile(lf.Path, lf.BaseName)
	if err != nil {
		return nil, err
	}

	file, err := liveFile.Stat()
	if err != nil {
		return nil, err
	}
	lf.file = file
	return liveFile, nil

}
func (lf *LiveFile) Remove() error {
	err := os.Remove(lf.Path + "/" + lf.file.Name())
	if err != nil {
		return err
	}
	return nil
}

func (lf *LiveFile) Exists() bool {
	_, err := os.Stat(lf.Path + "/" + lf.file.Name())
	if err == nil {
		return true
	} else {
		return false
	}
}
