package main

import (
	"testing"
	"bytes"
)

func TestGreet(t *testing.T){
	buffer := bytes.Buffer{} // Writerインターフェースの実装→ｆｍｔ．Printfのテストで使える
	Greet(&buffer, "Chris")

	got := buffer.String()
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}