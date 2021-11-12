package main

import (
	"fmt"

	"github.com/eiannone/keyboard"
)

//creo un channel que va a ser una rune (un solo string)
var keyPressChan chan rune

func main() {
	//go dispara esta funcion y sigue adelante

	//primero se ejecuta el print, pq la func q esta dsp de go se imprime cuando termina

	// go doSomething("Hello world")

	// fmt.Println("this is another msg")
	// for {
	// 	//do something
	// }

	//DEFINICION DE UN CHANEL
	//para hacer un channel funcional es necesario usar la palabra make
	keyPressChan := make(chan rune)
	go listenForKeyPress()
	fmt.Println("Press any key, or q to quit")
	_ = keyboard.Open() //abrimos la keyboard

	defer func() {
		keyboard.Close()
	}()

	for {
		char, _, _ := keyboard.GetSingleKey() //escuchamos la keyborad y la guradamos en l avar char
		if char == 'q' || char == 'Q' {
			break
		}

		keyPressChan <- char //le enviamos la info al channel, char va a ser el rune que ocupe el argumento?
	}
}

//func doSomething(s string) {
// 	until := 0
// 	for {
// 		fmt.Println("s is", s)
// 		until = until + 1
// 		if until >= 5 {
// 			break
// 		}
// 	}
//}

func listenForKeyPress() {
	for {
		key := <-keyPressChan //esto frena cuando no reciba un valor, ley tiene el valor de lo que reciba el channel keyPressChan
		fmt.Println("You pressed", string(key))
	}
}

//Channels permiten enviar info de un lugar a otro. Son usados en GoRoutines
