package main

import (
	"fmt"
)

/**************FUNCTIONS************/
type Animal struct {
	Name         string
	Sound        string
	NumberOfLegs int
}

func main() {
	// myTotal := sumMany(2, 3, 4, 5)
	// fmt.Println(myTotal)

	var dog Animal
	dog.Name = "dog"
	dog.Sound = "woof"
	dog.Says()
}

//VARIADICT FUNCTION
//FUNC QUE SUM UNA CANT INDETERMINADA DE NUMEROS REICINIDOS POR ARG
//esto significa que puedo tener muchos num en el arg
// func sumMany(nums ...int) int {
// 	total := 0

// 	for _, x := range nums {
// 		total = total + x
// 	}

// 	return total
// }

//esta funcion va a tomar el valor del type Animal
//la funcion se llama Says()
func (a *Animal) Says() {
	fmt.Printf("A %s says %s", a.Name, a.Sound)
}

// func main() {
// 	z := addTwoNumbers(2, 4)
// 	fmt.Println(z)
// }

// func addTwoNumbers(x, y int) int {
// 	return x + y
// }

/**************MAPS*************/
//tienen keys y values [key]value
//reference type
// func main() {
// 	intMap := make(map[string]int)

// 	intMap["one"] = 1
// 	intMap["two"] = 2
// 	intMap["three"] = 3
// 	intMap["four"] = 4

// 	for key, value := range intMap {
// 		fmt.Println(key, value)
// 	} //a veces se muestra de forma random, no ordenada como lo declaramos

// 	//DELETE
// 	delete(intMap, "four")
// 	fmt.Println(intMap)

// 	//ENCONTRAR ALGO EN EL MAP
// 	//el es el key
// 	el, ok := intMap["one"]
// 	if ok {
// 		fmt.Println(el, "is in map") //devueleve el value del key
// 	} else {
// 		fmt.Println(el, "is not in the map") //devuelve el valor de el como 0 cuando no lo encuentra por default
// 	}

// }

/***********SLICES*******************/
// func main() {
// 	var animals []string
// 	animals = append(animals, "dog")
// 	animals = append(animals, "fish")
// 	animals = append(animals, "cat")
// 	animals = append(animals, "horse")

// 	fmt.Println(animals) //devuelve todo dentro del array

// 	//devuelve no por uno
// 	//i= index, para evitarlo se pone _
// 	//x cada elem del array
// 	for i, x := range animals {
// 		fmt.Println(i, x)
// 	}

// 	fmt.Println("Element 0 is", animals[0])
// 	fmt.Println("First two elements are", animals[0:2]) //muestra desde la posicion 0 dos elem, contando desde el 0

// 	fmt.Println("El largo del array es", len(animals))

// 	//SORT SLICE
// 	fmt.Println("Is ir sortted?", sort.StringsAreSorted(animals)) //Para saber si esta ordenado por letra

// 	sort.Strings(animals)
// 	fmt.Println("Is ir sortted?", sort.StringsAreSorted(animals)) //Para saber si esta ordenado por letra
// 	fmt.Println(animals)

// 	animals = DeleteFromSlice(animals, 1)
// 	fmt.Println(animals)

// }

// func DeleteFromSlice(a []string, i int) []string {
// 	a[i] = a[len(a)-1] //en la posicion dle index colocamos el ultimo elem
// 	a[len(a)-1] = ""   //borramos el ultimo elem
// 	a = a[:len(a)-1]   // el array ahora tiene todo excepto el ultimo elem
// 	return a
// }

/*******POINTERS**********/
//algo que apunta a un ligar en memoria especifico
// func main() {
// 	x := 10

// 	myFirstPointer := &x

// 	fmt.Println("x is", x)
// 	fmt.Println("myFirstPointer is", myFirstPointer) //devuelve el lugar en memoria es un conj de num y letras

// 	*myFirstPointer = 15 //le cambio el valor a la variable que esta en la direcc de myFirstPointer
// 	fmt.Println("x is now", x)

// 	changeValueOfPointer(&x)
// 	fmt.Println("x is now", x)

// }

// func changeValueOfPointer(num *int) {
// 	*num = 25
// }

/***************STRUCT*************/
// type Car struct {
// 	NumberOfTires int
// 	Luxury        bool
// 	BucketsSeats  bool
// 	Make          string
// 	Model         string
// 	Year          int
// }

// func main() {

// var myCar Car
// myCar.BucketsSeats = true
// myCar.Year = 2021
// myCar.Luxury = false

// 	myCar := Car{
// 		NumberOfTires: 4,
// 		Luxury:        true,
// 		BucketsSeats:  true,
// 		Make:          "Volvo",
// 		Model:         "XC()",
// 		Year:          2019,
// 	}
// 	fmt.Printf("My car is a %d %s %s", myCar.Year, myCar.Make, myCar.Model)
// }

/*************ARRAY*********/
// func main() {

// 	var myStrings [3]string //array con 3 elem de valor string
// 	myStrings[0] = "cat"
// 	myStrings[1] = "dog"
// 	myStrings[2] = "fish"

// 	fmt.Println("First element ", myStrings[0])

// }

/*************VARIABLES************/
// import "log"

// //basic types (numbers, string, booleans)
// var myInt int
// var myInt16 int16 //dependiendo si el programa procesa 16 32 o 64 bits

// var myUint uint //solo pueden ser valores positivos o 0

// var myFloat float32
// var myFloat64 float64 //puede ser un num mucho + largo

// func main() {
// 	myInt = 10
// 	myUint = 8
// 	myFloat = 12.45545
// 	myFloat64 = 100.245678

// 	log.Println(myFloat, myFloat64, myInt, myUint)

// 	myString := "S"
// 	log.Println(myString)

// 	myString = "John" //los strign son inmutables, x lo que hace es
// 	// crear un nuevo string y lo coloca en la var

// 	log.Println(myString)

// 	var myBool = true
// 	myBool = false
// 	log.Println("myBool: ", myBool)

// }
