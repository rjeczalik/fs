package fsutil

import "github.com/rjeczalik/fs/memfs"

var trees = []memfs.FS{
	// .
	// ├── data
	// │   └── github.com
	// │       └── user
	// │           └── example
	// │               ├── .git/
	// │               ├── assets
	// │               │   ├── css
	// │               │   │   └── default.css
	// │               │   └── js
	// │               │       ├── app.js
	// │               │       └── link.js
	// │               └── dir
	// │                   └── dir.txt
	// └── src
	//     └── github.com
	//         └── user
	//             └── example
	//                 ├── .git/
	//                 ├── dir
	//                 │   └── dir.go
	//                 └── example.go
	0: memfs.Must(memfs.UnmarshalTab([]byte(".\ndata\n\tgithub.com\n\t\tuser\n\t\t" +
		"\texample\n\t\t\t\t.git/\n\t\t\t\tdir\n\t\t\t\t\tdir.txt\n\t\t\t\tas" +
		"sets\n\t\t\t\t\tjs\n\t\t\t\t\t\tapp.js\n\t\t\t\t\t\tlink.js\n\t\t\t" +
		"\t\tcss\n\t\t\t\t\t\tdefault.css\nsrc\n\tgithub.com\n\t\tuser\n\t\t" +
		"\texample\n\t\t\t\t.git/\n\t\t\t\tdir\n\t\t\t\t\tdir.go\n\t\t\t\tex" +
		"ample.go"))),
	// .
	// ├── data
	// │   └── github.com
	// │       └── user
	// │           └── example
	// │               ├── dir
	// │               │   └── dir.dat
	// │               ├── first
	// │               │   ├── css
	// │               │   │   └── first.css
	// │               │   └── js
	// │               │       └── first.js
	// │               └── second
	// │                   ├── css
	// │                   │   └── second.css
	// │                   └── js
	// │                       └── second.js
	// └── src
	//     └── github.com
	//         └── user
	//             └── example
	//                 ├── dir
	//                 │   └── dir.go
	//                 └── example.go
	1: memfs.Must(memfs.UnmarshalTab([]byte(".\ndata\n\tgithub.com\n\t\tuser\n\t" +
		"\t\texample\n\t\t\t\tdir\n\t\t\t\t\tdir.dat\n\t\t\t\tfirst\n\t\t\t\t" +
		"\tcss\n\t\t\t\t\t\tfirst.css\n\t\t\t\t\tjs\n\t\t\t\t\t\tfirst.js\n\t" +
		"\t\t\tsecond\n\t\t\t\t\tcss\n\t\t\t\t\t\tsecond.css\n\t\t\t\t\tjs\n" +
		"\t\t\t\t\t\tsecond.js\nsrc\n\tgithub.com\n\t\tuser\n\t\t\texample\n" +
		"\t\t\t\tdir\n\t\t\t\t\tdir.go\n\t\t\t\texample.go"))),
	// .
	// ├── schema
	// │   └── licstat
	// │       └── schema
	// │           ├── databasequery
	// │           │   ├── reqaddaliasls.json
	// │           │   ├── reqdeletef.json
	// │           │   ├── reqdeletels.json
	// │           │   ├── reqmergels.json
	// │           │   └── reqquerystatus.json
	// │           ├── definitions.json
	// │           ├── generalinfo
	// │           │   └── reqinstallpath.json
	// │           ├── license
	// │           │   └── reqlicensedetail.json
	// │           └── monitorconf
	// │               ├── reqaddls.json
	// │               ├── reqcheckls.json
	// │               ├── reqeditls.json
	// │               ├── reqremovels.json
	// │               └── reqstatusls.json
	// └── src
	//     └── licstat
	//         └── schema
	//             ├── schema.go
	//             └── tmp/
	2: memfs.Must(memfs.UnmarshalTab([]byte(".\nschema\n\tlicstat\n\t\tschema\n\t" +
		"\t\tdatabasequery\n\t\t\t\treqaddaliasls.json\n\t\t\t\treqdeletef.j" +
		"son\n\t\t\t\treqdeletels.json\n\t\t\t\treqmergels.json\n\t\t\t\treq" +
		"querystatus.json\n\t\t\tdefinitions.json\n\t\t\tgeneralinfo\n\t\t\t" +
		"\treqinstallpath.json\n\t\t\tlicense\n\t\t\t\treqlicensedetail.json" +
		"\n\t\t\tmonitorconf\n\t\t\t\treqaddls.json\n\t\t\t\treqcheckls.json" +
		"\n\t\t\t\treqeditls.json\n\t\t\t\treqremovels.json\n\t\t\t\treqstat" +
		"usls.json\nsrc\n\tlicstat\n\t\tschema\n\t\t\tschema.go\n\t\t\ttmp/"))),
	// .
	// ├── .1
	// │   ├── .2
	// │   │   ├── .3
	// │   │   │   ├── .txt
	// │   │   │   └── txt
	// │   │   ├── .3.txt
	// │   │   ├── 3
	// │   │   │   ├── .txt
	// │   │   │   └── txt
	// │   │   └── 3.txt
	// │   ├── .2.txt
	// │   ├── 2
	// │   │   ├── .3
	// │   │   │   ├── .txt
	// │   │   │   └── txt
	// │   │   ├── .3.txt
	// │   │   ├── 3
	// │   │   │   ├── .txt
	// │   │   │   └── txt
	// │   │   └── 3.txt
	// │   └── 2.txt
	// ├── .1.txt
	// ├── .abc
	// │   ├── .1
	// │   │   ├── .2
	// │   │   │   ├── .3
	// │   │   │   │   ├── .txt
	// │   │   │   │   └── txt
	// │   │   │   ├── .3.txt
	// │   │   │   ├── 3
	// │   │   │   │   ├── .txt
	// │   │   │   │   └── txt
	// │   │   │   └── 3.txt
	// │   │   ├── .2.txt
	// │   │   ├── 2
	// │   │   │   ├── .3
	// │   │   │   │   ├── .txt
	// │   │   │   │   └── txt
	// │   │   │   ├── .3.txt
	// │   │   │   ├── 3
	// │   │   │   │   ├── .txt
	// │   │   │   │   └── txt
	// │   │   │   └── 3.txt
	// │   │   └── 2.txt
	// │   ├── .efg
	// │   │   ├── .hij
	// │   │   │   ├── .txt
	// │   │   │   └── txt
	// │   │   ├── .hij.txt
	// │   │   ├── hij
	// │   │   │   ├── .txt
	// │   │   │   └── txt
	// │   │   └── hij.txt
	// │   ├── .efg.txt
	// │   ├── 1
	// │   │   ├── .2
	// │   │   │   ├── .3
	// │   │   │   │   ├── .txt
	// │   │   │   │   └── txt
	// │   │   │   ├── .3.txt
	// │   │   │   ├── 3
	// │   │   │   │   ├── .txt
	// │   │   │   │   └── txt
	// │   │   │   └── 3.txt
	// │   │   ├── .2.txt
	// │   │   ├── 2
	// │   │   │   ├── .3
	// │   │   │   │   ├── .txt
	// │   │   │   │   └── txt
	// │   │   │   ├── .3.txt
	// │   │   │   ├── 3
	// │   │   │   │   ├── .txt
	// │   │   │   │   └── txt
	// │   │   │   └── 3.txt
	// │   │   └── 2.txt
	// │   ├── efg
	// │   │   ├── .hij
	// │   │   │   ├── .txt
	// │   │   │   └── txt
	// │   │   ├── .hij.txt
	// │   │   ├── hij
	// │   │   │   ├── .txt
	// │   │   │   └── txt
	// │   │   └── hij.txt
	// │   └── efg.txt
	// ├── .abc.txt
	// ├── 1
	// │   ├── .2
	// │   │   ├── .3
	// │   │   │   ├── .txt
	// │   │   │   └── txt
	// │   │   ├── .3.txt
	// │   │   ├── 3
	// │   │   │   ├── .txt
	// │   │   │   └── txt
	// │   │   └── 3.txt
	// │   ├── .2.txt
	// │   ├── 2
	// │   │   ├── .3
	// │   │   │   ├── .txt
	// │   │   │   └── txt
	// │   │   ├── .3.txt
	// │   │   ├── 3
	// │   │   │   ├── .txt
	// │   │   │   └── txt
	// │   │   └── 3.txt
	// │   └── 2.txt
	// ├── 1.txt
	// ├── a.txt
	// ├── abc
	// │   ├── .1
	// │   │   ├── .2
	// │   │   │   ├── .3
	// │   │   │   │   ├── .txt
	// │   │   │   │   └── txt
	// │   │   │   ├── .3.txt
	// │   │   │   ├── 3
	// │   │   │   │   ├── .txt
	// │   │   │   │   └── txt
	// │   │   │   └── 3.txt
	// │   │   ├── .2.txt
	// │   │   ├── 2
	// │   │   │   ├── .3
	// │   │   │   │   ├── .txt
	// │   │   │   │   └── txt
	// │   │   │   ├── .3.txt
	// │   │   │   ├── 3
	// │   │   │   │   ├── .txt
	// │   │   │   │   └── txt
	// │   │   │   └── 3.txt
	// │   │   └── 2.txt
	// │   ├── .efg
	// │   │   ├── .hij
	// │   │   │   ├── .txt
	// │   │   │   └── txt
	// │   │   ├── .hij.txt
	// │   │   ├── hij
	// │   │   │   ├── .txt
	// │   │   │   └── txt
	// │   │   └── hij.txt
	// │   ├── .efg.txt
	// │   ├── 1
	// │   │   ├── .2
	// │   │   │   ├── .3
	// │   │   │   │   ├── .txt
	// │   │   │   │   └── txt
	// │   │   │   ├── .3.txt
	// │   │   │   ├── 3
	// │   │   │   │   ├── .txt
	// │   │   │   │   └── txt
	// │   │   │   └── 3.txt
	// │   │   ├── .2.txt
	// │   │   ├── 2
	// │   │   │   ├── .3
	// │   │   │   │   ├── .txt
	// │   │   │   │   └── txt
	// │   │   │   ├── .3.txt
	// │   │   │   ├── 3
	// │   │   │   │   ├── .txt
	// │   │   │   │   └── txt
	// │   │   │   └── 3.txt
	// │   │   └── 2.txt
	// │   ├── efg
	// │   │   ├── .hij
	// │   │   │   ├── .txt
	// │   │   │   └── txt
	// │   │   ├── .hij.txt
	// │   │   ├── hij
	// │   │   │   ├── .txt
	// │   │   │   └── txt
	// │   │   └── hij.txt
	// │   └── efg.txt
	// └── abc.txt
	3: memfs.Must(memfs.UnmarshalTab([]byte(".\n\t.1\n\t\t.2\n\t\t\t.3\n\t\t\t\t" +
		".txt\n\t\t\t\ttxt\n\t\t\t.3.txt\n\t\t\t3\n\t\t\t\t.txt\n\t\t\t\ttxt" +
		"\n\t\t\t3.txt\n\t\t.2.txt\n\t\t2\n\t\t\t.3\n\t\t\t\t.txt\n\t\t\t\tt" +
		"xt\n\t\t\t.3.txt\n\t\t\t3\n\t\t\t\t.txt\n\t\t\t\ttxt\n\t\t\t3.txt\n" +
		"\t\t2.txt\n\t.1.txt\n\t.abc\n\t\t.1\n\t\t\t.2\n\t\t\t\t.3\n\t\t\t\t" +
		"\t.txt\n\t\t\t\t\ttxt\n\t\t\t\t.3.txt\n\t\t\t\t3\n\t\t\t\t\t.txt\n\t" +
		"\t\t\t\ttxt\n\t\t\t\t3.txt\n\t\t\t.2.txt\n\t\t\t2\n\t\t\t\t.3\n\t\t" +
		"\t\t\t.txt\n\t\t\t\t\ttxt\n\t\t\t\t.3.txt\n\t\t\t\t3\n\t\t\t\t\t.tx" +
		"t\n\t\t\t\t\ttxt\n\t\t\t\t3.txt\n\t\t\t2.txt\n\t\t.efg\n\t\t\t.hij\n" +
		"\t\t\t\t.txt\n\t\t\t\ttxt\n\t\t\t.hij.txt\n\t\t\thij\n\t\t\t\t.txt\n" +
		"\t\t\t\ttxt\n\t\t\thij.txt\n\t\t.efg.txt\n\t\t1\n\t\t\t.2\n\t\t\t\t" +
		".3\n\t\t\t\t\t.txt\n\t\t\t\t\ttxt\n\t\t\t\t.3.txt\n\t\t\t\t3\n\t\t\t" +
		"\t\t.txt\n\t\t\t\t\ttxt\n\t\t\t\t3.txt\n\t\t\t.2.txt\n\t\t\t2\n\t\t" +
		"\t\t.3\n\t\t\t\t\t.txt\n\t\t\t\t\ttxt\n\t\t\t\t.3.txt\n\t\t\t\t3\n\t" +
		"\t\t\t\t.txt\n\t\t\t\t\ttxt\n\t\t\t\t3.txt\n\t\t\t2.txt\n\t\tefg\n\t" +
		"\t\t.hij\n\t\t\t\t.txt\n\t\t\t\ttxt\n\t\t\t.hij.txt\n\t\t\thij\n\t\t" +
		"\t\t.txt\n\t\t\t\ttxt\n\t\t\thij.txt\n\t\tefg.txt\n\t.abc.txt\n\t1\n" +
		"\t\t.2\n\t\t\t.3\n\t\t\t\t.txt\n\t\t\t\ttxt\n\t\t\t.3.txt\n\t\t\t3\n" +
		"\t\t\t\t.txt\n\t\t\t\ttxt\n\t\t\t3.txt\n\t\t.2.txt\n\t\t2\n\t\t\t.3" +
		"\n\t\t\t\t.txt\n\t\t\t\ttxt\n\t\t\t.3.txt\n\t\t\t3\n\t\t\t\t.txt\n\t" +
		"\t\t\ttxt\n\t\t\t3.txt\n\t\t2.txt\n\t1.txt\n\ta.txt\n\tabc\n\t\t.1\n" +
		"\t\t\t.2\n\t\t\t\t.3\n\t\t\t\t\t.txt\n\t\t\t\t\ttxt\n\t\t\t\t.3.txt" +
		"\n\t\t\t\t3\n\t\t\t\t\t.txt\n\t\t\t\t\ttxt\n\t\t\t\t3.txt\n\t\t\t.2" +
		".txt\n\t\t\t2\n\t\t\t\t.3\n\t\t\t\t\t.txt\n\t\t\t\t\ttxt\n\t\t\t\t." +
		"3.txt\n\t\t\t\t3\n\t\t\t\t\t.txt\n\t\t\t\t\ttxt\n\t\t\t\t3.txt\n\t\t" +
		"\t2.txt\n\t\t.efg\n\t\t\t.hij\n\t\t\t\t.txt\n\t\t\t\ttxt\n\t\t\t.hi" +
		"j.txt\n\t\t\thij\n\t\t\t\t.txt\n\t\t\t\ttxt\n\t\t\thij.txt\n\t\t.ef" +
		"g.txt\n\t\t1\n\t\t\t.2\n\t\t\t\t.3\n\t\t\t\t\t.txt\n\t\t\t\t\ttxt\n" +
		"\t\t\t\t.3.txt\n\t\t\t\t3\n\t\t\t\t\t.txt\n\t\t\t\t\ttxt\n\t\t\t\t3" +
		".txt\n\t\t\t.2.txt\n\t\t\t2\n\t\t\t\t.3\n\t\t\t\t\t.txt\n\t\t\t\t\t" +
		"txt\n\t\t\t\t.3.txt\n\t\t\t\t3\n\t\t\t\t\t.txt\n\t\t\t\t\ttxt\n\t\t" +
		"\t\t3.txt\n\t\t\t2.txt\n\t\tefg\n\t\t\t.hij\n\t\t\t\t.txt\n\t\t\t\t" +
		"txt\n\t\t\t.hij.txt\n\t\t\thij\n\t\t\t\t.txt\n\t\t\t\ttxt\n\t\t\thi" +
		"j.txt\n\t\tefg.txt\n\tabc.txt\n"))),
}
