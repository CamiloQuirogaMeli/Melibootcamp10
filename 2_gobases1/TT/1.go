package main
import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	letrasDeUnaPalabra()
}

func letrasDeUnaPalabra() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Letras de una palabra")
	fmt.Println("---------------------")
  
	for {
	  fmt.Print("-> ")
	  text, _ := reader.ReadString('\n')
	  // convert CRLF to LF (Only Windows)
	  text = strings.Replace(text, "\n", "", -1)
	  
	  fmt.Printf("La palabra %s tiene %d caracteres, y se deletrea de la siguiente manera:\n", text, len(text))
	  for i := 0; i < len(text); i++ {
		  fmt.Printf("%c\n", text[i])
	  }
	}
}