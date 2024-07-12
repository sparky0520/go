package main

import (
	"errors"
	"fmt"
)

func side() {
	name := "Vasu"
	printHello(name)

	numerator := 44
	denominator := 0
	var quotient, remainder, err = intDivision(numerator, denominator)
	if err != nil { // If error is not empty for say
		fmt.Println(err)
	} else if remainder == 0 {
		fmt.Printf("Quotient returned: %v without remainder", quotient)
	} else {
		fmt.Printf("Quotient returned: %v and Remainder returned: %v", quotient, remainder)
	}
}

func printHello(name string) {
	fmt.Printf("Hello %v!\n", name)
}

func intDivision(numerator int, denominator int) (int, int, error) {
	var myError error // defaults to nil
	if denominator == 0 {
		myError = errors.New("Cannot divide by 0 !")
		return 0, 0, myError
	}
	var quotient int = numerator / denominator
	var remainder int = numerator % denominator
	return quotient, remainder, myError
}
