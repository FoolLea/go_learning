/*
 * @Author: your name
 * @Date: 2020-07-27 01:07:45
 * @LastEditTime: 2020-07-27 01:09:39
 * @LastEditors: your name
 * @Description:
 * @FilePath: /learn-go-with-tests/arrays/sum_test.go
 * @可以输入预定的版权声明、个性签名、空行等
 */
package main

import "testing"

func TestSUm(t *testing.T) {
	numbers := [5]int{1, 2, 3, 4, 5}

	got := Sum(numbers)
	want := 15

	if got != want {
		t.Errorf("got %d want %d given, %v", got, want, numbers)
	}
}
