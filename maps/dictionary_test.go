package maps

import (
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("searching", func(t *testing.T) {
		assertDefinition(t, dictionary, "test", "this is just a test")
	})

	t.Run("unkown word", func(t *testing.T) {
		_, got := dictionary.Search("unkown")
		assertError(t, got, ErrNotFound)
	})
}

func ExampleDictionary_Search() {
	d := Dictionary{"test": "this is just a test"}
	definition, err := d.Search("test")
	if err == ErrNotFound {
		fmt.Println("word not found")
	}
	fmt.Println(definition)
	//Output: this is just a test
}

func ExampleDictionary_Add() {
	d := Dictionary{}
	err := d.Add("test", "definition")
	if err == ErrWordExists {
		fmt.Println("cant override words with Add")
	}
	w, _ := d.Search("test")
	fmt.Println(w)
	//Output: definition
}

func ExampleDictionary_Update() {
	d := Dictionary{"test": "definition"}
	err := d.Update("test", "this is just a test")
	if err == ErrWordDoesNotExist {
		fmt.Println("can't update a word that doesn't exist")
	}
	definition, err := d.Search("test")
	if err == ErrNotFound {
		fmt.Println("word not found")
	}
	fmt.Println(definition)
	//Output: this is just a test
}

func ExampleDictionary_Delete() {
	d := Dictionary{"test": "definition"}
	err := d.Delete("test")
	if err == ErrDeleteInvalid {
		fmt.Println("can't delete word that doesn't exist")
	}
	_, err = d.Search("test")
	if err == ErrNotFound {
		fmt.Println("word not found")
		//Output: word not found
	}

}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		dictionary.Add("test", "this is just a test")
		assertDefinition(t, dictionary, "test", "this is just a test")
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

func TestUpdate(t *testing.T) {
	t.Run("update existing word", func(t *testing.T) {
		word := "test"
		dictionary := Dictionary{word: "this is just a test"}
		err := dictionary.Update(word, "new test")
		assertNoError(t, err)
		assertDefinition(t, dictionary, word, "new test")
	})
	t.Run("update non existing word", func(t *testing.T) {
		dictionary := Dictionary{}
		err := dictionary.Update("none", "test")
		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	t.Run("deleting a word", func(t *testing.T) {
		word := "test"
		dictionary := Dictionary{word: "this is just a test"}
		dictionary.Delete(word)
		_, err := dictionary.Search(word)
		assertError(t, err, ErrNotFound)
	})
	t.Run("delete invalid", func(t *testing.T) {
		dictionary := Dictionary{}
		err := dictionary.Delete("invalid")
		assertError(t, err, ErrDeleteInvalid)
	})
}

func assertDefinition(t *testing.T, dictionary Dictionary, word string, want string) {
	t.Helper()
	got, err := dictionary.Search("test")
	if err != nil {
		t.Fatal("Shouldn't receive an error, got: ", err)
	}
	assertStrings(t, got, want)
}

func assertError(t *testing.T, err error, want error) {
	t.Helper()
	if err != want {
		t.Fatal("expected error: ", want, "but got", err)
	}
}
func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal("didn't expect an error.")
	}
}

func assertStrings(t *testing.T, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q, given %q", got, want, "test")
	}
}
