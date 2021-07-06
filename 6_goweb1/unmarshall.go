package main
​
import (
	"encoding/json"
	"fmt"
	"log"
)
​
type product struct {
	Name      string
	Price     int
	Published bool
}
​
func main(){
	jsonData := `{"Name": "MacBook Air", "Price": 900, "Published": true}`//se necesita
	//la comilla invertida para escapar las dobles comillas del medio
​
	p := product{}
​
	if err := json.Unmarshal([]byte(jsonData), &p); err != nil {
		log.Fatal(err)
	}
​
	fmt.Println(p)
}