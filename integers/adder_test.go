/*
 * @Author: your name
 * @Date: 2020-07-25 17:41:43
 * @LastEditTime: 2020-07-26 23:56:53
 * @LastEditors: your name
 * @Description:
 * @FilePath: /learn-go-with-tests/integers/adder_test.go
 * @可以输入预定的版权声明、个性签名、空行等
 */
package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
