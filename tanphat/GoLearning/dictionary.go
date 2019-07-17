package GoLearning

type Dictionary map[string]string

const (
	ErrNotFound = DictionaryErr("could not find the word you looking for")
	ErrWordExist = DictionaryErr("cannot add word because it already exists")
)

type DictionaryErr string

func (e DictionaryErr) Error() string{
	return string(e)
}

func Search(dictionary map[string]string, word string) string {
	return dictionary[word]
}

func (d Dictionary) Add (word string , definition string) error{
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExist
	default :
		return err
	}

	return nil
}

func (d Dictionary) Search(word string) (string,error){
	definition, ok := d[word]
	if !ok{
		return "", ErrNotFound
	}
	return definition, nil
}



