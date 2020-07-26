/*
 * @Descripttion:
 * @version:
 * @Author: sueRimn
 * @Date: 2020-07-25 14:42:23
 * @LastEditors: your name
 * @LastEditTime: 2020-07-25 20:24:23
 */
package main

import (
	"fmt"
)

const eHelloPrefix = "Hello, "
const spanish = "Spanish"
const sHelloPrefix = "Hola, "
const french = "French"
const fHelloPrefix = "Hola, "

/**
 * @description:
 * @param {type}
 * @return:
 */
func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	// if language == spanish {
	// 	return sHelloPrefix + name
	// }

	// prefix := eHelloPrefix

	// switch language {
	// case french:
	// 	prefix = fHelloPrefix
	// case spanish:
	// 	prefix = sHelloPrefix
	// }

	return greetingPrfix(language) + name
}

func greetingPrfix(language string) (prefix string) {
	switch language {
	case french:
		prefix = fHelloPrefix
	case spanish:
		prefix = sHelloPrefix
	default:
		prefix = eHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("world", "Enlish"))
}
