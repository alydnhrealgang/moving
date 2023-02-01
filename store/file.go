package store

import (
	"io"
	"os"
	"path/filepath"
	"strings"
)

const (
	fileMode      = os.FileMode(0666)
	pathSeparator = string(os.PathSeparator)
)

func CreateFile(path string) *File {
	return &File{
		path: path,
	}
}

type File struct {
	path string
	host string
}

func (f File) DownloadTo(path string, _ int64, w io.Writer) error {
	data, err := f.readFile(f.fixPath(path))
	if nil != err {
		return err
	}
	_, err = w.Write(data)
	return err
}

func (f File) UploadFrom(path, _ string, r io.Reader) (size int64, err error) {
	data, err := io.ReadAll(r)
	if nil != err {
		return
	}

	size, err = f.saveFile(f.fixPath(path), data)
	if nil != err {
		return
	}
	return
}

func (f File) fixPath(path string) string {
	return strings.ReplaceAll(path, "/", pathSeparator)
}

func (f File) readFile(pathToRead string) (data []byte, err error) {
	pathToRead = filepath.Join(f.path, pathToRead)
	return os.ReadFile(pathToRead)
}

func (f File) saveFile(path string, data []byte) (size int64, err error) {
	pathToSave := filepath.Join(f.path, path)
	if !filepath.IsAbs(pathToSave) {
		cwd, err := os.Getwd()
		if nil != err {
			return 0, err
		}
		pathToSave = filepath.Join(cwd, pathToSave)
	}
	dir := filepath.Dir(pathToSave)
	if _, err = os.Stat(dir); nil != err {
		if os.IsNotExist(err) {
			err = os.MkdirAll(dir, fileMode)
			if nil != err {
				return
			}
		} else {
			return
		}
	}

	err = os.WriteFile(pathToSave, data, fileMode)
	if nil != err {
		return
	}
	size = int64(len(data))
	return
}
