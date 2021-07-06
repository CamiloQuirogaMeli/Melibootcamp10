package main
import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

/* Una tienda de ropa quiere ofrecer a sus clientes un descuento sobre sus productos, para ello
necesitan una aplicación que les permita calcular el descuento en base a 2 variables, su
precio y el descuento en porcentaje. Espera obtener como resultado el valor con el
descuento aplicado y luego imprimirlo en consola.
● Crear la aplicación de acuerdo a los requerimientos. */

func main() {
	descuento()
}

func descuento() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Descuento")
	fmt.Println("Input format: price discount, ex:")
	fmt.Println("-> 100 20")
	fmt.Println("---------------------")
  
	for {
	  fmt.Print("-> ")
	  text, _ := reader.ReadString('\n')
	  // convert CRLF to LF (Only Windows)
	  text = strings.Replace(text, "\n", "", -1)
	  arr := strings.Fields(text)

	  price, _ := strconv.Atoi(arr[0])
	  discount, _ := strconv.Atoi(arr[1])
	  dis := price / 100 * discount
	  fmt.Println("El precio", arr[0], "aplicando el descuento de", arr[1], "se reduce a", dis)
	}
}