package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

func TestCountdown(t *testing.T) {
	t.Run("prints 3 to go", func(t *testing.T) { // 表示のテスト
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
	})

	t.Run("sleep before every print", func(t *testing.T) { // スリープと表示の順序テスト
		spySleepPrinter := &CountdownOperationsSpy{}
		Countdown(spySleepPrinter, spySleepPrinter)
		want := []string{
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
			t.Errorf("wanted calls %v got %v", want, spySleepPrinter.Calls)
		}
	})
}

func TestConfigurabelSleeper(t *testing.T){
	sleepTime := 5 * time.Second

	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime{
		t.Errorf("should have slept for %v but slept for %v",sleepTime, spyTime.durationSlept)
	}
}
