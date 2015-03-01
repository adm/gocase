package main

import (
	"flag"
	"fmt"
	"os"
)

/*
gocase
------
Suggests idiomatic gocase for function names, based on provided vocabulary.

usage: gocase -p pkg -d dir wordsfiles..
	it requires either -p  or -d for package or dir name.
	atleast a single word file should be provided.

	suggestions are printed on stdout.
	gocase comments lines for names, for which no suggestion was offered.
	gocase delete dash('-') character. see examples below.

words file:
	a line can have single word or multiple words separated by comma(,).
	a word is either bare word or of the form
		word = exp
	see below for usage.
	Blank & lines starting with '#' are ignored.

examples:
	assuming following words defined in wordsfile.
		go, case, fetch, send
		pkg = package, buf = buffer
	this what gocase will offer:
		gocase -> goCase
		fetchdata -> fetchData
		sendbuf -> sendBuffer
		pre-load -> preLoad

	checkout runtime.out file, it was run against go1.4.2 runtime package based on 'words' file.

TODO:
	+ preload common programming vocabulary.
	+ expand to all types of identifiers like var, type etc.
	+ print scope/frequency of occurence.

*/

func usage(msg string) {
	if msg != "" {
		fmt.Fprintln(os.Stderr, msg)
	}
	fmt.Fprintln(os.Stderr, "usage: gocase -p pkg -d dir words_files...")
	os.Exit(2)
}

var (
	pkg = flag.String("p", "", "package name")
	dir = flag.String("d", "", "dir name.")
)

func main() {
	flag.Parse()

	if *pkg == "" && *dir == "" {
		usage("use either pkg or dir flag.")
	}

	var wordsFiles []string
	for i := 0; i < flag.NArg(); i++ {
		wordsFiles = append(wordsFiles, flag.Arg(i))
	}
	if len(wordsFiles) == 0 {
		usage("please provide at least one words file.")
	}

	if *dir != "" {
		SuggestForDir(*dir, wordsFiles)
	} else if *pkg != "" {
		SuggestForPkg(*pkg, wordsFiles)
	} else {
		usage("")
	}
}
