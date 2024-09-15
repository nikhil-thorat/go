package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Nikhil")

	got := buffer.String()
	want := "Hello, Nikhil"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
