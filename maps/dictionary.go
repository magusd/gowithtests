package maps

type DictionaryErr string

func (d DictionaryErr) Error() string {
	return string(d)
}

const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("can't add a word that already exists")
	ErrWordDoesNotExist = DictionaryErr("can't update a word that doesn't exist")
	ErrDeleteInvalid    = DictionaryErr("can't delete a word that doesn't exist")
)

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	}
	return err
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
		return nil
	}
	return err
}

func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		return ErrDeleteInvalid
	case nil:
		delete(d, word)
		return nil
	}
	return err
}
