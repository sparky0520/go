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
	t0 := time.Now()
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go count()
		// go dbCall()
	}
	wg.Wait()
	fmt.Printf("\nTotal Execution Time: %v\n", time.Since(t0))
	// fmt.Println("The results are ", results)
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
