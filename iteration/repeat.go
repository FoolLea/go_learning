/*
 * @Author: your name
 * @Date: 2020-07-27 00:55:52
 * @LastEditTime: 2020-07-27 00:58:05
 * @LastEditors: your name
 * @Description:
 * @FilePath: /learn-go-with-tests/iteration/repeat.go
 * @可以输入预定的版权声明、个性签名、空行等
 */
package iteration

func Repeat(character string) string {
	var repeated string
	for i := 0; i < 5; i++ {
		repeated += character
	}
	return repeated
}
