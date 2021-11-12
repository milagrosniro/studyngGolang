package main

//comandos
//go build -o eliza.exe main.go
//lugo ./eliza
import (
	"bufio"
	"fmt" //fmt= format package
	"myapp/doctor"
	"os"
	"strings"
)

//Eliza
func main() {
	reader := bufio.NewReader(os.Stdin) //le decimos que queremos leer

	// var whatToSay string
	whatToSay := doctor.Intro()
	fmt.Println(whatToSay)

	//El unico loop
	for {
		fmt.Print("-> ")
		userInput, _ := reader.ReadString('\n') //lee cualq type, pero hay q decirle cuando deja de typing, en nuestro caso es cuando apretamos enter
		//userInput va a imprimir lo que nosotros escrubamos

		//aca cortamos el string de la rta, para que no tenga en cuenta el enter y solo lo escrito
		//(palabra q queremos reemplazar, parte que le queremos cortar, cuantas veces queremos hacer)
		//aca ponemos -1 que es cada vez q lo encuentre
		userInput = strings.Replace(userInput, "\r\n", "", -1) //windows
		userInput = strings.Replace(userInput, "\n", "", -1)   //linux,mac

		if userInput == "quit" {
			break //para salir del loop
		} else {
			fmt.Println(doctor.Response(userInput))
		}

	}

}

//func main() {
// fmt.Println("Hello World!") //imprime en lineas separadas
// fmt.Println("otra linea")
// fmt.Print("this is come text") //imprime todo en una misma linea
// fmt.Print("this is some text")

//DECLARAR Y DEFINIR UNA VARIABLE
// var whatToSay string = "hola Mundo"
// whatToSay = "Hello World!" //cambia el valor de la var

//Otra forma resumida de definir una var
//aca Go reconoce solo el tipo de datp
// 	whatToSay := "hello world"

// 	sayHelloWorld(whatToSay)
// }

// func sayHelloWorld(whatToSay string) {
// 	fmt.Println(whatToSay)
// }
