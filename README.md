fs [![Build Status](https://img.shields.io/travis/rjeczalik/fs/master.svg)](https://travis-ci.org/rjeczalik/fs "linux_amd64") [![Build Status](https://img.shields.io/travis/rjeczalik/fs/osx.svg)](https://travis-ci.org/rjeczalik/fs "darwin_amd64") [![Build status](https://img.shields.io/appveyor/ci/rjeczalik/fs.svg)](https://ci.appveyor.com/project/rjeczalik/fs "windows_amd64") [![Coverage Status](https://img.shields.io/coveralls/rjeczalik/fs/master.svg)](https://coveralls.io/r/rjeczalik/fs?branch=master)
=====

Interface and mocks for the `os` package. Plus utility and codegen commands.

* Commands
  * [cmd/gotree](README.md#cmdgotree-)
  * [cmd/mktree](README.md#cmdmktree-)

* Packages
  * [fs](README.md#fs-)
  * [fs/memfs](README.md#fsmemfs-)
  * [fsutil](README.md#fsfsutil-)

## cmd/gotree [![GoDoc](https://godoc.org/github.com/rjeczalik/fs/cmd/gotree?status.png)](https://godoc.org/github.com/rjeczalik/fs/cmd/gotree)

Command `gotree` is Go implementation of the Unix `tree` command.

*Installation*

```bash
~ $ go get -u github.com/rjeczalik/fs/cmd/gotree
```

*Documentation*

[godoc.org/github.com/rjeczalik/fs/cmd/gotree](http://godoc.org/github.com/rjeczalik/fs/cmd/gotree)

*Usage*

```bash
~/src $ gotree github.com/rjeczalik/fs
github.com/rjeczalik/fs/.
├── fs.go
├── fsutil
│   ├── fsutil.go
│   ├── fsutil_test.go
│   ├── tee.go
│   └── tee_test.go
└── memfs
    ├── memfs.go
    ├── memfs_test.go
    ├── tree.go
    ├── tree_test.go
    ├── util.go
    └── util_test.go

2 directories, 11 files
```

**NOTE** `fs.Filesystem` does not support symlinks yet ([#1](https://github.com/rjeczalik/fs/issues/1)), that's why `gotree` will print any symlink as regular file or directory. Moreover it won't follow nor resolve any of them.

```bash
~/src $ gotree -go=80 github.com/rjeczalik/fs
memfs.Must(memfs.UnmarshalTab([]byte(".\n\tfs.go\n\tfsutil\n\t\tfsutil.go" +
	"\n\t\tfsutil_test.go\n\t\ttee.go\n\t\ttee_test.go\n\tmemfs\n\t\tmem" +
	"fs.go\n\t\tmemfs_test.go\n\t\ttree.go\n\t\ttree_test.go\n\t\tutil.g" +
	"o\n\t\tutil_test.go\n")))
```
```bash
~/src $ gotree -var=fspkg github.com/rjeczalik/fs
var fspkg = memfs.Must(memfs.UnmarshalTab([]byte(".\n\tfs.go\n\tfsutil\n\t" +
	"\tfsutil.go\n\t\tfsutil_test.go\n\t\ttee.go\n\t\ttee_test.go\n\tmem" +
	"fs\n\t\tmemfs.go\n\t\tmemfs_test.go\n\t\ttree.go\n\t\ttree_test.go\n" +
	"\t\tutil.go\n\t\tutil_test.go\n")))
```

## cmd/mktree [![GoDoc](https://godoc.org/github.com/rjeczalik/fs/cmd/mktree?status.png)](https://godoc.org/github.com/rjeczalik/fs/cmd/mktree)

Command mktree creates a file tree out of `tree` output read from standard input.

*Installation*

```bash
~ $ go get -u github.com/rjeczalik/fs/cmd/mktree
```

*Documentation*

[godoc.org/github.com/rjeczalik/fs/cmd/mktree](http://godoc.org/github.com/rjeczalik/fs/cmd/mktree)

*Usage*

```bash
~ $ gotree
.
├── dir
│   └── file.txt
└── file.txt

1 directory, 2 files

~ $ gotree | mktree -o /tmp/mktree

~ $ gotree /tmp/mktree
/tmp/mktree
├── dir
│   └── file.txt
└── file.txt

1 directory, 2 files
```
```bash
~ $ gotree > tree.txt

~ $ mktree -o /tmp/mktree2 tree.txt

~ $ gotree /tmp/mktree2
/tmp/mktree2
├── dir
│   └── file.txt
└── file.txt

1 directory, 2 files
```

## fs [![GoDoc](https://godoc.org/github.com/rjeczalik/fs?status.png)](https://godoc.org/github.com/rjeczalik/fs)

Package fs provides an interface for the filesystem-related functions from the `os` package.

*Installation*

```bash
~ $ go get -u github.com/rjeczalik/fs
```

*Documentation*

[godoc.org/github.com/rjeczalik/fs](http://godoc.org/github.com/rjeczalik/fs)

## memfs [![GoDoc](https://godoc.org/github.com/rjeczalik/fs/memfs?status.png)](https://godoc.org/github.com/rjeczalik/fs/memfs)

Package memfs provides an implementation for an in-memory filesystem.

*Installation*

```bash
~ $ go get -u github.com/rjeczalik/fs/memfs
```

*Documentation*

[godoc.org/github.com/rjeczalik/fs/memfs](http://godoc.org/github.com/rjeczalik/fs/memfs)

## fsutil [![GoDoc](https://godoc.org/github.com/rjeczalik/fs/fsutil?status.png)](https://godoc.org/github.com/rjeczalik/fs/fsutil)

Package fsutil is a collection of various filesystem utility functions.

*Installation*

```bash
~ $ go get -u github.com/rjeczalik/fs/fsutil
```

*Documentation*

[godoc.org/github.com/rjeczalik/fs/fsutil](http://godoc.org/github.com/rjeczalik/fs/fsutil)
