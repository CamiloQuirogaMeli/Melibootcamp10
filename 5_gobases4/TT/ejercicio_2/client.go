package main

import "errors"

type ClientInfo struct {
	Legajo			int
	Nombre 			string
	Apellido		string
	DNI				int
	NumeroTelefono	int
	Domicilio		string
}

func CreateClient(
	legajo 			int,
	nombre 			string,
	apellido 		string,
	dni		 		int,
	numeroTelefono 	int,
	domicilio		string,
) (ClientInfo, error) {

	if legajo == 0 {

		return ClientInfo{}, errors.New("error: el legajo no puede ser 0")
	} 

	if nombre == "" {

		return ClientInfo{}, errors.New("error: el nombre no puede ser vacio")
	}

	if apellido == "" {

		return ClientInfo{}, errors.New("error: el apellido no puede ser vacio")
	}

	if dni == 0 {

		return ClientInfo{}, errors.New("error: el DNI no puede ser 0")
	}

	if numeroTelefono == 0 {

		return ClientInfo{}, errors.New("error: el numero de telefono no puede ser 0")
	}

	if domicilio == "" {

		return ClientInfo{}, errors.New("error: el domicilio no puede ser vacio")
	}

	return ClientInfo {
		Nombre : nombre,
		Apellido : apellido,
		DNI : dni,
		NumeroTelefono : numeroTelefono,
		Domicilio: domicilio,
	}, nil
}
