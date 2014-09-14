package fsutil

import (
	"path/filepath"
	"testing"

	"github.com/rjeczalik/fs"
	"github.com/rjeczalik/fs/memfs"
)

func equal(lhs, cas []string) bool {
	if len(lhs) != len(cas) {
		return false
	}
	for i := range cas {
		cas[i] = filepath.FromSlash(cas[i])
	}
Loop:
	for i := range lhs {
		for j := range cas {
			if lhs[i] == cas[j] {
				continue Loop
			}
		}
		return false
	}
	return true
}

func equaldiff(lhs, cas map[string][]string) bool {
	if len(lhs) != len(cas) {
		return false
	}
	for k, cas := range cas {
		if lhs, ok := lhs[filepath.FromSlash(k)]; !ok || !equal(lhs, cas) {
			return false
		}
	}
	return true
}

func TestReadpaths(t *testing.T) {
	t.Skip("TODO(rjeczalik)")
}

func TestReaddirpaths(t *testing.T) {
	cases := [...]struct {
		c    Control
		dirs map[string][]string
	}{
		0: {
			Control{FS: trees[0]},
			map[string][]string{
				filepath.FromSlash("/data/github.com/user/example"): {
					"assets",
					"dir",
				},
				filepath.FromSlash("/src/github.com/user/example"): {
					"dir",
				},
			},
		},
		1: {
			Control{FS: trees[0], Hidden: true},
			map[string][]string{
				filepath.FromSlash("/data/github.com/user/example"): {
					"assets",
					"dir",
					".git",
				},
				filepath.FromSlash("/src/github.com/user/example"): {
					"dir",
					".git",
				},
			},
		},
		2: {
			Control{FS: trees[1]},
			map[string][]string{
				filepath.FromSlash("/"): {
					"data",
					"src",
				},
				filepath.FromSlash("/data/github.com/user/example"): {
					"dir",
					"first",
					"second",
				},
				filepath.FromSlash("/src"): {
					"github.com",
				},
			},
		},
		3: {
			Control{FS: trees[2]},
			map[string][]string{
				filepath.FromSlash("/"): {
					"schema",
					"src",
				},
				filepath.FromSlash("/schema/licstat/schema"): {
					"databasequery",
					"generalinfo",
					"license",
					"monitorconf",
				},
				filepath.FromSlash("/src/licstat/schema"): {
					"tmp",
				},
			},
		},
	}
	for i, cas := range cases {
		for dir, v := range cas.dirs {
			paths := cas.c.Readdirpaths(dir)
			if paths == nil {
				t.Errorf("want paths!=nil (i=%d, dir=%s)", i, dir)
				continue
			}
			if !equal(paths, v) {
				t.Errorf("want paths=%v; got %v (i=%d, dir=%s)", v, paths, i, dir)
			}
		}
	}
}

func TestReaddirnames(t *testing.T) {
	t.Skip("TODO(rjeczalik)")
}

func TestCatchSpy(t *testing.T) {
	cases := [...]struct {
		depth int
		c     Control
		dirs  map[string][]string
	}{
		0: {
			1, Control{FS: trees[3]},
			map[string][]string{
				filepath.FromSlash("/"): {
					"/1",
					"/1.txt",
					"/a.txt",
					"/abc",
					"/abc.txt",
				},
				filepath.FromSlash("/1"): {
					"/1/2",
					"/1/2.txt",
				},
				filepath.FromSlash("/abc"): {
					"/abc/1",
					"/abc/efg",
					"/abc/efg.txt",
				},
			},
		},
		1: {
			1, Control{FS: trees[3], Hidden: true},
			map[string][]string{
				filepath.FromSlash("/"): {
					"/.1",
					"/.1.txt",
					"/.abc",
					"/.abc.txt",
					"/1",
					"/1.txt",
					"/a.txt",
					"/abc",
					"/abc.txt",
				},
				filepath.FromSlash("/1"): {
					"/1/.2",
					"/1/.2.txt",
					"/1/2",
					"/1/2.txt",
				},
				filepath.FromSlash("/abc"): {
					"/abc/.1",
					"/abc/.efg",
					"/abc/.efg.txt",
					"/abc/1",
					"/abc/efg",
					"/abc/efg.txt",
				},
			},
		},
		2: {
			3, Control{FS: trees[3]},
			map[string][]string{
				filepath.FromSlash("/abc/1"): {
					"/abc/1/2",
					"/abc/1/2/3",
					"/abc/1/2/3/txt",
					"/abc/1/2/3.txt",
					"/abc/1/2.txt",
				},
				filepath.FromSlash("/1"): {
					"/1/2",
					"/1/2/3",
					"/1/2/3/txt",
					"/1/2/3.txt",
					"/1/2.txt",
				},
				filepath.FromSlash("/abc"): {
					"/abc/1",
					"/abc/1/2",
					"/abc/1/2/3",
					"/abc/1/2/3.txt",
					"/abc/1/2.txt",
					"/abc/efg",
					"/abc/efg/hij",
					"/abc/efg/hij/txt",
					"/abc/efg/hij.txt",
					"/abc/efg.txt",
				},
			},
		},
	}
	for i, cas := range cases {
		for dir, v := range cas.dirs {
			spy := memfs.New()
			for j, fs := range []fs.Filesystem{
				0: cas.c.FS,
				1: TeeFilesystem(cas.c.FS, spy),
				2: spy,
			} {
				c := cas.c
				c.FS = fs
				found := c.Find(dir, cas.depth)
				if found == nil {
					t.Errorf("want found!=nil (i=%d, dir=%s, j=%d)", i, dir, j)
					continue
				}
				if !equal(found, v) {
					t.Errorf("want found=%v; got %v (i=%d, dir=%s, j=%d)", v,
						found, i, dir, j)
				}
			}
			found := (Control{FS: spy, Hidden: true}).Find(dir, cas.depth)
			if found == nil {
				t.Errorf("want found!=nil (i=%d, dir=%s)", i, dir)
				continue
			}
			if !equal(found, v) {
				t.Errorf("want found=%v; got %v (i=%d, dir=%s)", v, found, i, dir)
			}
		}
	}
}

func TestIntersect(t *testing.T) {
	cases := [...]struct {
		c    Control
		dirs []string
		src  string
		dst  string
	}{
		0: {
			Control{FS: trees[0]},
			[]string{
				"github.com/user/example",
				"github.com/user/example/dir",
			},
			"/src", "/data",
		},
		1: {
			Control{FS: trees[0], Hidden: true},
			[]string{
				"github.com/user/example",
				"github.com/user/example/dir",
				"github.com/user/example/.git",
			},
			"/src", "/data",
		},
		2: {
			Control{FS: trees[2]},
			[]string{
				"licstat/schema",
			},
			"/src", "/schema",
		},
		3: {
			Control{FS: trees[2], Hidden: true},
			[]string{
				"licstat/schema",
			},
			"/src", "/schema",
		},
		4: {
			Control{FS: trees[1]},
			[]string{
				"github.com/user/example",
				"github.com/user/example/dir",
			},
			"/src", "/data",
		},
		5: {
			Control{FS: trees[1], Hidden: true},
			[]string{
				"github.com/user/example",
				"github.com/user/example/dir",
			},
			"/src", "/data",
		},
	}
	for i, cas := range cases {
		dirs := cas.c.Intersect(
			filepath.FromSlash(cas.src),
			filepath.FromSlash(cas.dst),
		)
		if len(dirs) == 0 {
			t.Errorf("want len(dirs)!=0 (i=%d)", i)
			continue
		}
		if !equal(dirs, cas.dirs) {
			t.Errorf("want dirs=%v; got %v (i=%d)", cas.dirs, dirs, i)
		}
	}
}

func TestIntersectInclude(t *testing.T) {
	cases := [...]struct {
		c    Control
		diff map[string][]string
		src  string
		dst  string
	}{
		0: {
			Control{FS: trees[1]},
			map[string][]string{
				"github.com/user/example/dir": nil,
				"github.com/user/example": {
					"github.com/user/example/first",
					"github.com/user/example/second",
				},
			},
			"/src", "/data",
		},
		1: {
			Control{FS: trees[1], Hidden: true},
			map[string][]string{
				"github.com/user/example/dir": nil,
				"github.com/user/example": {
					"github.com/user/example/first",
					"github.com/user/example/second",
				},
			},
			"/src", "/data",
		},
		2: {
			Control{FS: trees[0]},
			map[string][]string{
				"github.com/user/example": {
					"github.com/user/example/assets",
				},
				"github.com/user/example/dir": nil,
			},
			"/src", "/data",
		},
		3: {
			Control{FS: trees[0], Hidden: true},
			map[string][]string{
				"github.com/user/example/.git": nil,
				"github.com/user/example": {
					"github.com/user/example/assets",
				},
				"github.com/user/example/dir": nil,
			},
			"/src", "/data",
		},
		4: {
			Control{FS: trees[2]},
			map[string][]string{
				"licstat/schema": nil,
			},
			"/src", "/schema",
		},
		5: {
			Control{FS: trees[2], Hidden: true},
			map[string][]string{
				"licstat/schema": nil,
			},
			"/src", "/schema",
		},
	}
	for i, cas := range cases {
		if i != 3 {
			continue
		}
		diff := cas.c.IntersectInclude(
			filepath.FromSlash(cas.src),
			filepath.FromSlash(cas.dst),
		)
		if len(diff) == 0 {
			t.Errorf("want len(diff)!=0 (i=%d)", i)
			continue
		}
		if !equaldiff(diff, cas.diff) {
			t.Errorf("want diff=%v; got %v (i=%d)", cas.diff, diff, i)
		}
	}
}

func TestFind(t *testing.T) {
	t.Skip("TODO(rjeczalik)")
}
