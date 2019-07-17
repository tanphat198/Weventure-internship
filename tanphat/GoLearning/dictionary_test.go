package GoLearning

import (
	"testing"
)

func TestSearch(t *testing.T){
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := Search("unknown")
		want := ErrNotFound

		assertError(t, err, want)

		assertStrings(t, err.Error(), Error())
	})
}


func TestAdd(t *testing.T){
	t.Run("New word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "this is just a test"

		err := Add(word,definition)
		assertError(t , err , nil)
		assertDefinition(t, dictionary,word,definition)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		err := Add(word,"new test")

		assertError(t, err, ErrWordExist)
		assertDefinition(t, dictionary, word, definition)
	})
}

func assertDefinition(t *testing.T, dictionary Dictionary, word, definition string){
	t.Helper()

	got, err := Search(word)
	if err != nil {
		t.Fatal("Should find added word: ",err)
	}

	if definition != got {
		t.Errorf("got '%s' but want '%s'",got ,definition)
	}
}

func assertStrings(t *testing.T, got, want string){
	t.Helper()

	if got != want {
		t.Errorf("got '%s' want '%s'",got ,want)
	}
}

func assertError(t *testing.T, got , want error){
	t.Helper()

	if got != want {
		t.Errorf("got error '%s' but want '%s'",got , want)
	}
}