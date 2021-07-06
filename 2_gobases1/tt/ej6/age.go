package ej6

type Employees struct{
	EmployeesList map[string]int
}

func (e *Employees) GetEmployees() map[string]int {
	return e.EmployeesList
}

func (e *Employees) GetEmployeeByName(name string) (string, int) {
	return name, e.EmployeesList[name]
}

func (e *Employees) AddEmployee(name string, age int)  {
	e.EmployeesList[name] = age
}

func (e *Employees) CountEmployees() int {
	return len(e.EmployeesList)
}

func (e *Employees) DeleteEmployee(name string)  {
	delete(e.EmployeesList, name)
}

func (e *Employees) EmployeesOver21() map[string]int {
	emp := make(map[string]int)
	for name, age := range e.EmployeesList{
		if age > 21{
			emp[name] = age
		}
	}
	return emp
}