package main

/* メモ
・mapは参照型→ポインタとして渡さなくても変更できる
・map型の基本的なデータ構造はhash tables or has map
・nilマップは読み取りができるが、書き込みではランタイムエラー
→空のマップ変数を初期化しないこと。
いい var dictionary = map[string]string{}
いい var dictionary = make(map[string]string)
だめ var m map[string]string
*/

type Dict map[string]string // mapのラッパー

// var (
// 	ErrNotFound = errors.New("could not find the word you are looking for")
// 	ErrWordExists = errors.New("cannot add word because it already exists")
// )
const (
	ErrNotFound = DictError("could not find the word you are looking for")
	ErrWordExists = DictError("cannot add word because it already exists")
	ErrWordDoesNotExist = DictError("cannot update word because it does not exist")
)

type DictError string 
func (e DictError) Error() string{
	return string(e)
}

func (d Dict) Search(word string) (string, error) { // メソッド
	definition, ok := d[word] // okはキーが正常に検出されたかどうか
	if !ok { // キーがない場合
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dict) Add(word, definition string) error { // メソッド
	_, err := d.Search(word)
	switch err {
	case ErrNotFound: // 存在しない→OK
		d[word] = definition
	case nil: // すでに存在する
		return ErrWordExists
	default: // 安全策
		return err
	}
	return nil
}

func (d Dict) Update(word, newDefinition string) error { // メソッド
	_, err := d.Search(word)
	switch err {
	case ErrNotFound: // 存在しない→だめ
		return ErrWordDoesNotExist
	case nil: // すでに存在する→OK
		d[word] = newDefinition
	default: // 安全策
		return err
	}
	return nil
}


