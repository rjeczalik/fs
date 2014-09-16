// Command mktree creates a file tree out of tree output read from standard input.
// The support format is output from gotree or Unix tree commands.
//
// Usage
//
//   NAME:
//     mktree - mirrors a file tree when pipped to gotree command
//
//   USAGE:
//     mktree [OPTION]... [FILE]
//
//   OPTIONS:
//     -a          Creates also hidden files and directories
//     -o <dir=.>  Destination directory - defaults to current directory
//     -d          Create directories only
//     -h          Print usage information
//     FILE        Read tree output from the file instead of standard input
//
// Example
//
// The following example mirrors file tree of tools/fs package under /tmp/mktree
// directory.
//
//   src/github.com/rjeczalik/tools/fs $ gotree
//   .
//   ├── fs.go
//   ├── fsutil
//   │   ├── fsutil.go
//   │   ├── fsutil_test.go
//   │   ├── tee.go
//   │   └── tee_test.go
//   └── memfs
//       ├── memfs.go
//       ├── memfs_test.go
//       ├── tree.go
//       ├── tree_test.go
//       ├── util.go
//       └── util_test.go
//
//   2 directories, 11 files
//   src/github.com/rjeczalik/tools/fs $ gotree | mktree -o /tmp/mktree
//   src/github.com/rjeczalik/tools/fs $ gotree /tmp/mktree
//   /tmp/mktree/.
//   ├── fs.go
//   ├── fsutil
//   │   ├── fsutil.go
//   │   ├── fsutil_test.go
//   │   ├── tee.go
//   │   └── tee_test.go
//   └── memfs
//       ├── memfs.go
//       ├── memfs_test.go
//       ├── tree.go
//       ├── tree_test.go
//       ├── util.go
//       └── util_test.go
//
//   2 directories, 11 files
package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/rjeczalik/fs"
	"github.com/rjeczalik/fs/fsutil"
	"github.com/rjeczalik/fs/memfs"
)

const sep = string(os.PathSeparator)

const usage = `NAME:
	mktree - mirrors a file tree when pipped to gotree command

USAGE:
	mktree [OPTION]... [FILE]

OPTIONS:
	-a          Creates also hidden files and directories
	-o <dir=.>  Destination directory - defaults to current directory
	-d          Create directories only
	-h          Print usage information
	FILE        Read tree output from the file instead of standard input`

var (
	dironly bool
	all     bool
	output  string
)

var flags = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

func die(v interface{}) {
	fmt.Fprintln(os.Stderr, v)
	os.Exit(1)
}

func init() {
	output, _ = os.Getwd()
	flags.BoolVar(&all, "a", false, "")
	flags.StringVar(&output, "o", output, "")
	flags.BoolVar(&dironly, "d", false, "")
	flags.Usage = func() { fmt.Println(usage) }
	flags.Parse(os.Args[1:])
}

func cp(lhs, rhs fs.Filesystem, all bool, out string) {
	fs := fsutil.Tee(lhs, fsutil.Rel(rhs, out))
	c := fsutil.Control{FS: fs, Hidden: all}
	c.Find(sep, 0)
}

func main() {
	if len(flags.Args()) > 1 {
		die(usage)
	}
	var in io.Reader = os.Stdin
	if len(flags.Args()) == 1 {
		f, err := os.Open(flags.Args()[0])
		if err != nil {
			die(err)
		}
		defer f.Close()
		in = f
	}
	tree, err := memfs.Unix.Decode(in)
	if err != nil {
		die(err)
	}
	cp(tree, fs.Default, all, output)
}
