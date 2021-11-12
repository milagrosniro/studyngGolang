package main

import "fmt"

type Employee struct {
	Name     string
	Age      int
	Salary   int
	Fulltime bool
}

func main() {
	jack := Employee{
		Name:     "Kack",
		Age:      27,
		Salary:   40000,
		Fulltime: false,
	}

	jill := Employee{
		Name:     "Jill",
		Age:      33,
		Salary:   60000,
		Fulltime: false,
	}

	var employees []Employee
	employees = append(employees, jack)
	employees = append(employees, jill)

	for _, x := range employees {
		if x.Age > 30 {
			fmt.Println(x.Name, "is older")
		} else {
			fmt.Println(x.Name, "is young")
		}

		if x.Age > 30 && x.Salary > 50000 {
			fmt.Println(x.Name, "makes more money")
		} else {
			fmt.Println(x.Name, "makes less money")

		}
	}
}

/*
// func main() {
// 	apples := 18
// 	oranges := 9

// 	//boolean expression
// 	fmt.Println(apples == oranges)
// 	fmt.Println(apples != oranges)

// 	//hay que comparar los mismos type de datos

// 	fmt.Println(apples > oranges)
// 	fmt.Println(apples <= oranges)

// }
*/
