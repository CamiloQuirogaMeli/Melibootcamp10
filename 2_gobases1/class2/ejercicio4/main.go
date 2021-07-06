package main

import "fmt"

var months = map[int]string{
	1:  "January",
	2:  "February",
	3:  "March",
	4:  "April",
	5:  "May",
	6:  "June",
	7:  "July",
	8:  "August",
	9:  "September",
	10: "October",
	11: "November",
	12: "December",
}

func main() {

	var m int

	fmt.Scanf("%d", &m)

	whichMonth(m)
}

func whichMonth(monthNumber int) {

	if monthNumber > 12 || monthNumber < 1 {
		fmt.Println("The input month number must be between 1 and 12")
	} else {
		fmt.Println(months[monthNumber])
	}

}
