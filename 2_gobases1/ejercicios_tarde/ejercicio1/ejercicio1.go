package main
import( 
	"fmt"
	"strings"
	)

func main(){
	fmt.Print("Ingrese una palabra: ")
	var s string
	fmt.Scanln(&s)
        fmt.Println("La palabra", strings.ToUpper(s), "tiene", len(s), "letras")

	for i := 0; i < len(s); i++ {
 		fmt.Printf("%c\n", s[i])
	}
}
