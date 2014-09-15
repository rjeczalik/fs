package fsutil

import (
	"path/filepath"
	"testing"

	"github.com/rjeczalik/fs/memfs"
)

func dirlen(d memfs.Directory) (n int) {
	n = len(d)
	if _, ok := d[""].(memfs.Property); ok {
		n--
	}
	return
}

// TODO(rjeczalik)
func ExampleRelFilesystem() {
	// Output:
}

func move(fs memfs.FS, off string) memfs.FS {
	mv := memfs.New()
	if err := mv.MkdirAll(off, 0xD); err != nil {
		panic(err)
	}
	if dirlen(mv.Tree) == 0 {
		return fs
	}
	dir, base := filepath.Split(off)
	tmp := memfs.Must(mv.Cd(dir))
	tmp.Tree[base] = fs.Tree
	return mv
}

func TestRelFilesystem(t *testing.T) {
	// To randomize.
	cases := map[string]struct{}{
		"/":                      {},
		"C:/":                    {},
		"/a":                     {},
		"X:/a":                   {},
		"/tmp":                   {},
		"C:/Windows/Temp":        {},
		"/private/var/tmp":       {},
		"C:/Program Files (x86)": {},
	}
	for cas := range cases {
		cas = filepath.FromSlash(cas)
		for i, tree := range trees {
			// This is a hack, which emulates:
			//
			//   func EqualFilesystem(lhs memfs.FS, rhs fs.FS) bool
			//
			// Which should probably hit memfs/util.go file eventually.
			exp, spy := move(tree, cas), memfs.New()
			(Control{FS: TeeFilesystem(tree, RelFilesystem(spy, cas)), Hidden: true}).Find(sep, 0)
			if !memfs.Equal(spy, exp) {
				t.Errorf("want spy=exp (cas=%s, i=%d)", cas, i)
			}
			return
		}
	}
}
