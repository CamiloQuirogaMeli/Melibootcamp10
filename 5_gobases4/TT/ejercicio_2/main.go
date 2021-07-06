package main

import "fmt"
import "io/ioutil"

func RegisterClient(
	legajo 			int,
	nombre 			string,
	apellido 		string,
	dni		 		int,
	numeroTelefono 	int,
	domicilio		string,
) {
	defer func(){
		err := recover()
		fmt.Println(err)
		fmt.Println("Fin de la ejecución")
	}()

	clientToRegister, err := CreateClient(legajo, nombre, apellido, dni, numeroTelefono, domicilio)

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Se registró el client:", clientToRegister)
	}
}


func main() {

	
	defer func() {

		err := recover()

		if err != nil {
			fmt.Println(err)
			RegisterClient(1293, "ruperto", "almeja", 19292, 19288191, "")
			fmt.Println("Se detectaron varios errores en tiempo de ejecución")
		}
		fmt.Println("No han quedado archivos abiertos")
	}()

	content, err := ioutil.ReadFile("customers.txt")

	if err != nil {
		panic("error: el archivo indicado no fue encontrado o está dañado")
	} else {
		fmt.Printf("Contenido del archivo: %s\n", string(content))
	}

}