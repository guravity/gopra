package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {

	t.Run("return fastServerURL", func(t *testing.T) {
		// テストHTTPサーバー
		slowServer := makeDleayedServer(20 * time.Millisecond)
		fastServer := makeDleayedServer(0 * time.Millisecond)

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		defer slowServer.Close() // 最後に実行
		defer fastServer.Close()

		want := fastURL
		got, err := Racer(slowURL, fastURL)

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}

		if err != nil{
			t.Errorf("not expected an error, but got %q", err)
		}
	})

	t.Run("return an error when a server doesn't response within 10s", func(t *testing.T) {
		server := makeDleayedServer(25 * time.Millisecond)

		defer server.Close() 

		_,  err := ConfigurableRacer(server.URL, server.URL, 20*time.Millisecond)

		if err == nil {
			t.Errorf("expected an error but didn't get one")
		}
	})

}

func makeDleayedServer(delay time.Duration) *httptest.Server { // 絶対にOKを返すテストサーバー
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
