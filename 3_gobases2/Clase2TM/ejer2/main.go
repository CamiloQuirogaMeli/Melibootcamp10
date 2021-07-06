package main

import "fmt"

func calculateAvarage(scores ...int) float64 {

	var scoreAvarage float64

	for _, score := range scores {
		scoreAvarage += float64(score)
	}

	return scoreAvarage / float64(len(scores))
}

func main() {

	studentNotes := calculateAvarage(6, 6, 8)

	fmt.Println(studentNotes)
}
