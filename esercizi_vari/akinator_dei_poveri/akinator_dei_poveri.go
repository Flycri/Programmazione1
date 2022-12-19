package main

import (
	"fmt"
	"reflect"
)

type Question struct {
	id       int
	question string
	id_left  int
	id_right int
}

func FindAnimal(risposte []bool, animals map[string][]bool) string {
	for k, v := range animals {
		if reflect.DeepEqual(risposte, v) {
			return k
		}
	}
	return "-1"
}

func main() {
	animals := make(map[string][]bool)
	var question []Question
	var input string

	animals = readJsonFromFile("animals.json")
	question = readJsonFromFile("domande.json")
	risposte := make([]bool, len(question))

	actual_question := question[0]

	for {
		fmt.Printf("%s t/f", actual_question.question)
		fmt.Scan(&input)
		if input == "t" {
			risposte[actual_question.id] = true
			if actual_question.id_left == -1 {
				break
			}
			actual_question.id = actual_question.id_left
		} else if input == "f" {
			risposte[actual_question.id] = false
			if actual_question.id_right == -1 {
				break
			}
			actual_question.id = actual_question.id_right
		} else {
			continue
		}
	}

	answer_animal := FindAnimal(risposte, animals)

	fmt.Println("Stavi pensando all'animale ", answer_animal, "? t/f")
	fmt.Scan(&input)
	if input == "t" {
		fmt.Println("ho indovinato")
	} else {
		fmt.Println("mi dispiace")
	}
}
