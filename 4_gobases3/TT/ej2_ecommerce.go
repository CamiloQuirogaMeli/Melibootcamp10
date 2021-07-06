package main

/*
Una importante empresa de ventas web necesita agregar una funcionalidad para agregar
productos a los usuarios. Para ello requieren que tanto los usuarios como los productos
tengan la misma dirección de memoria en el main del programa como en las funciones.
Se necesitan las estructuras:
- Usuario: Nombre, Apellido, Correo, Productos (array de productos).
- Producto: Nombre, precio, cantidad.
Se requieren las funciones:
- Nuevo producto: recibe nombre y precio, y retorna un producto.
- Agregar producto: recibe usuario, producto y cantidad, no retorna nada, agrega el
producto al usuario.
- Borrar productos: recibe un usuario, borra los productos del usuario.
*/
import "fmt"

type UsuarioEcom struct {
	NombreUsEcom   string
	ApellidoUEcom  string
	CorreoUEcom    string
	ProductosUEcom []*ProductosEcom
}

type ProductosEcom struct {
	NombrePEcom   string
	PrecioPEcom   float64
	CantidadPEcom int
}

func nuevoProducto(nombre string, precio float64) *ProductosEcom {
	return &ProductosEcom{NombrePEcom: nombre, PrecioPEcom: precio}
}

func agregarProducto(u *UsuarioEcom, pr *ProductosEcom, cant int) {
	pr.CantidadPEcom = cant
	u.ProductosUEcom = append(u.ProductosUEcom, pr)
	fmt.Println("Se ha agregado el producto ", pr, "al usuario ", u.NombreUsEcom)
	fmt.Println("Los datos completos son: ", u)
}

func borrarProducto(u *UsuarioEcom) {
	u.ProductosUEcom = []*ProductosEcom{}
	fmt.Println("Se han eliminado los productos del usuario correctamente.", u)
}
