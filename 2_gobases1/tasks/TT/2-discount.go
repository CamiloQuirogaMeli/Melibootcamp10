package main

func discount(price int, discount float64) (result float64) {
	percent := float64(price) * discount

	result = float64(price) - percent

	return result
}
