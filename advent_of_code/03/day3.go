package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

type Prio struct {
	value       int
	occurrences int
}

func findSameRune(s1, s2 string) rune {
	var r rune
	for _, r = range s1 {
		if strings.ContainsRune(s2, r) {
			return r
		}
	}
	return -1
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	priority_map := make(map[rune]Prio)
	var tot int

	for scanner.Scan() {
		comp1 := scanner.Text()[:len(scanner.Text())/2]
		comp2 := scanner.Text()[len(scanner.Text())/2:]
		same_rune := findSameRune(comp1, comp2)
		if unicode.IsUpper(same_rune) {
			new_element := priority_map[same_rune]
			new_element.occurrences++
			new_element.value = int(byte(same_rune) - 38)
			priority_map[same_rune] = new_element
		} else {
			new_element := priority_map[same_rune]
			new_element.occurrences++
			new_element.value = int(byte(same_rune) - 96)
			priority_map[same_rune] = new_element
		}
	}

	for _, v := range priority_map {
		tot += v.value * v.occurrences
	}

	fmt.Println(tot)

}
