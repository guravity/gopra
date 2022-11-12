package main

import "testing"

func TestHello(t *testing.T){
	got := Hello("Go World")
	want := "Hello, Go World"

	if got != want {
		t.Errorf("got %q want %q", got , want)
	}
}