/*
 * @Author: your name
 * @Date: 2020-08-03 21:07:52
 * @LastEditTime: 2020-08-05 21:21:14
 * @LastEditors: your name
 * @Description:
 * @FilePath: /learn-go-with-tests/structs/perimeter_test.go
 * @可以输入预定的版权声明、个性签名、空行等
 */
package structs

import (
	"fmt"
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
	areaTests := []struct {
		shpae Shape
		want  float64
	}{
		{Rectangle{12, 6}, 72.0},
		{Circle{10}, 314.1592653589793},
	}

	for _, tt := range areaTests {
		fmt.Println(tt)
		got := tt.shpae.Area()
		if got != tt.want {
			t.Errorf("got %.2f want %.2f", got, tt.want)
		}
	}
}
