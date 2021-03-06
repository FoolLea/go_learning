/*
 * @Author: your name
 * @Date: 2020-08-03 21:07:52
 * @LastEditTime: 2020-08-05 21:57:05
 * @LastEditors: your name
 * @Description:
 * @FilePath: /learn-go-with-tests/structs/perimeter_test.go
 * @可以输入预定的版权声明、个性签名、空行等
 */
package structs

import (
	"testing"
)

func TestPremiter(t *testing.T) {
	got := Premeter(10.0, 10.0)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	t.Run("rectangle", func(t *testing.T) {
		// got := Area(10.0, 10.0)
		// want := 100.0
		rectangle := Rectangle{10.0, 10.0}
		// got := rectangle.Area(rectangle)
		got := rectangle.Area()
		want := 100.0

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})

	t.Run("circle", func(t *testing.T) {
		circle := Circle{10}
		// got := circle.Area(circle)
		got := circle.Area()
		want := 314.1592653589793

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})
}

func TestArea2(t *testing.T) {
	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	}

	t.Run("rectangle", func(t *testing.T) {
		rectangle := Rectangle{11, 11}
		checkArea(t, rectangle, 121)
	})

	t.Run("circle", func(t *testing.T) {
		circle := Circle{10}
		checkArea(t, circle, 314.1592653589793)
	})

}

func TestArea3(t *testing.T) {
	// areaTests := []struct {
	// 	shpae Shape
	// 	want  float64
	// }{
	// 	{Rectangle{12, 6}, 72.0},
	// 	{Circle{10}, 314.1592653589793},
	// 	{Triangle{12, 6}, 36.0},
	// }

	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		// fmt.Println(tt)
		// got := tt.shape.Area()
		// if got != tt.hasArea {
		// 	t.Errorf("%#v got %.2f want %.2f", tt.shape, got, tt.hasArea)
		// }

		// 在每个用例中使用 t.Run，测试用例的错误输出中会包含用例的名字
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %.2f want %.2f", tt.shape, got, tt.hasArea)
			}
		})
	}
}
