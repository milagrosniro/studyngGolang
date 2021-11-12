package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

//maneras de declarar variables
// func main(){

// var firstNumber int
// firstNumber = 2

// var secondNumber = 5

// subtraccion := 7
// }

const prompt = "and don't type your number in, just press ENTER when ready"

func main() {
	//seed de random number generator
	rand.Seed(time.Now().UnixNano())

	var firstNumber = rand.Intn(8) + 2 //devuelve un num random del 2-10
	var secondNumber = rand.Intn(8) + 2
	var subtraccion = rand.Intn(8) + 2
	var answer = firstNumber*secondNumber - subtraccion

	game(firstNumber, secondNumber, subtraccion, answer)
}

func game(firstNumber, secondNumber, subtraction, answer int) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Guess the Number Game")
	fmt.Println("---------------------")
	fmt.Println("")

	fmt.Println("Think of a number between 1-10", prompt)
	reader.ReadString('\n')

	//take them through the games

	fmt.Println("Multiply your number by", firstNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Now multiply the result by", secondNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Divide the result by the number you originally thought of", prompt)
	reader.ReadString('\n')

	fmt.Println("Now subtract", subtraction, prompt)
	reader.ReadString('\n')

	//give them the  answer

	fmt.Println("The answer is", answer)
}
