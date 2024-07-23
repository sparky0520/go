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

// intArray := [...]uint8{128, 135, 119, 226, 255, 250}
// 	fmt.Println(intArray)

// 	fmt.Println(&intArray[0])
// 	fmt.Println(&intArray[1])
// 	fmt.Println(&intArray[2])

// var myString = []rune("Résumé")
// 	var indexed = myString[2]
// 	fmt.Printf("%v, %T\n", indexed, indexed)
// 	for i, v := range myString {
// 		fmt.Println(i, v)
// 	}
// 	fmt.Printf("\nThe length of myString is %v\n", len(myString))

// 	var myRune = 'a'
// 	fmt.Printf("\nmyRune = %v, Type = %T\n", myRune, myRune)

// 	var strSlice = []string{"s", "u", "s"}
// 	var strBuilder strings.Builder
// 	for i := range strSlice {
// 		strBuilder.WriteString(strSlice[i])
// 	}
// 	var catString = strBuilder.String()
// 	fmt.Printf("\n%v\n", catString)

// type gasEngine struct {
// 	gallons uint8
// 	mpg     uint8
// 	owner
// }
// type electricEngine struct {
// 	mpkwh uint8
// 	kwh   uint8
// 	owner
// }

// type owner struct {
// 	name string
// }

// func (e gasEngine) milesLeft() uint8 {
// 	return (e.gallons * e.mpg)
// }

// func (e electricEngine) milesLeft() uint8 {
// 	return (e.kwh * e.mpkwh)
// }

// func canMakeIt(e engine, miles uint8) {
// 	if miles <= e.milesLeft() {
// 		fmt.Println("You can make it!")
// 	} else {
// 		fmt.Println("Fuel up. Hurry!")
// 	}
// }

// // checks if whatever passed has a milesLeft() method with following signature
// type engine interface {
// 	milesLeft() uint8
// }

// var newEngine = gasEngine{8, 20, owner{"Darsh"}}
// 	fmt.Println(newEngine.name, newEngine.milesLeft())
// 	canMakeIt(newEngine, 150)
// 	canMakeIt(electricEngine{8, 12, owner{"Darsh"}}, 80)

// var p *int32 = new(int32)
// 	var i int32 = 9
// 	fmt.Printf("The value at *p is %v\n", *p)
// 	fmt.Printf("The value at i is %v\n", i)
// 	*p = 10
// 	fmt.Printf("The value at *p is %v\n", *p)
// 	p = &i
// 	fmt.Printf("The value at *p is %v\n", *p)
// 	*p = 3
// 	fmt.Printf("The value at i is %v\n", i)
// 	fmt.Printf("The value at *p is %v\n", *p)

// var thing1 = [5]float64{1, 2, 3, 4, 5}
// fmt.Printf("Memory Location of thing1: %p\n", &thing1)
// var result [5]float64 = square(&thing1)
// fmt.Printf("The result is: %v\n", result)
// fmt.Println("The value of thing1 is ", thing1)

// func square(thing2 *[5]float64) [5]float64 {
// 	fmt.Printf("The memory location of thing2: %p\n", thing2)
// 	fmt.Printf("The value of thing2: %v\n", *thing2)
// 	for i := range *thing2 {
// 		thing2[i] = thing2[i] * thing2[i]
// 	}
// 	return *thing2
// }
