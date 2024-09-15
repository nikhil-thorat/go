package maps

type Dictionary map[string]string

const (
	ErrWrongKey        = DictionaryErr("wrong key")
	ErrWordExists      = DictionaryErr("duplicate key error")
	ErrKeyDoesNotExist = DictionaryErr("cannot update key that does not exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(key string) (string, error) {

	definition, ok := d[key]

	if !ok {
		return "", ErrWrongKey
	}

	return definition, nil
}

func (d Dictionary) Add(key string, definition string) error {
	_, err := d.Search(key)

	switch err {
	case ErrWrongKey:
		d[key] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(key string, definition string) error {

	_, err := d.Search(key)

	switch err {
	case ErrWrongKey:
		return ErrKeyDoesNotExist
	case nil:
		d[key] = definition
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(key string) {
	delete(d, key)
}
