/*
 * @Author: your name
 * @Date: 2020-07-27 00:52:21
 * @LastEditTime: 2020-08-03 21:29:05
 * @LastEditors: your name
 * @Description:
 * @FilePath: /learn-go-with-tests/iteration/repeat_test.go
 * @可以输入预定的版权声明、个性签名、空行等
 */
package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected '%q' but got '%q'", expected, repeated)
	}
}

// testing.B 可使你访问隐性命名（cryptically named）b.N.
// 基准测试运行时，代码会运行 b.N 次，并测量需要多长时间
// go test -bench=.
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
