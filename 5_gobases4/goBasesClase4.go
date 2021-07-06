package main

import (
	"errors"
	"fmt"
	//"os"
	//"time"
)

/*
//declaracion de la interface error
type error interface {
	Error() string
}

// sólo se requiere crear un tipo que implemente el método Error()
type myCustomError struct {
	status int
	msg    string
}

//hacemos que nuestro tipo struct implemente el método Error()
func (e *myCustomError) Error() string {
	return fmt.Sprintf("%d - %v", e.status, e.msg)
}
func myCustomErrorTest(status int) (int, error) { //devolvemos un objeto de la interface error, que está siendo creado en el momento
	if status >= 300 {
		return 400, &myCustomError{ //el & es para usarlo fuera de la funcion y manipular el objeto
			status: status,
			msg:    "algo salió mal",
		}
	}
	return 200, nil
}
func main() {
	//error
	status, err := myCustomErrorTest(200) //llamamos a nuestra func
	if err != nil {                       //hacemos una validación del valor de err
		fmt.Println(err) //si err no es nil, imprimimos el error y...
		os.Exit(1)       //utilizamos este método para salir del programa
	}
	fmt.Printf("Status %d, Funciona!\n", status)

	//errors.New
	statusCode := 398
	if statusCode > 399 {
		fmt.Println(errors.New("La petición ha fallado."))
		return
	}
	fmt.Println("El programa finalizó correctamente.")

	//fmt.Errorf
	ahora := time.Now()
	eror := fmt.Errorf("momento del error: %v", ahora)
	fmt.Println("error ocurrido: ", eror)

}
*/

//PAQUETE ERRORS - AS() func As(err error, target interface{}) bool
//errors.As solo funciona con una interface o con un tipo que implementa error (no con error en forma directa)
/*
// definimos un type struct
type myError struct {
	msg string
	x   string
}

//hacemos que nuestro type struct implemente el método Error()
func (e *myError) Error() string {
	return fmt.Sprintf("ha ocurrido un error: %s, %s", e.msg, e.x)
}
func main() {
	e := &myError{"nuevo error", "404"} //& refiere a la direccion de memoria a la interface
	var err *myError
	isMyError := errors.As(e, &err) // compara los tipo de dato de los errores
	fmt.Println(isMyError)          //imprime true porque los errores coinciden, ambos son de tipo: myErrors
}
*/
/*
//PAQUETE ERRORS - IS() func Is(err, target error) bool
var err1 = errors.New("error número 1")

func x() error {
	return fmt.Errorf("información extra del error: %w", err1)
}

func main() {
	e := x()
	coincidence := errors.Is(e, err1)
	fmt.Println(coincidence) //imprime true
}
*/

//PAQUETE ERRORS - UNWRAP()
/*
type errorTwo struct{}

func (e errorTwo) Error() string {
	return "error two happended"
}
func main() {
	e2 := errorTwo{}
	e1 := fmt.Errorf("e2: %w", e2)
	e3 := fmt.Errorf("e1: %w", e1)
	fmt.Println(errors.Unwrap(e1)) //imprime e2
	fmt.Println(errors.Unwrap(e2)) //imprime nil
	fmt.Println(errors.Unwrap(e3)) //imprime e2: error two happened
}
*/

//RETORNO DE ERRORES
func SayHello(name string) (string, error) {
	if name == "" {
		return "", errors.New("nil name")
	}
	return fmt.Sprintf("Hola %s", name), nil
}

func clase4TM() {
	name := ""
	greeting, err := SayHello(name)
	//validamos si hubo error
	if err != nil {
		fmt.Println("no se puede saludar", err)
	}
	//si no hubo error, continuamos con nuestra ejecucion
	fmt.Println(greeting)
}
