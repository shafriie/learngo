package main

import (
	"fmt"
	"strconv"
)

func main (){

	//convert to int to string
	// nilai := 100000

	// nilaiString := strconv.Itoa(nilai)


	//convert to string to int
	bonus := "10000"

	bonusInt, _ := strconv.Atoi(bonus)

	bonusTotal := bonusInt + 90000

	fmt.Println("Hello world " , bonusTotal)
}