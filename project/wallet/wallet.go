package wallet

import (
	"errors"
	"fmt"
)

/* メモ
・Goでは、値を関数/メソッドに渡すときに値をコピーする。→参照としてポインタを渡す。
・ポインタはnilにできる。→nilチェックが必要。
・エラーを定義できる。→テストがしやすくなる。
*/

var ErrInsufficientFunds = errors.New("you are greedy.")

// type Stringer interface { // fmt パッケージで定義されたインターフェース
// 	String() string
// }

type Bitcoin int // int型からBitcoin型を作成
func (b Bitcoin) String() string { // オーバーライド的な
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin // 小文字はプライベート
}

func (w *Wallet) Deposit(amount Bitcoin) { // ポインタを取得
	// fmt.Printf("Deposit method, %v \n", &w.balance) // アドレス確認
	w.balance += amount
}
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	} else {
		w.balance -= amount
		return nil
	}
}
func (w Wallet) Balance() Bitcoin { return w.balance }
