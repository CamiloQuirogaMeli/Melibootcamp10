package ej4



func NumToMonth(month uint8) string {
	months := map[uint8]string{
		1 : "january",
		2 : "february",
		3 : "march",
		4 : "april",
		5 : "may",
		6 : "june",
		7 : "july",
		8 : "august",
		9 : "september",
		10 : "october",
		11 : "november",
		12 : "december",
	}
	if month <= 12 && month > 0{
		return months[month]
	}else {
		return "Wrong Month"
	}
}
