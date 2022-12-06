package main

import "testing"

func TestSearch(t *testing.T) {
	// dictionary := map[string]string{"test":"this is just a test."}; // Pythonでいうdict
	dict := Dict{"test": "this is just a test."} // Pythonでいうdict

	t.Run("known word", func(t *testing.T) {
		got, _ := dict.Search("test")
		want := "this is just a test."

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dict.Search("unknown")

		assertError(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dict := Dict{}
		word := "test"
		definition := "this is just a test!"
		err := dict.Add(word, definition)

		assertError(t, err, nil) // エラーチェック
		assertDefinition(t, dict, word, definition) // 値チェック
	})
	t.Run("existing word", func(t *testing.T) { // すでに存在する場合
		word := "test"
		definition := "this is just a test!"
		dict := Dict{word: definition}
		err := dict.Add(word, "new test")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dict, word, definition) // 値チェック（変わっていない）
	})
}

func TestUpdate(t *testing.T){
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "just a test"
		dict := Dict{word:definition}
		newDefinition := "new definition"
		err := dict.Update(word, newDefinition)
		assertError(t, err, nil)
		assertDefinition(t, dict, word, newDefinition)
	})
	t.Run("new word", func(t *testing.T) { // 存在しないのに更新しようとした場合
		dict := Dict{}
		word := "test"
		err := dict.Update(word, "new definition")
		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T){
	word := "test"
	definition := "just a test"
	dict := Dict{word:definition}
	dict.Delete(word)
	_, err := dict.Search(word)
	assertError(t, err, ErrNotFound)

}

func assertDefinition(t *testing.T, dict Dict, word, definition string) {
	t.Helper()
	got, err := dict.Search(word)
	if err != nil {
		t.Fatal("shold find added word:", err)
	}
	assertStrings(t, got, definition)
}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q but want %q given", got, want)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got error %q but want %q given", got, want)
	}
	if got == nil{
		if want == nil{
			return
		}
		t.Fatal("expected got ge an error")
	}
}
