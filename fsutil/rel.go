package fsutil

import (
	"os"
	"path/filepath"

	"github.com/rjeczalik/fs"
)

type relfs struct {
	rel string // store path + separator and just use rf.rel + p?
	fs  fs.Filesystem
}

func (rf relfs) Create(p string) (fs.File, error) {
	return rf.fs.Create(filepath.Join(rf.rel, p))
}

func (rf relfs) Mkdir(p string, m os.FileMode) error {
	return rf.fs.Mkdir(filepath.Join(rf.rel, p), m)
}

func (rf relfs) MkdirAll(p string, m os.FileMode) error {
	return rf.fs.MkdirAll(filepath.Join(rf.rel, p), m)
}

func (rf relfs) Open(p string) (fs.File, error) {
	return rf.fs.Open(filepath.Join(rf.rel, p))
}

func (rf relfs) Remove(p string) error {
	return rf.fs.Remove(filepath.Join(rf.rel, p))
}

func (rf relfs) RemoveAll(p string) error {
	return rf.fs.RemoveAll(filepath.Join(rf.rel, p))
}

func (rf relfs) Stat(p string) (os.FileInfo, error) {
	return rf.fs.Stat(filepath.Join(rf.rel, p))
}

func (rf relfs) Walk(p string, fn filepath.WalkFunc) error {
	return rf.fs.Walk(filepath.Join(rf.rel, p), fn)
}

// RelFilesystem returns a filesystem which prepends rel to each path passed to
// the fs.Filesystem methods it implements.
func RelFilesystem(fs fs.Filesystem, rel string) fs.Filesystem {
	return relfs{
		rel: rel,
		fs:  fs,
	}
}
