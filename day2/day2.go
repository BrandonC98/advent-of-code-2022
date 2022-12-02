package main

import (
	"io/ioutil"
	"log"
	"strings"
)

const(
	win = 6
	draw = 3
	lost = 0

	rock = 1
	paper = 2
	scissors = 3

	oRock = "A"
	oPaper = "B"
	oScissors = "C"

	pRock = "X"
	pPaper = "Y"
	pSissors = "Z"

)

func battle(player string, opp string) int {
	score := 0

	p, o := getChoice(player, opp)

	switch  {
	case p == o: score += draw
	case p == rock && o == scissors: score += win
	case p == paper && o == rock: score += win
	case p == scissors && o == paper: score += win
	default: score += lost
	}

	return score + p
}

func getChoice(player string, opp string) (int, int) {
	var playerChoice int
	var oppsChoice int
	switch player {
	case "X": playerChoice = rock
	case "Y": playerChoice = paper
	case "Z": playerChoice = scissors
	}
	
	switch opp {
	case "A": oppsChoice = rock
	case "B": oppsChoice = paper
	case "C": oppsChoice = scissors
	}

	return playerChoice, oppsChoice
}

func part2() {
	path := "inputs/1.txt"
	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	s := string(file)

	ss := strings.Split(s, "\n")

	totalScore := 0

	//for each round
	for n := range ss {
		round := ss[n]
		opp := string(round[0])
		player := string(round[2])

		// println("round 		", round)
		// println("opp 		", opp)
		// println("player 		", player)

		totalScore += battle(player, opp)

	}

	println("total score ", totalScore)

}

func part1() {
	path := "inputs/2.txt"
	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	s := string(file)

	ss := strings.Split(s, "\n")

	totalScore := 0

	//for each round
	for n := range ss {
		round := ss[n]
		opp := string(round[0])
		player := string(round[2])

		// println("round 		", round)
		// println("opp 		", opp)
		// println("player 		", player)

		totalScore += battle(player, opp)

	}

	println("total score ", totalScore)

}

func main() {
	println("Start")
	part2()
}