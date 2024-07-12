package main

import "fmt"

func main() {
	intArray := [...]uint8{128, 135, 119, 226, 255, 250}
	fmt.Println(intArray)

	fmt.Println(&intArray[0])
	fmt.Println(&intArray[1])
	fmt.Println(&intArray[2])
}
