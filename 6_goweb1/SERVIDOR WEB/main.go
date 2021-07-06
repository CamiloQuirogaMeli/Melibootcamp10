package main

/*
1
Según la temática elegida, genera un JSON que cumpla con las siguientes claves según la temática.
Los productos varían por id, nombre, color, precio, stock, código (alfanumérico), publicado (si-no), fecha de creación.
Los usuarios varían por id, nombre, apellido, email, edad, altura, activo (si-no), fecha de creación.
Las transacciones: id, código de transacción (alfanumérico), moneda, monto, emisor (string), receptor (string), fecha de transacción.
Dentro de la carpeta go-web crea un archivo temática.json, el nombre tiene que ser el tema elegido, ej: products.json.
Dentro del mismo escribí un JSON que permita tener un array de productos, usuarios o transacciones con todas sus variantes.

2
Crea dentro de la carpeta go-web un archivo llamado main.go
Crea un servidor web con Gin que te responda un JSON que tenga una clave “message” y diga Hola seguido por tu nombre.
Pegale al endpoint para corroborar que la respuesta sea la correcta.

3
Ya habiendo creado y probado nuestra API que nos saluda, generamos una ruta que devuelve un listado de la temática elegida.
Dentro del main.go, crea una estructura según la temática con los campos correspondientes.
Genera un endpoint cuya ruta sea /temática (en plural)
Genera un handler para el endpoint llamado GetAll
Instanciar un array de la estructura y devolverla en nuestro endpoint.
*/
/*
import (
	"fmt"

	"github.com/gin-gonic/gin"
)
*/
func main() {
	/*
		router := gin.Default()
		router.GET("/hello-world", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Hello World!, que bueno crear este servidor web",
			})
		})
		router.Run()
	*/
	/*
		router := gin.Default()
		//Cada vez que llamamos a GET y le pasamos una ruta, definimos un nuevo endpoint.
		router.GET("/", HandlerRaiz)
		router.GET("/gophers", HandlerGophers)
		router.GET("/gophers/get", HandlerGetGopher)
		router.GET("/gophers/info", HandlerGetInfo)
		router.GET("/about", HandlerAbout)
		router.Run(":8081")
	*/
}

/*
//Esta función nos permite ver la anatomía de un mensaje Request de una
//forma más práctica.
func Ejemplo(context *gin.Context) {
	//El body, header y method están contenidos en el contexto de gin.
	contenido := context.Request.Body
	header := context.Request.Header
	metodo := context.Request.Method
	fmt.Println("¡He recibido algo!")
	fmt.Printf("\tMetodo: %s\n", metodo)
	fmt.Printf("\tContenido:\n")
	for key, value := range header {
		fmt.Printf("\t\t%s -> %s\n", key, value)
	}
	fmt.Printf("\tCotenido:%s\n", contenido)
	fmt.Println("¡Yay!")
	context.String(200, "¡Lo recibí!") //Respondemos al cliente con 200 OK y un bonito mensaje.
}
*/
