package main

import (
	"fmt"
)

// func main() {
// 	// multiplication
// 	var radius = 12.0

// 	area := math.Pi * radius * radius
// 	fmt.Println("area", area)

// 	// integer division
// 	half := 1 / 2
// 	fmt.Println("Half", half)

// 	halfFloat := 1.0 / 2.0
// 	fmt.Println("half float", halfFloat)

// 	// squaring, and raising to power
// 	badThreeSquared := 3 ^ 2 // this is the bitwise XOR operator
// 	fmt.Println("Three squared is not", badThreeSquared)
// 	goodThreeSquared := math.Pow(3.0, 2.0)
// 	fmt.Println("Three squared is", goodThreeSquared)

// 	// modulus operator
// 	remainder := 50 % 3
// 	fmt.Println("remaider is", remainder)

// 	// unary operators
// 	x := 3
// 	// can't do this
// 	//y := x++
// 	// can do this
// 	var y = x
// 	y++
// 	fmt.Println("x is", x, "and y is", y)
// }

// // 	// precedence

// // func main() {
// // 	// multiplication and division
// // 	a := 12.0 * 3.0 / 4.0
// // 	b := (12.0 * 3.0) / 4.0
// // 	c := 12.0 * (3.0 / 4.0)

// // 	fmt.Println("a", a, "b", b, "c", c)

// // 	// integer division
// // 	unclear := 12 * (3 / 4)
// // 	fmt.Println("unclear", unclear) //devuelve un 0 pq el rdo toma el primer numeor sin los decimales

// // 	// parenthesis
// // 	f := 12.0 / 3.0 / 4.0
// // 	fmt.Println("f", f) //1
// // 	f = 12.0 / (3.0 / 4.0)
// // 	fmt.Println("f", f) //16

// // 	// addition/subtraction
// // 	fmt.Println()
// // 	x := 12 + 3 - 4 // 11
// // 	y := (12 + 3) - 4 //11
// // 	z := 12 + (3 - 4) //11
// // 	fmt.Println("x", x, "y", y, "z", z)
// // 	x = 12 + 3*4 //24
// // 	y = (12 + 3) * 4 //60
// // 	z = 12 + (3 * 4) //24
// // 	fmt.Println("x", x, "y", y, "z", z)
// // }

// /**************************************/
// //Using the modulus Operator

// func main() {
// 	// does one number divide exactly into another?
// 	// x := 12
// 	// y := 5

// 	// if x%y == 0 {
// 	// 	fmt.Println(y, "divides exactly into", x)
// 	// } else {
// 	// 	fmt.Println(y, "does not divide exactly into", x)
// 	// }

// 	// thisMonth := 12
// 	// fmt.Println("The month after", thisMonth, "is", thisMonth + 1)

// 	for m := 1; m <= 12; m++ {
// 		fmt.Println("The month after", m, "is", m%12+1)
// 	}
// }

// func main() {
// 	second := 31
// 	minute := 1

// 	if minute < 59 && (second+1) > 59 {
// 		minute++
// 	}
// 	fmt.Println(minute)
// }

// func main() {
// 	a := 12
// 	b := 6
// 	//evitar dividir por 0
// 	// if b != 0 && divideTwoNumbers(a, b) == 2 {
// 	// 	fmt.Println("We found a 2")
// 	// }

// 	// otra forma de evitar que algo se divida por 0
// 	// 	if b != 0 {

// 	// 		c := divideTwoNumbers(a, b)

// 	// 		if c == 2 {

// 	// 			fmt.Println("We found a 2")
// 	// 		}
// 	// 	}

// 	/******************evitar divir en 0 en la func****/
// 	c, err := divideTwoNumbers(a, b)
// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		if c == 2 {
// 			fmt.Println("we foun two ")
// 		}
// 	}
// }

// func divideTwoNumbers(x, y int) (int, error) {
// 	if y == 0 {
// 		return 0, errors.New("cannot divide by 0")
// 	}
// 	return x / y, nil
// }

/****************************************/
func main() {
	x := 12
	x++

	fmt.Println(x)

	y := 10
	y /= 2 //5
	fmt.Println(y)
}
