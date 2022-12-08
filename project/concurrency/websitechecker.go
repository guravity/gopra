package concurrency


type WebsiteChecker func(string) bool
type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)   // スライス定義
	resultChannel := make(chan result) //

	for _, url := range urls {
		// 毎回のループで goroutine ゴルーチンを開始
		go func(u string) { // 無名関数
			// result構造体を送信
			resultChannel <- result{u, wc(u)} // send statement
			// results[u] = wc(u)
		}(url) // ()は実行
	}

	for i := 0; i < len(urls); i++ {
		result := <-resultChannel // recieve expression→sendと合わせて、mapへの書き込み制御
		results[result.string] = result.bool
	}

	return results
}
