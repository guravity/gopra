package concurrency

import (
	"testing"
	"time"
)

func slowStubWebsiteChecker(_ string) bool { // checkerのスタブ
	time.Sleep(20 * time.Microsecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 1000)
	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}
	for i := 0; i < b.N; i++ {
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}
