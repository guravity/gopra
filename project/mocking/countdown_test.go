package main

import (
	"testing"
	"bytes"
)

func TestCountdown(t *testing.T){
	buffer := &bytes.Buffer{} // テストできるように出力先をBufferに送信
	spySleeper := &SpySleeper{}

	Countdown(buffer, spySleeper)
	
	got := buffer.String()
	want := "3\n2\n1\nGo!"

	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
	if spySleeper.Calls != 4 {
		t.Errorf("not enough calls to sleeper, want 4 but got %d", spySleeper.Calls)
	}


}