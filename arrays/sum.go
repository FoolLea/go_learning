/*
 * @Author: your name
 * @Date: 2020-07-27 01:10:05
 * @LastEditTime: 2020-07-27 01:12:54
 * @LastEditors: your name
 * @Description:
 * @FilePath: /learn-go-with-tests/arrays/sum.go
 * @可以输入预定的版权声明、个性签名、空行等
 */
package main

func Sum(numbers [5]int) (sum int) {
	for i := 0; i < 5; i++ {
		sum += numbers[i]
	}
	return sum
}

// func Sum(numbers [5]int) int {
// 	sum := 0
// 	for i := 0; i < 5; i++ {
// 		sum += numbers[i]
// 	}
// 	return sum
// }
