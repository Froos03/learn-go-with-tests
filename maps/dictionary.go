package dictionary

type Dictionary map[string]string

type DictionaryErr string

const (
	ErrUnkownWord       = DictionaryErr("We could not find the word you were looking for")
	ErrExistingWord     = DictionaryErr("This is an existing word, cannot add")
	ErrWordDoesNotExist = DictionaryErr("Cannot Perform Operation on non-existing word")
)

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(keyword string) (string, error) {
	definition, ok := d[keyword]

	if !ok {
		return "", ErrUnkownWord
	}
	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrUnkownWord:
		d[word] = definition
	case nil:
		return ErrExistingWord
	default:
		return err
	}
	return nil
}

func (d Dictionary) Update(word, newDefinition string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		d[word] = newDefinition
	case ErrUnkownWord:
		return ErrWordDoesNotExist
	default:
		return err
	}
	return nil
}

func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		delete(d, word)
	case ErrUnkownWord:
		return ErrWordDoesNotExist
	default:
		return err
	}
	return nil
}
