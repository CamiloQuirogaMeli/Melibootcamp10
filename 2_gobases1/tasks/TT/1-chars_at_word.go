package main

func amountOfChars(word string) {
	println(len(word))
	for _, char := range word {
		println(string(char))
	}
}
