package main

import (
	"fmt"
	"io"
	"os"
	"time"
)
const finalWord = "Go!"
const countDownStart = 3

type Sleeper interface { // Sleepをモックするためにインターフェース定義
	Sleep()
}

type SpySleeper struct {
	Calls int
}
func (s *SpySleeper) Sleep(){ // インターフェースの実装
	s.Calls++ // メソッドが呼び出された回数をカウント
}

type DefaultSleeper struct {}
func (d *DefaultSleeper) Sleep(){
	time.Sleep(1*time.Second)
}

func Countdown(out io.Writer, sleeper Sleeper){
	// fmt.Fprint(out, "3") // バッファに書き込み
	for i := countDownStart; i>0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}
	sleeper.Sleep()
	fmt.Fprint(out, finalWord)
}

func main(){
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}
