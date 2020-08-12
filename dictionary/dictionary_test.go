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
	dictionary := Dictionary{}
	dictionary.Add("test", "this is just a test")

	want := "this is just a test"
	got, err := dictionary.Search("test")
	if err != nil {
		t.Fatal("should find added word:", err)
	}

	if want != got {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
