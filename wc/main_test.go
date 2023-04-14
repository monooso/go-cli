package main

import (
	"bytes"
	"testing"
)

func TestCountBytes(t *testing.T) {
	b := bytes.NewBufferString("word1 word2")

	exp := 11
	res := countBytes(b)

	if res != exp {
		t.Errorf("Expected %d bytes, got %d", exp, res)
	}
}

func TestCountLines(t *testing.T) {
	b := bytes.NewBufferString("word 1 word 2\nword 3 word 4\nword 5")

	exp := 3
	res := countLines(b)

	if res != exp {
		t.Errorf("Expected %d lines, got %d", exp, res)
	}
}

func TestCountWords(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3 word4\n")

	exp := 4
	res := countWords(b)

	if res != exp {
		t.Errorf("Expected %d words, got %d", exp, res)
	}
}
