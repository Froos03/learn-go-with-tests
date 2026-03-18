package dictionary

import (
	"testing"
)

var dictionary = Dictionary{"Hi": "Hello"}

func TestSearch(t *testing.T) {

	t.Run("Searching for a known word", func(t *testing.T) {
		got, _ := dictionary.Search("Hi")
		want := "Hello"

		assertStrings(t, got, want)
	})
	t.Run("Searching for an unknown word", func(t *testing.T) {
		_, err := dictionary.Search("Bye")
		if err == nil {
			t.Fatal("Expected to get an error")
		}

		assertError(t, err, ErrUnkownWord)
	})
}

func TestAdd(t *testing.T) {
	t.Run("Adding a new word", func(t *testing.T) {
		const word = "Bye"
		err := dictionary.Add(word, "Goodbye")
		want := "Goodbye"

		assertError(t, err, nil)
		assertDefinition(t, word, want)
	})
	t.Run("Adding an existing word should not work", func(t *testing.T) {
		const word = "Hi"
		err := dictionary.Add(word, "Mar7aba")
		want := "Hello"

		assertError(t, err, ErrExistingWord)
		assertDefinition(t, word, want)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Updating an existing word", func(t *testing.T) {
		const word = "Hi"
		err := dictionary.Update(word, "Mar7aba")
		want := "Mar7aba"

		assertError(t, err, nil)
		assertDefinition(t, word, want)
	})
	t.Run("Updating an unknown word", func(t *testing.T) {
		const word = "Boi"
		err := dictionary.Update(word, "Goodbye")

		assertError(t, err, ErrWordDoesNotExist)
	})
}
func TestDelete(t *testing.T) {
	t.Run("Deleting an existing word", func(t *testing.T) {
		const word = "Hi"
		err := dictionary.Delete(word)
		assertError(t, err, nil)

		_, err = dictionary.Search(word)
		assertError(t, err, ErrUnkownWord)
	})
	t.Run("Deleting an unknown word", func(t *testing.T) {
		const word = "Yo"
		err := dictionary.Delete(word)

		assertError(t, err, ErrWordDoesNotExist)
	})
}

func assertDefinition(t *testing.T, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatalf("Should have found the added word: %q", err)
	}
	assertStrings(t, got, definition)
}

func assertStrings(t testing.TB, word, definition string) {
	t.Helper()

	if word != definition {
		t.Errorf("got %q want %q", word, definition)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
