/*
 * @Author: your name
 * @Date: 2020-08-10 23:22:49
 * @LastEditTime: 2020-08-19 00:25:48
 * @LastEditors: your name
 * @Description:
 * @FilePath: /learn-go-with-tests/dictionary/dictionary_test.go
 * @可以输入预定的版权声明、个性签名、空行等
 */
package dictionary

import (
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := map[string]string{"test": "this is a test"}
	got := Search(dictionary, "test")
	want := "this is a test"

	assertString(t, got, want)
}

func assertString(t *testing.T, got string, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got error '%s' want '%s'", got, want)
	}
}

func TestSearch2(t *testing.T) {
	dictionary := Dictionary{"test": "this is a test"}

	t.Run("know word", func(t *testing.T) {
		// 返回第二个参数，它是一个 Error 类型
		got, _ := dictionary.Search("test")
		want := "this is a test"

		assertString(t, got, want)
	})

	t.Run("unknow word", func(t *testing.T) {
		// 返回第二个参数，它是一个 Error 类型
		_, got := dictionary.Search("unknown")
		// want := "could not find the word you were looking for"

		// if err == nil {
		// 	t.Fatal("expected to get an error.")
		// }

		// assertString(t, err.Error(), want)
		assertError(t, got, ErrNotFound)
	})

}

func TestAdd(t *testing.T) {
	// dictionary := Dictionary{}
	// dictionary.Add("test", "this is just a test")

	// want := "this is just a test"
	// got, err := dictionary.Search("test")
	// if err != nil {
	// 	t.Fatal("should find added word:", err)
	// }

	// if want != got {
	// 	t.Errorf("got '%s' want '%s'", got, want)
	// }

	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "this is just a test"

		err := dictionary.Add(word, definition)

		assertError(t, err, nil)

		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, "new test")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})
}

func assertDefinition(t *testing.T, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}

	if definition != got {
		t.Errorf("got '%s' want '%s'", got, definition)
	}
}
