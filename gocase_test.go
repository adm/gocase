package main

import (
	"reflect"
	"testing"
)

var gc *gocaser

func init() {
	gc = NewGocaser([]string{"words"})
}

func TestGoCase(t *testing.T) {
	cases := []struct {
		word string
		gc   string
	}{
		{"func", "func"},
		{"funcdefer", "funcDefer"},
		{"forcegc", "forceGc"},
		{"memeq", "memEq"},
		{"forcegchelper", "forceGcHelper"},
		{"timerproc", "timerProcess"},
	}

	for _, c := range cases {
		got := gc.GoCase(c.word)
		if got != c.gc {
			t.Errorf("failed for (%v) got (%s) expected (%s)", c.word, got, c.gc)
		}
	}
}

func TestSplit(t *testing.T) {
	cases := []struct {
		w     string
		words []string
	}{
		{"func", []string{"func"}},
		{"defer", []string{"defer"}},
		{"funcdefer", []string{"func", "defer"}},
		{"forcegc", []string{"force", "gc"}},
		{"memeq", []string{"mem", "eq"}},
		{"forcegchelper", []string{"force", "gc", "helper"}},
		{"timerproc", []string{"timer", "process"}},
	}

	for _, c := range cases {
		words := gc.split(c.w)
		if !reflect.DeepEqual(words, c.words) {
			t.Errorf("failed for (%s) got %v expected %v.", c.w, words, c.words)
		}
	}
}
