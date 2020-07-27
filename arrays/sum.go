/*
 * @Author: your name
 * @Date: 2020-07-27 01:10:05
 * @LastEditTime: 2020-07-27 22:45:10
 * @LastEditors: your name
 * @Description:
 * @FilePath: /learn-go-with-tests/arrays/sum.go
 * @可以输入预定的版权声明、个性签名、空行等
 */
package main

// func Sum(numbers [5]int) (sum int) {
// 	for i := 0; i < 5; i++ {
// 		sum += numbers[i]
// 	}
// 	return sum
// }

// func Sum(numbers [5]int) int {
// 	sum := 0
// 	for i := 0; i < 5; i++ {
// 		sum += numbers[i]
// 	}
// 	return sum
// }

/**
 * @description: range 会迭代数组，每次迭代都会返回数组元素的索引和值
				 使用 _ 空白标志符 来忽略索引
 * @param {type}
 * @return:
*/
func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}

	return sum
}

// func SumAll(numbersToSum ...[]int) (sums []int) {
// 	lengthOfNumbers := len(numbersToSum)

// 	// make 可以在创建切片的时候指定我们需要的长度和容量
// 	sums = make([]int, lengthOfNumbers)

// 	for i, numbers := range numbersToSum {
// 		sums[i] = Sum(numbers)
// 	}

// 	return
// }

// func SumAll(numbersToSum ...[]int) []int {
// 	lengthOfNumbers := len(numbersToSum)

// 	// make 可以在创建切片的时候指定我们需要的长度和容量
// 	sums := make([]int, lengthOfNumbers)

// 	for i, numbers := range numbersToSum {
// 		sums[i] = Sum(numbers)
// 	}

// 	return sums
// }

func SumAll(numbersToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numbersToSum {
		// append 函数，它能为切片追加一个新值
		sums = append(sums, Sum(numbers))
	}

	return sums
}
