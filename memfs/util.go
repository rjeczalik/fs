package memfs

func dirlen(d Directory) (n int) {
	n = len(d)
	if _, ok := d[""].(Property); ok {
		n--
	}
	return
}

// Equal returns true when the structure of the lhs and rhs is the same.
// It does not compare the value of the Files between the trees. If both trees
// are empty it returns true.
func Equal(lhs, rhs FS) bool {
	type node struct{ lhs, rhs Directory }
	var (
		glob = []node{{lhs: lhs.Tree, rhs: rhs.Tree}}
		nod  node
	)
	for len(glob) > 0 {
		nod, glob = glob[len(glob)-1], glob[:len(glob)-1]
		if dirlen(nod.lhs) != dirlen(nod.rhs) {
			return false
		}
		for k, lv := range nod.lhs {
			// Ignore special empty key.
			if k == "" {
				continue
			}
			rv, ok := nod.rhs[k]
			if !ok {
				return false
			}
			switch l := lv.(type) {
			case File:
				if _, ok := rv.(File); !ok {
					return false
				}
			case Directory:
				r, ok := rv.(Directory)
				if !ok {
					return false
				}
				glob = append(glob, node{lhs: l, rhs: r})
			default:
				return false
			}
		}
	}
	return true
}

// Fsck checks the fs Tree whether each node has proper type: either a File or
// a Directory. Moreover it fails if an empty key of a directory is defined and
// is not of Property type. Fsking empty tree gives true.
func Fsck(fs FS) bool {
	var (
		glob = []Directory{fs.Tree}
		dir  Directory
	)
	for len(glob) > 0 {
		dir, glob = glob[len(glob)-1], glob[:len(glob)-1]
	Loop:
		for k, v := range dir {
			if k == "" {
				if _, ok := v.(Property); ok {
					continue Loop
				}
				return false
			}
			switch v := v.(type) {
			case File:
			case Directory:
				glob = append(glob, v)
			default:
				return false
			}
		}
	}
	return true
}

// Must is a helper that wraps a call to a function returning (FS, error) and
// panics if the error is non-nil. It is intended for use in variable
// initializations such as:
//
//   var fs = memfs.Must(memfs.TabTree([]byte(".\ndir\n\tfile.txt")))
func Must(fs FS, err error) FS {
	if err != nil {
		panic(err)
	}
	return fs
}
