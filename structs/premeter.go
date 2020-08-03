/*
 * @Author: your name
 * @Date: 2020-08-03 21:11:08
 * @LastEditTime: 2020-08-03 23:48:36
 * @LastEditors: your name
 * @Description:
 * @FilePath: /learn-go-with-tests/structs/premeter.go
 * @可以输入预定的版权声明、个性签名、空行等
 */
package structs

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
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
func (r Rectangle) Area() float64 {
	return 0
}

func (c Circle) Area() float64 {
	return 0
}
