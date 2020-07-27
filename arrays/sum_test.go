/*
 * @Author: your name
 * @Date: 2020-07-27 01:07:45
 * @LastEditTime: 2020-07-27 22:27:41
 * @LastEditors: your name
 * @Description:
 * @FilePath: /learn-go-with-tests/arrays/sum_test.go
 * @可以输入预定的版权声明、个性签名、空行等
 */
package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	// numbers := [5]int{1, 2, 3, 4, 5}

	// got := Sum(numbers)
	// want := 15

	// if got != want {
	// 	t.Errorf("got %d want %d given, %v", got, want, numbers)
	// }

	t.Run("collection of 5 numbers", func(t *testing.T) {
		// numbers := [5]int{1, 2, 3, 4, 5}
		// 数组换成切片
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	// 不能对切片使用等号运算符
	// if got != want {
	// 	t.Errorf("got %v want %v", got, want)
	// }

	// reflect.DeepEqual 不是「类型安全」
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
