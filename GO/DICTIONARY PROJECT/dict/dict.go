package dict

import "errors"

// Dictionary type
type Dictionary map[string]string // "alias". 단순히 별명이라고 볼 수 있음

var (
	errNotFound   = errors.New("Not Found")
	errDuplicated = errors.New("Duplicated Word")
	errCantUpdate = errors.New("Can't update non-existing word")
	errCantDelete = errors.New("Can't delete non-existing word")
)

// type은 method를 가질 수 있다 which is super cool.

// method of the map
// i, ok := m["route"]
// Search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}

// Add a word to the dictionary
func (d Dictionary) Add(word string, def string) error {
	// we don't need the * on the receiver because maps on Go are automatically using *
	// dictionary is a hashmap, and by default a hashmap already has * included
	_, err := d.Search(word)
	if err == errNotFound {
		d[word] = def
		return nil
	}
	return errDuplicated
}

// Update the definition of the word
func (d Dictionary) UpdateDef(word string, def string) error {
	_, err := d.Search(word)
	if err == errNotFound {
		return errCantUpdate
	}
	d[word] = def
	return nil
}

// Delete the word
func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)
	if err == errNotFound {
		return errCantDelete
	}
	delete(d, word)
	return nil
}
