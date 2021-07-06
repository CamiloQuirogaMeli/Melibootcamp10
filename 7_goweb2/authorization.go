package main

import "errors"

func ValidateToken(token string) (bool, error) {
	if token != "123456" {
		return false, errors.New("no tiene permisos para realizar la petición solicitada")
	}
	return true, nil
}
