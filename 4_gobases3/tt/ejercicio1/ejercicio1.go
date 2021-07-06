package main

import "fmt"

type Person struct {
	name    string
	surname string
	age     int
	email   string
	pass    string
}

func rename(p *Person, s string, n string) {
	p.surname = s
	p.name = n
}

func changeAge(p *Person, a int) {
	p.age = a
}

func changeEmail(p *Person, e string) {
	p.email = e
}

func changePass(p *Person, pa string) {
	p.pass = pa
}

func main() {
	var (
		name    = "Karen"
		surname = "Cuadrado"
		age     = 31
		email   = "cuadrado.karen@gmail.com"
		pass    = "12345435"
	)

	p1 := Person{
		name:    "nombre",
		surname: "apellido",
		age:     0,
		email:   "correo",
		pass:    "contraseña",
	}
	fmt.Println(p1)
	fmt.Println(&p1.name, &p1.surname, &p1.age, &p1.email, &p1.pass)

	rename(&p1, surname, name)
	fmt.Println(p1)
	fmt.Println(&p1.name, &p1.surname, &p1.age, &p1.email, &p1.pass)

	changeAge(&p1, age)
	fmt.Println(p1)
	fmt.Println(&p1.name, &p1.surname, &p1.age, &p1.email, &p1.pass)

	changeEmail(&p1, email)
	fmt.Println(p1)
	fmt.Println(&p1.name, &p1.surname, &p1.age, &p1.email, &p1.pass)

	changePass(&p1, pass)
	fmt.Println(p1)
	fmt.Println(&p1.name, &p1.surname, &p1.age, &p1.email, &p1.pass)
}
