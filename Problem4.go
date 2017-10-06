package main

import "fmt"

func main() {
	var i int
	var j int
	var answer int //couldn't figure out how to get it to hold a bigger int

	fmt.Println("\nEnter an integer value: ")
	_, err := fmt.Scanf("%d", &i)
	
	if err != nil {
		fmt.Println(err)
	}

	j = i 
	answer = i
	fmt.Print("\n", i,"! = ", i,)

	for j >= 2 {
		j--
		fmt.Print(" x ", j)
		answer *= j
	}

	fmt.Println(" =", answer)
}
//Resource: https://www.socketloop.com/tutorials/golang-how-to-read-integer-value-from-standard-input