package overlayfs

import (
	"os"
	"time"

	"github.com/spf13/afero"
)

type Fs []afero.Fs

func (f *Fs) Create(name string) (afero.File, error) {
	return nil, os.ErrPermission
}

func (f *Fs) Mkdir(name string, perm os.FileMode) error {
	return os.ErrPermission
}

func (f *Fs) MkdirAll(path string, perm os.FileMode) error {
	return os.ErrPermission
}

func (f Fs) Open(name string) (file afero.File, err error) {
	for _, fs := range f {
		file, err = fs.Open(name)
		if err == nil {
			return file, nil
		}
	}
	return nil, os.ErrNotExist
}

func (f Fs) OpenFile(name string, flag int, perm os.FileMode) (file afero.File, err error) {
	if flag&(os.O_WRONLY|os.O_RDWR|os.O_APPEND|os.O_CREATE|os.O_TRUNC) != 0 {
		return nil, os.ErrPermission
	}
	for _, fs := range f {
		file, err = fs.OpenFile(name, flag, perm)
		if err == nil {
			return file, nil
		}
	}
	return nil, os.ErrNotExist
}

func (f *Fs) Remove(name string) error {
	return os.ErrPermission
}

func (f *Fs) RemoveAll(path string) error {
	return os.ErrPermission
}

func (f *Fs) Rename(oldname string, newname string) error {
	return os.ErrPermission
}

func (f Fs) Stat(name string) (fi os.FileInfo, err error) {
	for _, fs := range f {
		fi, err = fs.Stat(name)
		if err == nil {
			return fi, nil
		}
	}
	return nil, os.ErrNotExist
}

func (f *Fs) Name() string {
	return "Overlay"
}

func (f *Fs) Chmod(name string, mode os.FileMode) error {
	return os.ErrPermission
}

func (f *Fs) Chtimes(name string, atime time.Time, mtime time.Time) error {
	return os.ErrPermission
}
