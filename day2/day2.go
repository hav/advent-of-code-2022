package day2

import (
	"strings"

	utils "besharati.se/advent-of-code-2022/utils"
)

type Move int

const (
	Rock     Move = 1
	Paper    Move = 2
	Scissors Move = 3
)

func (us Move) result(them Move) int {
	switch us {
	case Rock:
		if them == Scissors {
			return 6
		} else if them == Paper {
			return 0
		}
		return 3

	case Paper:
		if them == Rock {
			return 6
		} else if them == Scissors {
			return 0
		}
		return 3

	case Scissors:
		if them == Paper {
			return 6
		} else if them == Rock {
			return 0
		}
		return 3
	default:
		return 0
	}
}

var inputToMove = map[string]Move{
	"A": Rock, "B": Paper, "C": Scissors,
	"X": Rock, "Y": Paper, "Z": Scissors,
}

func chooseMove(them Move, desiredResult string) Move {
	switch desiredResult {
	case "X":
		switch them {
		case Rock:
			return Scissors
		case Paper:
			return Rock
		case Scissors:
			return Paper
		default:
			break
		}
	case "Y":
		return them
	case "Z":
		switch them {
		case Rock:
			return Paper
		case Paper:
			return Scissors
		case Scissors:
			return Rock
		default:
			break
		}
	default:
		return 0
	}
	return 0
}

func rockPaperScissors(input []string, choose bool) int {
	res := []int{}
	for _, v := range input {
		c := strings.Split(v, " ")
		var them, us Move

		if choose {
			them = inputToMove[strings.ToUpper(c[0])]
			us = chooseMove(them, c[1])
		} else {
			them, us = inputToMove[strings.ToUpper(c[0])], inputToMove[strings.ToUpper(c[1])]
		}

		res = append(res, us.result(them)+int(us))
	}
	return utils.SumIntSlice(res)
}
