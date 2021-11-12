package main

import (
	// "fmt"
	"myapp/packageone"
)

var myVar = "myVar"

func main() {
	//variables
	//main package
	//declare a package level var for the main
	//package named myVar

	//declare a block level var for the main func
	//called blockVar
	var blockVar = "block var"

	//package main
	//in the main func print out the values of myVar,
	//blockVar and packageVar on one line using the PrintMe
	//func in packageone

	packageone.PrintMe(myVar, blockVar)
}

// var one = "One"

// func main() {
// 	var one = "two"
// 	fmt.Println(one)
// 	myFunc()

// 	newString := packageone.PublicVar
// 	fmt.Println("from package", newString)

// 	packageone.Export()
// }

// func myFunc() {
// 	// var one = "the number one"
// 	fmt.Println(one)
// }
