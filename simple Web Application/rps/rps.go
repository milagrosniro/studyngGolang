package rps

import (
	"math/rand"
	"time"
)

const (
	ROCK         = 0 // beats scissors. (scissors + 1) % 3 = 0
	PAPER        = 1 // beats rock. (rock + 1) % 3 = 1
	SCISSORS     = 2 // beats paper. (paper + 1) % 3 = 2
	PLAYERWINS   = 1
	COMPUTERWINS = 2
	DRAW         = 3
)

type Round struct {
	//Winner         int    `json:"winner"`
	ComputerChoice string `json:"computer_choice"`
	RoundResult    string `json:"round_result"`
	Message        string `json:"round_message"`
}

var winMessages = []string{
	"Good Job!",
	"Nice work!",
	"You should buy a lottery ticket!",
}

var loseMessages = []string{
	"Too bad!",
	"Try Again!",
	"This is just not your day",
}

var drawMessages = []string{
	"Great minds think alike!",
	"Uh, oh. Try again!",
	"YNobody wins, but you can try again",
}

func PlayRound(playerValue int) Round {
	rand.Seed(time.Now().UnixNano()) //numero random
	computerValue := rand.Intn(3)    //elige un nro random del entre 0<=n<3
	computerChoice := ""
	roundResult := ""
	//winner := 0

	switch computerValue {
	case ROCK:
		computerChoice = "Computer chose ROCK"
		break
	case PAPER:
		computerChoice = "Computer chose PAPER"
		break
	case SCISSORS:
		computerChoice = "Computer chose SCISSORS"
		break
	default:
	}
	messageInt := rand.Intn(3)
	message := ""

	if playerValue == computerValue {
		roundResult = "It's a draw"
		//winner = DRAW
		message = drawMessages[messageInt]
	} else if playerValue == (computerValue+1)%3 {
		roundResult = "Player wins!"
		//winner = PLAYERWINS
		message = winMessages[messageInt]
	} else {
		roundResult = "Computer wins!"
		//winner = COMPUTERWINS
		message = loseMessages[messageInt]
	}

	var result Round
	//result.Winner = winner
	result.ComputerChoice = computerChoice
	result.RoundResult = roundResult
	result.Message = message

	return result
}
