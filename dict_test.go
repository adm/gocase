package main

import (
	"testing"
)

func TestMatch(t *testing.T) {
	cases := []struct {
		w        string
		nmatched int
		good     bool
	}{
		{"func", 4, true},
		{"defer", 5, true},
		{"for", 3, true},
		{"force", 5, true},
		{"XXX", 0, false},
		{"forcegc", 5, true},
		{"forcegcmem", 5, true},
		{"mem", 3, true},
		{"memeq", 3, true},
	}

	for _, c := range cases {
		nmatched, good := gc.dict.Match(c.w)
		if nmatched != c.nmatched || good != c.good {
			t.Errorf("failed (%s) nmatched=%d,good=%t", c.w, nmatched, good)
		}
	}
}
