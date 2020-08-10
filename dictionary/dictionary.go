/*
 * @Author: your name
 * @Date: 2020-08-10 23:25:57
 * @LastEditTime: 2020-08-10 23:46:36
 * @LastEditors: your name
 * @Description:
 * @FilePath: /learn-go-with-tests/dictionary/dictionary.go
 * @可以输入预定的版权声明、个性签名、空行等
 */
package dictionary

type Dictionary map[string]string

func Search(dictionary map[string]string, word string) string {
	return dictionary[word]
}

func (d Dictionary) Search(word string) string {
	return d[word]
}
