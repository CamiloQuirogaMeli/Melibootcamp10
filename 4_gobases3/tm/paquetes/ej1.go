package paquetes

import (
	"errors"
	"io/ioutil"
)

func GuardarArchivo(path string, productos string) error {
	textToSave := []byte(productos)
	err := ioutil.WriteFile("./file.txt", textToSave, 0644)

	if err != nil {
		return errors.New("No se pudo guardar el archivo")
	} else {
		return nil
	}
}
