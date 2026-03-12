package main

import (
	"fmt"
	"sync"
	"time"
)

var mutex = sync.RWMutex{}
var waitGroup = sync.WaitGroup{}
var dbValues = []string{"String1", "String2", "String3", "String4", "String5"} 
var results = []string{}

func main() {
	startedAt := time.Now()

	// for index := range dbValues {
	// 	waitGroup.Add(1)
	// 	go dbCall(index)
	// }

	for i := 0; i <= 1000; i++ {
		waitGroup.Add(1)
		count()
	}
	waitGroup.Wait()

	fmt.Println("Execution time: ", time.Since(startedAt))
	fmt.Println(results)
}

func dbCall(index int) {
	time.Sleep(time.Duration(2000) * time.Millisecond)
	save(dbValues[index])
	log()
	waitGroup.Done()
}

func save(result string) {
	mutex.Lock()
	results = append(results, result)
	mutex.Unlock()
}

func log() {
	mutex.RLock()
	fmt.Println("Results until now:", results)
	mutex.RUnlock()
}

func count() {
	var sum int32 = 0

	for i :=  0; i <= 100000000; i++ {
		sum += 1
	}

	waitGroup.Done()
}