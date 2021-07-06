package main

import (
	"fmt"
)

func (u *Usuario) setProducto(prod *Producto) {
	u.productos = append(u.productos, *prod)
}

func nuevoProducto(nombre string, precio float64) (*Producto, error) {
	if precio <= 0 {
		return nil, fmt.Errorf("el precio ingresado %f.2 no puede ser menor a 1", precio)
	}
	ptrPod := &Producto{
		nombre: nombre,
		precio: precio,
	}
	return ptrPod, nil
}

func agregarProducto(user *Usuario, product *Producto, cantidad int) error {
	if cantidad <= 0 {
		return fmt.Errorf("la cantidad no puede ser menos que 1 unidad")
	}
	product.cantidad = cantidad
	user.setProducto(product)
	return nil
}

func borrarProductos(user *Usuario) error {
	if user.productos == nil {
		return fmt.Errorf("el usuario ingresado no tiene productos para eliminar")
	}
	user.productos = nil
	return nil
}

func main() {
	user := &Usuario{
		nombre:   "Franco",
		apellido: "Ballet",
		correo:   "franco.ballet@mercadolibre.com",
	}
	var (
		nombre   string
		precio   float64
		cantidad int
	)
	fmt.Println("Ingrese el nombre del nuevo producto")
	fmt.Scanln(&nombre)
	fmt.Println("Ingrese el precio del mismo")
	fmt.Scanln(&precio)
	produ, err := nuevoProducto(nombre, precio)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Ingrese la cantidad a agregar")
		fmt.Scanln(&cantidad)
		er := agregarProducto(user, produ, cantidad)
		if er != nil {
			fmt.Println(er)
		} else {
			fmt.Println("Imprimiendo el usuario")
			fmt.Println(user)
			fmt.Println("Se borraran sus productos")
			err := borrarProductos(user)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Productos eliminados, imprimiendo usuario")
				fmt.Println(user)
			}
		}
	}

}
