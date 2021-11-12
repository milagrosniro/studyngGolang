package main

import "fmt"

func main() {
	age := 10
	name := "Jack"
	rigthHanded := true

	fmt.Printf("%s is %d years old.Right handed: %t", name, age, rigthHanded)
	fmt.Println()

	ageInTenYears := age + 10 //esto es una expresion pq puede ser evaluado como un unico valor

	fmt.Printf("In ten years, %s will be %d years old", name, ageInTenYears)
	fmt.Println()

	isATeenager := age >= 13

	fmt.Println(name, "is a teenager:", isATeenager)

}
