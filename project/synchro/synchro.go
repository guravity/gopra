package main

import "sync"

type Counter struct{
	mu sync.Mutex // ロックできる
	value int //ゼロで初期化される
}

func (c *Counter) Inc() {
	c.mu.Lock() // Lockしている間は、他のゴルーチンは待たなければいけない。
	defer c.mu.Unlock() // 上位ブロックがreturnするまで遅延→c.value++が終わるのを待つ。
	c.value++
}

func (c *Counter) Value() int {
	return c.value
}