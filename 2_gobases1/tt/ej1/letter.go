package ej1

func LetterOfWord(word string) (int, []string){
	lenWord := len(word)
	//letter := strings.Split(word, "")
	letter := make([]string, lenWord)
	for i := 0; i < lenWord; i++ {
		letter[i] = string(word[i])
	}
	return lenWord, letter
}