package main

import (
  "fmt"
  "github.com/extmatperez/meli_bootcamp10/tm/ej1"
  "github.com/extmatperez/meli_bootcamp10/tm/ej2"
  "github.com/extmatperez/meli_bootcamp10/tm/ej3"
  "github.com/extmatperez/meli_bootcamp10/tm/ej4"
  ej12 "github.com/extmatperez/meli_bootcamp10/tt/ej1"
  ej22 "github.com/extmatperez/meli_bootcamp10/tt/ej2"
  ej32 "github.com/extmatperez/meli_bootcamp10/tt/ej3"
  ej42 "github.com/extmatperez/meli_bootcamp10/tt/ej4"
  "github.com/extmatperez/meli_bootcamp10/tt/ej5"
  "github.com/extmatperez/meli_bootcamp10/tt/ej6"
  "strconv"
)

const separator = "-------------------------------------------------------------------------"

func main(){
  fmt.Println("1. " + separator)
  ej1.PrintMyName()

  fmt.Println("2. " + separator)
  ej2.Weather()

  fmt.Println("3. " + separator)
  ej3.Variables()

  fmt.Println("4. " + separator)
  ej4.DataTypes()

  fmt.Println("1. " + separator)
  lenWord, letter := ej12.LetterOfWord("RANDOM")
  fmt.Printf("length %d, %s\n", lenWord, letter)

  fmt.Println("2. " + separator)
  fmt.Println("Price " + strconv.Itoa(200000) + " Discount 30%\nTotal Price " + fmt.Sprintf("%.0f", ej22.CalculateDiscount(200000.0,0.3)))

  fmt.Println("3. " + separator)
  fmt.Println(">Juan 21 años de edad, antiguedad en el trabajo de 24 meses, salario de 100000\n" + ej32.BankLoan(21,24,100000))

  fmt.Println("4. " + separator)
  for i := uint8(1); i <= 12; i++ {
    fmt.Println(ej42.NumToMonth(i))
  }
  fmt.Println("5. " + separator)
  students := ej5.Students{NamesList: []string{"Benjamin", "Nahuel", "Brenda", "Marcos", "Pedro", "Axel", "Alez", "Dolores", "Federico", "Hernan","Leandro", "Eduardo", "Duvraschka"}}
  students.AddStudent("Marlon")
  students.AddStudent("Andres")
  students.AddStudent("Gabriela")
  fmt.Println(students.GetList())

  fmt.Println("6. " + separator)
  employees := ej6.Employees{EmployeesList:map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}}
  fmt.Println(employees.GetEmployeeByName("Benjamin"))
  fmt.Println(employees.EmployeesOver21())
  employees.AddEmployee("Federico",25)
  employees.DeleteEmployee("Pedro")
  fmt.Println(employees.GetEmployees())

}
