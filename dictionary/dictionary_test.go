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

func TestSearch2(t *testing.T) {
	dictionary := Dictionary{"test": "this is a test"}

	t.Run("know word", func(t *testing.T) {
		// 返回第二个参数，它是一个 Error 类型
		got, _ := dictionary.Search("test")
		want := "this is a test"

		assertString(t, got, want)
	})

	t.Run("nnknow word", func(t *testing.T) {
		// 返回第二个参数，它是一个 Error 类型
		_, err := dictionary.Search("unknown")
		want := "could not find the word you were looking for"

		if err == nil {
			t.Fatal("expected to get an error.")
		}

		assertString(t, err.Error(), want)
	})

}
