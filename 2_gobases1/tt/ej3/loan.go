package ej3

const(
	msg1 = "Debes tener mas de 22 años para acceder a un prestamo\n"
	msg2 = "Debes tener mas de un año de antiguedad en tu trabajo para acceder a un prestamo\n"
	msg3 = "Se te otorga el prestamo pero debes pagar intereses porque tu salario es inferior o igual a $100.000\n"
	msg4 = "Se te otorga el prestamo, no pagas intereses porque tu salario es mayor a $100.000\n"
)

func BankLoan(age int, seniorityInMonths int, salary float32) string {
	if age < 22 {
		return msg1
	}else if seniorityInMonths < 12 {
		return msg2
	}else if salary < 100000 {
		return msg3
	}else{
		return msg4
	}
}
