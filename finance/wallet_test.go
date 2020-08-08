/*
 * @Author: your name
 * @Date: 2020-08-08 18:15:52
 * @LastEditTime: 2020-08-08 23:33:07
 * @LastEditors: your name
 * @Description:
 * @FilePath: /learn-go-with-tests/finance/wallet_test.go
 * @可以输入预定的版权声明、个性签名、空行等
 */
package finance

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	assertBalance := func(t *testing.T, wallet Wallet, want Bitcoin) {
		got := wallet.Balance()

		if got != want {
			// t.Errorf("got %d want %d", got, want)
			t.Errorf("got %s want %s", got, want)
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		// got := wallet.Balance()

		fmt.Println("address of balance in test is", &wallet.balance)

		want := Bitcoin(10)

		assertBalance(t, wallet, want)

		// if got != want {
		// 	// t.Errorf("got %d want %d", got, want)
		// 	t.Errorf("got %s want %s", got, want)
		// }
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		wallet.Withdraw(10)

		// got := wallet.Balance()

		want := Bitcoin(10)

		assertBalance(t, wallet, want)

		// if got != want {
		// 	t.Errorf("got %s want %s", got, want)
		// }
	})

}
