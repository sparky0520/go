package main

import (
	"fmt"
	"sync"
	"time"
)

// var m = sync.RWMutex{}
// var dbData = []string{"id1", "id2", "id3", "id4", "id5"}
// var results = []string{}
var wg = sync.WaitGroup{}

func main() {
	// intArray := [...]uint8{128, 135, 119, 226, 255, 250}
	// fmt.Println(intArray)

	// var intSlice []int32 = []int32{1, 2, 3, 4, 5}
	// fmt.Printf("The length is %v with capacity %v\n", len(intSlice), cap(intSlice))
	// intSlice = append(intSlice, 7)
	// fmt.Printf("The length is %v with capacity %v\n", len(intSlice), cap(intSlice))

	// var intSlice2 []int32 = []int32{6, 7, 8, 9}
	// intSlice = append(intSlice, intSlice2...)
	// fmt.Println(intSlice)

	// var intSlice3 []int32 = make([]int32, 3, 5)
	// fmt.Println(intSlice3)

	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	var myMap2 = map[string]uint8{"Adam": 32, "Sarah": 30}
	fmt.Println(myMap2["Sarah"])

	var age, ok = myMap2["Jason"]
	if ok {
		fmt.Println(age)
	} else {
		fmt.Println(ok)
	}

	myMap2["Jason"] = 37

	age, ok = myMap2["Jason"]
	if ok {
		fmt.Println(age)
	} else {
		fmt.Println(ok)
	}

	for name, age := range myMap2 {
		fmt.Printf("Name: %v, Age: %v\n", name, age)
	}

	for i, v := range myMap2 {
		fmt.Printf("Index: %v, Value: %v\n", i, v)
	}
	i := 0
	for {
		if i >= 10 {
			break
		}
		fmt.Println(i)
		i++
	}
	i = 0
	for i < 10 {
		fmt.Println(i)
		i++
	}
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func count() {
	var res int
	for i := 0; i < 100000000; i++ {
		res++
	}
	wg.Done()
}

// func dbCall() {
// 	var delay float32 = 2000
// 	time.Sleep(time.Duration(delay) * time.Millisecond)
// 	save(dbData[i])
// 	log()
// 	wg.Done()
// }

// func save(result string) {
// 	m.Lock()
// 	results = append(results, result)
// 	m.Unlock()
// }

// func log() {
// 	m.RLock()
// 	fmt.Printf("\nThe current results are: %v\n", results)
// 	m.RUnlock()
// }
