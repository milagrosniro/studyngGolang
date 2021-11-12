package main

//DO WHILE LOOP
func main() {
	for {
		//do some work

		//check a condition and break
		//VER EN CONSOLE-APP

	}
}

//NESTED LOOP, es un loop sin un loop
// func main() {
// 	//cada vez que el primer loop de una vuelta el segundo loop va a dar 3 vueltas
// 	for i := 1; i <= 10; i++ {
// 		fmt.Print("i is ", i)
// 		for j := 1; j <= 3; j++ {
// 			fmt.Print("   j: ", j)
// 		}
// 		fmt.Println()
// 	}
// }

/******************************/
//INFINTE LOOP, se puede usar con go routines
// func main() {
// 	//read input from the user 5 times and write it to a log

// 	reader := bufio.NewReader(os.Stdin)
// 	ch := make(chan string)

// 	go mylogger.ListenForLog(ch)

// 	fmt.Println("Enter something")

// 	for i := 0; i < 5; i++ {
// 		fmt.Print("->")
// 		input, _ := reader.ReadString('\n')
// 		ch <- input
// 		time.Sleep(time.Second)
// 	}
// }

//PATH LOOP
// func main() {
// 	//parth loop
// 	for i := 0; i <= 10; i++ {
// 		fmt.Println("i=", i)
// 	}
// }

//while looop

// func main() {
// 	rand.Seed(time.Now().UnixNano())
// 	i := 1000
// 	//execute a loop while i> 100
// 	for i > 100 {
// 		//get a random number between 1 and 1001
// 		i = rand.Intn(1000) + 1
// 		fmt.Println("i =", i)

// 		if i > 100 {
// 			fmt.Println("is is", i, "so loop keeps going")
// 		}
// 	}
// 	fmt.Println("Got", i, "and broke out of loop")
// }
