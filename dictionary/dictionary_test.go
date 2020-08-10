package dictionary

import "testing"

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
	got := dictionary.Search("test")
	want := "this is a test"

	assertString(t, got, want)
}
