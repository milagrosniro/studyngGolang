package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/eiannone/keyboard"
)

var reader *bufio.Reader

/***EXPERIMENTING WITH INTERPOLATION STRINGS***/

//crear un type de var
type User struct {
	UserName        string
	Age             int
	FavouriteNumber float64
	OwnsADog        bool
}

func main() {
	reader = bufio.NewReader(os.Stdin)

	var user User //user de tipo User

	user.UserName = readString("what is your name?")

	user.Age = readInt("How old are you?")
	user.FavouriteNumber = readFloat("What is your favourite number")
	user.OwnsADog = readBool("Do you want a dog?(y/n)")

	fmt.Printf("Your name is %s. You are %d years old. Your favourite number is %.2f, and is %t that you want a dog", user.UserName, user.Age, user.FavouriteNumber, user.OwnsADog)

}

func prompt() {
	fmt.Println("-->")
}

func readString(s string) string {
	for {

		fmt.Println(s)
		prompt()

		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1)
		userInput = strings.Replace(userInput, "\n", "", -1)
		if userInput == "" {
			fmt.Println("Please enter a value")
		} else {

			return userInput
		}
	}

}

func readInt(s string) int {
	for {

		fmt.Println(s) //imprimo la ??
		prompt()

		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1)
		userInput = strings.Replace(userInput, "\n", "", -1)

		num, err := strconv.Atoi(userInput)

		if err != nil {
			fmt.Println("Plase enter a whole number")
		} else {

			return num //el loop termina cuando tengamos esta rta
		}
	}
}

func readFloat(s string) float64 {
	for {

		fmt.Println(s) //imprimo la ??
		prompt()

		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1)
		userInput = strings.Replace(userInput, "\n", "", -1)

		num, err := strconv.ParseFloat(userInput, 64) //convierte un numero float en string

		if err != nil {
			fmt.Println("Plase enter a  y or n")
		} else {

			return num //el loop termina cuando tengamos esta rta
		}
	}
}

func readBool(s string) bool {
	err := keyboard.Open()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		_ = keyboard.Close()
	}()

	for {
		fmt.Println(s)
		char, _, err := keyboard.GetSingleKey()
		if err != nil {
			log.Fatal(err)
		}
		if char == 'n' || char == 'N' {
			return false
		}
		if char == 'y' || char == 'Y' {
			return true
		} else {
			fmt.Println("Plase enter a  number")
		}

	}

}

/**********************interpolation strings***/
// func main() {
// 	reader = bufio.NewReader(os.Stdin)
// 	userName := readString("what is your name?")

// 	age := readInt("How old are you?")

// 	//Forma no eficiente de concatenar
// 	// fmt.Println("your name is "+ userName + ". You are", age, "years old")

// 	//Otra forma
// 	// fmt.Println(fmt.Sprintf("Your name is %s. You are %d years old", userName, age))

// 	//la forma + eficiente, usa menos memoria y es + rapido.
// 	fmt.Printf("Your name is %s. You are %d years old", userName, age)

// }

// func prompt() {
// 	fmt.Println("-->")
// }

// func readString(s string) string {
// 	for {

// 		fmt.Println(s)
// 		prompt()

// 		userInput, _ := reader.ReadString('\n')
// 		userInput = strings.Replace(userInput, "\r\n", "", -1)
// 		userInput = strings.Replace(userInput, "\n", "", -1)
// 		if userInput == "" {
// 			fmt.Println("Please enter a value")
// 		} else {

// 			return userInput
// 		}
// 	}

// }

// func readInt(s string) int {
// 	for {

// 		fmt.Println(s) //imprimo la ??
// 		prompt()

// 		userInput, _ := reader.ReadString('\n')
// 		userInput = strings.Replace(userInput, "\r\n", "", -1)
// 		userInput = strings.Replace(userInput, "\n", "", -1)

// 		num, err := strconv.Atoi(userInput)

// 		if err != nil {
// 			fmt.Println("Plase enter a whole number")
// 		} else {

// 			return num //el loop termina cuando tengamos esta rta
// 		}
// 	}
// }
