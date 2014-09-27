package fsutil

import (
	"os"
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
func ExampleRel() {
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

func TestRel(t *testing.T) {
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
			if n := Copy(tree, Rel(spy, cas)); n == 0 {
				t.Errorf("want n!=0 (cas=%s, i=%d)", cas, i)
				continue
			}
			if !memfs.Equal(spy, exp) {
				t.Errorf("want spy=exp (cas=%s, i=%d)", cas, i)
			}
			return
		}
	}
}

func TestRelWalk(t *testing.T) {
	cases := [...]struct {
		rel string
		p   map[string]struct{}
	}{
		0: {
			"/",
			map[string]struct{}{
				"/":             {},
				"/file":         {},
				"/dir":          {},
				"/dir/file":     {},
				"/dir/dir":      {},
				"/dir/dir/file": {},
				"/dir/dir/dir":  {},
			},
		},
		1: {
			"/tmp",
			map[string]struct{}{
				"/tmp":              {},
				"/tmp/file":         {},
				"/tmp/dir":          {},
				"/tmp/dir/file":     {},
				"/tmp/dir/dir":      {},
				"/tmp/dir/dir/file": {},
				"/tmp/dir/dir/dir":  {},
			},
		},
		2: {
			"/private/var/tmp",
			map[string]struct{}{
				"/private/var/tmp":              {},
				"/private/var/tmp/file":         {},
				"/private/var/tmp/dir":          {},
				"/private/var/tmp/dir/file":     {},
				"/private/var/tmp/dir/dir":      {},
				"/private/var/tmp/dir/dir/file": {},
				"/private/var/tmp/dir/dir/dir":  {},
			},
		}}
	var (
		p  []string
		fi []string
	)
	fn := func(path string, fileinfo os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		p, fi = append(p, path), append(fi, fileinfo.Name())
		return nil
	}
	for i, cas := range cases {
		p, fi = p[:0], fi[:0]
		if err := Rel(trees[4], filepath.FromSlash(cas.rel)).Walk(sep, fn); err != nil {
			t.Errorf("want err=nil; got err (i=%d)", err, i)
			continue
		}
		if m, n, o := len(p), len(fi), len(cas.p); m != n || n != o {
			t.Errorf("want len(p)=len(fi)=%d; got len(p)=%d, len(fi)=%d (i=%d)", o, m, n, i)
			continue
		}
		for j := range p {
			if p[j] != fi[j] {
				t.Errorf("want p=fi; got p=%q, fi=%q (i=%d, j=%d)", p[j], fi[j], i, j)
			}
			if _, ok := cas.p[p[j]]; !ok {
				t.Errorf("%q not found in cas.p (i=%d, j=%d)", p[j], i, j)
			}
		}
	}
}
