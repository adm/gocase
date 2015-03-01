package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type node struct {
	tbl    map[rune]*node
	isWord bool
}

type Dict struct {
	tree *node
	exp  map[string]string
}

func newNode() *node {
	return &node{tbl: make(map[rune]*node)}
}

func NewDict() *Dict {
	return &Dict{tree: newNode(), exp: make(map[string]string)}
}

func (d *Dict) Load(fname string) {
	file, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ln := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(ln, "#") || ln == "" {
			continue
		}

		for _, w := range strings.Split(ln, ",") {
			if w == "" { //trailing comma
				continue
			}
			xs := strings.Split(w, "=")
			word := strings.TrimSpace(xs[0])
			exp := ""
			if len(xs) == 2 {
				exp = strings.TrimSpace(xs[1])
			} else {
				exp = word
			}
			d.InstallWord(word, exp)
		}
	}
}

func (d *Dict) Expand(w string) string {
	if e, ok := d.exp[w]; ok {
		return e
	}
	return w
}

func (d *Dict) InstallWord(word, exp string) {
	d.install(word, exp)
	// XXX: we can save this, by using single path for both
	// by always walking lower case and checking input words case of first rune for result.
	titlew := strings.Title(word)
	if word != titlew {
		d.install(titlew, strings.Title(exp))
	}
}

func (d *Dict) install(word, exp string) {
	n := d.tree
	for _, ch := range word {
		if _, ok := n.tbl[ch]; !ok {
			n.tbl[ch] = newNode()
		}
		n = n.tbl[ch]
	}
	n.isWord = true
	if word != exp {
		d.exp[word] = exp
	}
}

func (d *Dict) Match(w string) (int, bool) {
	tree := d.tree
	lastMatchPos := -1
	exact := false
	for i, ch := range w {
		if n, ok := tree.tbl[ch]; ok {
			if n.isWord {
				lastMatchPos = i
				exact = true
			}
		} else {
			break
		}

		tree = tree.tbl[ch]
	}
	return lastMatchPos + 1, exact
}
