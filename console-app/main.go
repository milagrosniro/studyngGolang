package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/eiannone/keyboard"
)

/******FUN QUE MUESTRA NOMBRE DE LETRA***/
func main() {
	err := keyboard.Open()

	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		_ = keyboard.Close()
	}()

	coffees := make(map[int]string) //creo un array cuyo subindice es de type in y su valor string
	coffees[1] = "capuccino"
	coffees[2] = "Late"
	coffees[3] = "Americano"
	coffees[4] = "Mocha"
	coffees[5] = "Macchiato"
	coffees[6] = "Espresso"

	fmt.Println("MENU")
	fmt.Println("-----")
	fmt.Println("1- Capuccino")
	fmt.Println("2- Late")
	fmt.Println("3- Americano")
	fmt.Println("4- Mocha")
	fmt.Println("5- Macchiato")
	fmt.Println("6- Espresso")
	fmt.Println("Q- Quit the program")

	//challengue
	/*typing either an uppercase q or a lowercase q*/
	char := ' '
	for char != 'q' && char != 'Q' {
		char, _, err := keyboard.GetSingleKey()

		if err != nil {
			log.Fatal(err)
		}

		i, _ := strconv.Atoi(string(char))
		_, ok := coffees[i] //busco el valor en el map
		if ok {

			fmt.Println(fmt.Sprintf("you chose %s", coffees[i]))
		}


//OTRO EJ DE DO WHILE LOOP
	// 	for ok := true; ok; ok = char != 'q' && char != 'Q'{
	// 		char, _, err := keyboard.GetSingleKey()

	// 		if err != nil {
	// 			log.Fatal(err)
	// 		}
	
	// 		i, _ := strconv.Atoi(string(char))
	// 		_, ok := coffees[i] //busco el valor en el map
	// 		if ok {
	
	// 			fmt.Println(fmt.Sprintf("you chose %s", coffees[i]))
	// 		}
	// 	}

	// }

	fmt.Println("program exiting!!!")

	/**********************EJEMPLO DE DO WHILE LOOP***********************/
	// for {
	// 	char, _, err := keyboard.GetSingleKey()

	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	//se pone antes de la seccion de imprimir para que no imprima you chose
	// 	if char == 'q' || char == 'Q' {
	// 		break
	// 	}
	// 	//declaro la var i, ignoro el err y llamo a la libreria strconv.Atoi() q transforma un str en un num
	// 	i, _ := strconv.Atoi(string(char))
	// 	t := fmt.Sprintf("you chose %s", coffees[i]) //el %q transforma el valor en string
	// 	fmt.Println(t)
	// 	//rune is a single character, the lower level than type string

	// }

	// fmt.Println("program exiting!!!")
}

/**********FUNC Q MUESTRA NUM TECLA*/
// func main() {
// 	err := keyboard.Open() //devuelve un error, si no hay error su valor es nil (q es null/nada)
// 	//abrimos el package keyboard

// 	//if we cant open de keyboard package
// 	if err != nil {
// 		log.Fatal(err) //el programa deja de andar
// 	}

// 	//defer no se ejecuta inmediatamente
// 	//se ejecuta apenas la funcion actual dnd esta la palabra defer (en este caso main) apenas termina
// 	//toma lugar defer
// 	//apenas main() termine que se ejecute lo q esta en keyboard.Close
// 	defer func() {
// 		_ = keyboard.Close()
// 	}() //funcion anomica se tiene q autoinvocar al terminar

// 	fmt.Println("press any key on the keyboard. Press ESC to quit")

// 	for {
// 		char, key, err := keyboard.GetSingleKey()

// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 		if key != 0 {
// 			fmt.Println("you pressed", char, key)
// 		} else {
// 			fmt.Println("you pressed", char) //devuelve el nro de character
// 		}

// 		if key == keyboard.KeyEsc {
// 			break
// 		}
// 	}

// 	fmt.Println("program exiting!!!")
// }

/*****************************************/
//REVISAR PORQUE NO FUNCIONO!!!!! video 21
// func main() {
// 	reader := bufio.NewReader(os.Stdin)

// 	for {
// 		fmt.Print("->")

// 		userInput, _ := reader.ReadString('\n') //el _ sirve para obviar la variable de error que devuelve

// 		userInput = strings.Replace(userInput, "\n", "", -1) //reemplazo el enter por nada

// 		if userInput == "quit" {
// 			break
// 		} else {
// 			fmt.Println(userInput)
// 		}
// 	}
// }
