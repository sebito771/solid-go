package main

type diccionario map[string]string
type DictionaryErr string

func Search(diccionario map[string]string, word string) string {
	return diccionario[word]
}

const (
	ErrorNotFound         = DictionaryErr("palabra no encontrada")
	ErrorWordExist        = DictionaryErr("la palabra ya existe")
	ErrorWordDoesNotExist = DictionaryErr("no se puede realizar el cambio porque la palabra no existe")
)

// FUNCIONES
func (dic diccionario) Search(word string) (string, error) {
	definition, ok := dic[word]
	if !ok {
		return "no se ha encontrado la palabra", ErrorNotFound
	}
	return definition, nil
}

func (d diccionario) Add(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrorNotFound:
		d[word] = definition
	case nil:
		return ErrorWordExist
	default:
		return err

	}
	return nil
}

func (dice DictionaryErr) Error() string {
	return string(dice)
}

func (d diccionario) Update(word, definition string) error {
	d[word] = definition
	_, err := d.Search(word)

	switch err {
	case ErrorNotFound:
		return ErrorWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}
	return nil
}

func (d diccionario) Delete(word string) error {
	_, err := d.Search(word)
	switch err {
	case ErrorNotFound:
		return ErrorWordDoesNotExist
	case nil:
		delete(d, word)
	default:
		return err
	}
	return nil
}
