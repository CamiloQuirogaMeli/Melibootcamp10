package main

func main() {

	var apellido string = "Gomez"
	//var edad int = "35". This is wrong, integer values can't be string literals
	var edad int = 35
	boolean := "false" //Althouhg this declaration is ok, it may be misleading, because the variable "boolean" is actually a string
	//var sueldo string = 45857.90. This is wrong because string values have to be wrapped in " "
	var sueldo string = "45857.90"
	var nombre string = "Julián"

	//It is worth noting that, despite being declared correctly, this variables haven't been used, so the compiler is going to throw errors

}
