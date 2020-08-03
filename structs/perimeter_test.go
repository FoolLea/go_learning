/*
 * @Author: your name
 * @Date: 2020-08-03 21:07:52
 * @LastEditTime: 2020-08-03 23:47:06
 * @LastEditors: your name
 * @Description:
 * @FilePath: /learn-go-with-tests/structs/perimeter_test.go
 * @可以输入预定的版权声明、个性签名、空行等
 */
package structs

import "testing"

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
		got := rectangle.Area(rectangle)
		want := 100.0

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})

	t.Run("circle", func(t *testing.T) {
		circle := Circle{10}
		got := circle.Area(circle)
		want := 314.16

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})
}
