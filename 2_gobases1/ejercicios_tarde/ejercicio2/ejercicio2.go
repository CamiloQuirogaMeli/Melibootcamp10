package main

import(
        "fmt"
        )

func main(){
        fmt.Print("Ingrese el precio de la prenda: ")
        var p float32
        fmt.Scanln(&p)
	if p < 0{
		fmt.Println("El precio no puede ser negativo")
	}else{
		fmt.Print("Ingrese el descuento que quiere aplicar: ")
        	var d int16
        	fmt.Scanln(&d)
		if d < 0{		
			fmt.Println("El descuento no puede ser negativo")
		}else{
		
			var descuento float32 = 1 - float32(d)/100.0
	
        		fmt.Printf("El precio con descuento es: %.2f\n", p*descuento)
		}
	}
	
}

