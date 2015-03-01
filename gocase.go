package main

import (
	"strings"
)

type gocaser struct {
	dict *Dict
}

func NewGocaser(wordsFiles []string) *gocaser {
	gc := &gocaser{NewDict()}
	for _, file := range wordsFiles {
		gc.dict.Load(file)
	}

	return gc
}

func (gc *gocaser) GoCase(word string) string {
	if word == "" {
		return ""
	}

	words := gc.split(word)
	goCased := words[0]
	for _, w := range words[1:] {
		if w == "_" {
			continue
		}
		goCased += strings.Title(w)
	}
	return goCased
}

func (gc *gocaser) split(word string) []string {
	var words []string
	unknown := ""

	dict := gc.dict
	for len(word) > 0 {
		nmatched, exact := dict.Match(word)
		if exact && nmatched > 0 {
			if len(unknown) > 0 {
				words = append(words, unknown)
				unknown = ""
			}
			words = append(words, dict.Expand(word[:nmatched]))
		} else {
			if nmatched == 0 {
				nmatched = 1
			}
			unknown += word[:nmatched]
		}
		word = word[nmatched:]
	}

	if len(unknown) > 0 {
		words = append(words, unknown)
	}

	return words
}
