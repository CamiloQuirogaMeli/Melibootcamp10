package tt

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type CleaningProducts struct{
	Id       uint    `json:"id"`
	Price    float64 `json:"price"`
	Quantity uint    `json:"stock"`
}

func (p *CleaningProducts) SaveInfo() {
	productString := fmt.Sprintln(p.Id,";",p.Price,";",p.Quantity,";")
	data, _ := ioutil.ReadFile("./data/products")
	data = append(data,productString...)
	_ = ioutil.WriteFile("./data/products", data,0644)
}

func PrintInvoice() {
	data, err := ioutil.ReadFile("./data/products")
	if err != nil {
		fmt.Println("File not exists")
	}
	dataClean := strings.Replace(string(data)," ","",-1)
	dataClean = strings.Replace(dataClean,"\n"," ",-1)
	str := strings.Split(dataClean,";")
	titleToPrint := fmt.Sprint("ID\t\tPrecio\tCantidad\t")
	fmt.Println(titleToPrint)
	for i := 0; i < len(str)-3; i=i+3 {
		firstSpace := 22 - len(str[i])
		line := fmt.Sprintf("%s%*s%10s\n", str[i], firstSpace, str[i+1], str[i+2])
		fmt.Print(line)
	}
}
