/*
 * @Author: your name
 * @Date: 2020-08-10 23:25:57
 * @LastEditTime: 2020-08-12 22:21:50
 * @LastEditors: your name
 * @Description:
 * @FilePath: /learn-go-with-tests/dictionary/dictionary.go
 * @可以输入预定的版权声明、个性签名、空行等
 */
package dictionary

import (
	"errors"
)

var ErrNotFound = errors.New("could not find the word you were looking for")

type Dictionary map[string]string

func Search(dictionary map[string]string, word string) string {
	return dictionary[word]
}

// func (d Dictionary) Search(word string) string {
// 	return d[word]
// }

func (d Dictionary) Search(word string) (string, error) {
	// map 查找的特性，返回两个值。第二个值是一个布尔值，表示是否成功找到 key
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(word, definition string) {
	d[word] = definition
}
