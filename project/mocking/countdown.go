package main

import (
	"fmt"
	"io"
	"os"
	"time"
)
const finalWord = "Go!"
const countDownStart = 3
const write = "write"
const sleep = "sleep"

type Sleeper interface { // Sleepをモックするためのインターフェース定義
	Sleep()
}

type SpySleeper struct {
	Calls int
}
func (s *SpySleeper) Sleep(){ // インターフェースのモック実装
	s.Calls++ // メソッドが呼び出された回数をカウント
}

type CountdownOperationsSpy struct { // →Sleeper, io.Writer のモック実装
	Calls []string
}
func (s *CountdownOperationsSpy) Sleep(){
	s.Calls = append(s.Calls, sleep) // 1つ目には追加元の配列
}
func (s *CountdownOperationsSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

type SpyTime struct { // timeのモック
	durationSlept time.Duration
}
func (s *SpyTime) Sleep(duration time.Duration){ // time.sleepのモックメソッド
	s.durationSlept = duration
}

// type DefaultSleeper struct {}
// func (d *DefaultSleeper) Sleep(){
// 	time.Sleep(1*time.Second)
// }

type ConfigurableSleeper struct { // カスタマブルなスリーパー
	duration time.Duration // スリープ時間
	sleep func(time.Duration)
}
func (c *ConfigurableSleeper) Sleep(){
	c.sleep(c.duration)
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
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}
