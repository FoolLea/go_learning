/*
 * @Author: your name
 * @Date: 2020-08-08 18:15:52
 * @LastEditTime: 2020-08-09 22:03:10
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
	// assertError := func(t *testing.T, err error) {
	// 	if err == nil {
	// 		t.Error("wanted an error but didnt get one")
	// 	}
	// }

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

	t.Run("Withdraw with funds", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		// wallet.Withdraw(Bitcoin(10))
		err := wallet.Withdraw(Bitcoin(10))

		// got := wallet.Balance()

		want := Bitcoin(10)

		assertBalance(t, wallet, want)
		assertNoError(t, err)

		// if got != want {
		// 	t.Errorf("got %s want %s", got, want)
		// }
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, startingBalance)

		// if err == nil {
		// 	t.Error("wanted an error but didn't get one")
		// }
		// assertError(t, err)
		assertError2(t, err, InsufficientFundsError)

	})

}

func assertBalance(t *testing.T, wallet Wallet, want Bitcoin) {
	got := wallet.Balance()

	if got != want {
		// t.Errorf("got %d want %d", got, want)
		t.Errorf("got %s want %s", got, want)
	}
}

func assertError2(t *testing.T, got error, want error) {
	if got == nil {
		// 如果Fatal被调用，它将停止测试。这是因为我们不希望对返回的错误进行更多断言。
		t.Fatal("didn't get an error but wanted one")
	}

	// if got.Error() != want {
	// 	t.Errorf("got '%s', want '%s'", got, want)
	// }

	if got != want {
		t.Errorf("got '%s', want '%s'", got, want)
	}
}

func assertNoError(t *testing.T, got error) {
	if got != nil {
		t.Fatal("got an error but didnt want one")
	}
}
