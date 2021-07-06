package paquetes

import (
	"errors"
	"io/ioutil"
	"strings"
)

func LeerArchivo(path string) ([]byte, error) {

	data, err := ioutil.ReadFile(path)

	if err != nil {
		return nil, errors.New("No se pudo leer el archivo")
	}

	return data, nil
}

func ReemplazarChars(str string) (string, string) {
	head := "ID\t\tPrecio\t\tCantidad\n"

	body := strings.ReplaceAll(str, ",", "\t\t")
	body = strings.ReplaceAll(body, ";", "\n")

	return head, body

}
