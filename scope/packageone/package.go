package packageone

import "fmt"

//package one
//declare a packege level var in the packageone
//package named PackageVar

//create an exported func in packageone called PrintMe
var PackageVar = "package var"

//NO OLVIDAR de definir el tipo de var que van a ser los arg
//como arg coloco solo los que voy a reicbir en el otro package
func PrintMe(myVar, blockVar string) {
	fmt.Println(myVar, blockVar, PackageVar)
}

// import "fmt"

// var privateVar = "iAm private"

// var PublicVar = "iam public" //las var que se exportan empiezan en mayuscula

// func notExported() {
// 	fmt.Println("privada")
// }

// func Export() {
// 	fmt.Println(privateVar)
// }
