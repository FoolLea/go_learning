/*
 * @Author: your name
 * @Date: 2020-08-03 21:11:08
 * @LastEditTime: 2020-08-04 23:48:05
 * @LastEditors: your name
 * @Description:
 * @FilePath: /learn-go-with-tests/structs/premeter.go
 * @可以输入预定的版权声明、个性签名、空行等
 */
package structs

import (
	"math"
)

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Shape interface {
	Area() float64
}

func Premeter(width float64, height float64) float64 {
	return 2 * (width + height)
}

// func Area(width float64, height float64) float64 {
// 	return width * height
// }

// 申明函数
// func Area(rectangle Rectangle) float64 {
// 	return rectangle.Width * rectangle.Height
// }

// 申明方法
// Rectangle 有一个返回值类型为 float64 的方法 Area，所以它满足接口 Shape
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Circle 有一个返回值类型为 float64 的方法 Area，所以它满足接口 Shape
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
