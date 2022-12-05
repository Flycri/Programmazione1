package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getScore(your_play string) int {
	switch your_play {
	case "X":
		return 1
	case "Y":
		return 2
	case "Z":
		return 3
	default:
		os.Exit(1)
	}
	return -1
}

func runMatch(opp_play, your_play string) int {
	if (opp_play == "A") && (your_play == "X") || (opp_play == "B") && (your_play == "Y") || (opp_play == "C") && (your_play == "Z") {
		return 3
	}
	if (opp_play == "A") && (your_play == "Y") || (opp_play == "B") && (your_play == "Z") || (opp_play == "C") && (your_play == "X") {
		return 6
	}
	return 0
}

func getPlay(opp_play, result string) string {
	switch opp_play {
	case "A":
		switch result {
		case "X":
			return "Z"
		case "Y":
			return "X"
		case "Z":
			return "Y"
		}
	case "B":
		switch result {
		case "X":
			return "X"
		case "Y":
			return "Y"
		case "Z":
			return "Z"
		}
	case "C":
		switch result {
		case "X":
			return "Y"
		case "Y":
			return "Z"
		case "Z":
			return "X"
		}
	}
	return ""
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	var tot_score int

	for scanner.Scan() {
		match_slice := strings.Split(scanner.Text(), " ")
		tot_score += getScore(getPlay(match_slice[0], match_slice[1]))
		tot_score += runMatch(match_slice[0], getPlay(match_slice[0], match_slice[1]))
	}

	fmt.Println(tot_score)

}
