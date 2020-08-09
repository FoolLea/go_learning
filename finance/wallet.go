/*
 * @Author: your name
 * @Date: 2020-08-08 18:15:10
 * @LastEditTime: 2020-08-09 21:53:56
 * @LastEditors: your name
 * @Description:
 * @FilePath: /learn-go-with-tests/finance/wallet.go
 * @可以输入预定的版权声明、个性签名、空行等
 */
package finance

import (
	"errors"
	"fmt"
)

// Go 允许从现有的类型创建新的类型。
// type MyName OriginalType
type Bitcoin int

type Stringer interface {
	String() string
}

type Wallet struct {
	// 小写开头 在定义它的包之外 就是私有的
	balance Bitcoin
}

// var 关键字允许我们定义包的全局值
var InsufficientFundsError = errors.New("cannot withdraw, insufficient funds")

// 指针让我们 指向 某个值，然后修改它。
// 所以，我们不是拿钱包的副本，而是拿一个指向钱包的指针，这样我们就可以改变它
func (w *Wallet) Deposit(amount Bitcoin) {
	fmt.Println("address of balance in Deposit is", &w.balance)
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	fmt.Println("address of balance in Withdraw is", &w.balance)

	if amount > w.balance {
		// return errors.New("cannot withdraw, insufficient funds")
		return InsufficientFundsError
	}
	w.balance -= amount
	return nil
}
