package main

import "testing"

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func assertError(t testing.TB, err, want error) {
	t.Helper()
	if err != want {
		t.Errorf("got %s, want %s", err, want)
	}
}

func assertDefinition(t testing.TB, dict Dictionary, word, definition string) {
	t.Helper()
	got, err := dict.Search(word)
	if err != nil {
		t.Fatal("Should find added word", err)
	}
	assertStrings(t, got, definition)
}
func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is a test"}
	t.Run("Found word", func(t *testing.T) {

		got, _ := dictionary.Search("test")
		want := "this is a test"
		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {

		_, err := dictionary.Search("unknown")
		if err == nil {
			t.Fatal("expected to get an error.")
		}
		assertError(t, err, ErrWordNotFound)
	})
}
func TestAdd(t *testing.T) {
	t.Run("Add new word", func(t *testing.T) {
		d := Dictionary{}
		word := "new"
		definition := "new word test"
		err := d.Add(word, definition)
		assertError(t, err, nil)
		assertDefinition(t, d, word, definition)
	})
	t.Run("existing word", func(t *testing.T) {
		word := "existing"
		definition := "existing word test"
		d := Dictionary{word: definition}
		err := d.Add(word, "new word now")
		assertError(t, err, ErrWordAlreadyExists)
		assertDefinition(t, d, word, definition)
	})
}
func TestUpdate(t *testing.T) {
	t.Run("update word", func(t *testing.T) {
		word := "update"
		definition := "def for update"
		dict := Dictionary{word: definition}
		newDef := "new def for update"
		err := dict.Update(word, newDef)
		assertError(t, err, nil)
		assertDefinition(t, dict, word, newDef)
	})
	t.Run("update word doesn't exist", func(t *testing.T) {
		word := "update"
		dict := Dictionary{}
		newDef := "new def for update"
		err := dict.Update(word, newDef)
		assertError(t, err, ErrWordDoesNotExist)
	})
}
func TestDelete(t *testing.T) {
	t.Run("delete", func(t *testing.T) {
		word := "test"
		dict := Dictionary{word: "delete word"}
		dict.Delete(word)
		_, err2 := dict.Search(word)
		assertError(t, err2, ErrWordNotFound)
	})
}
