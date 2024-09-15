package maps

import "testing"

func TestSearch(t *testing.T) {

	// t.Run("Known key", func(t *testing.T) {
	// 	got, _ := Dictionary.Search(Dictionary{}, "test")
	// 	want := "This is just a test"
	// 	assertString(t, got, want)
	// })

	t.Run("Unknown key", func(t *testing.T) {
		_, got := Dictionary.Search(Dictionary{}, "unknown")

		if got == nil {
			t.Fatal("Expected an error")
		}

		assertError(t, got, ErrWrongKey)

	})
}

func TestAdd(t *testing.T) {
	dictionary := Dictionary{}

	key := "test"
	definition := "This is just a test"

	dictionary.Add(key, definition)

	assertDefinition(t, dictionary, key, "This is just a test")

	t.Run("New Word", func(t *testing.T) {
		dictionary := Dictionary{}

		key := "test"
		definition := "This is just a test"

		err := dictionary.Add(key, definition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, key, definition)
	})

	t.Run("Existing Word", func(t *testing.T) {
		key := "test"
		definition := "This is just a test"

		dictionary := Dictionary{key: definition}

		err := dictionary.Add(key, definition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, key, definition)
	})
}

func TestUpdate(t *testing.T) {

	t.Run("Existing Word", func(t *testing.T) {
		key := "test"
		definition := "This is just a test"

		dictionary := Dictionary{key: definition}
		newDefinition := "This is just a test, but with a new definition"

		err := dictionary.Update(key, newDefinition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, key, newDefinition)
	})

	t.Run("New Word", func(t *testing.T) {
		key := "new_test"
		definition := "This is just a test"

		dictionary := Dictionary{}

		err := dictionary.Update(key, definition)

		assertError(t, err, ErrKeyDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	key := "test"
	dictionary := Dictionary{key: "This is just a test"}

	dictionary.Delete(key)

	_, err := dictionary.Search(key)

	assertError(t, err, ErrWrongKey)
}

func assertString(t *testing.T, got string, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func assertError(t *testing.T, got error, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertDefinition(t *testing.T, dictionary Dictionary, key string, definition string) {
	t.Helper()

	got, err := dictionary.Search(key)

	if err != nil {
		t.Fatal("Should find the definition", err)
	}

	assertString(t, got, definition)
}
