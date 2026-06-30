package main

import (
	"fmt"
	"math/rand"
	"time"
	"sync"
)

var dbData = []string{"id1", "id2", "id3", "id4", "id5", "id6"}

func main() {
	sequentialRun()
	concurrentRun()
	sharedMemoryRun()
	mutexRun()
}

func sequentialRun() {
	fmt.Printf("\n=== Sequential Run ===\n")
	t0 := time.Now()
	for i:=0; i<len(dbData); i++ {
		dbCallSequential(i)
	}
	fmt.Printf("\nTotal execution time: %v\n", time.Since(t0))
}
func dbCallSequential(i int) {
	// Simulate DB call delay
	var delay float32 = rand.Float32()*400
	time.Sleep(time.Duration(delay)*time.Millisecond)
	fmt.Println("The result from the database is:", dbData[i])
}

// go starts a goroutine (async) for that func and moves program to next line
// WaitGroup is a counter to let us block goroutines as needed
var wg = sync.WaitGroup{}
func concurrentRun() {
	fmt.Printf("\n=== Concurrent Run ===\n")
	t0 := time.Now()
	for i:=0; i<len(dbData); i++ {
		wg.Add(1)
		go dbCallConcurrent(i)
	}
	wg.Wait() // goes forward here only when counter is 0
	fmt.Printf("\nTotal execution time: %v\n", time.Since(t0))
}
func dbCallConcurrent(i int) {
	var delay float32 = rand.Float32()*400
	time.Sleep(time.Duration(delay)*time.Millisecond)
	fmt.Println("The result from the database is:", dbData[i])
	wg.Done()
}

// shared memory
// results can be random (missing ids)
var results = []string{}
var wg2 = sync.WaitGroup{}
func sharedMemoryRun() {
	fmt.Printf("\n=== Shared Memory Run ===\n")
	t0 := time.Now()
	for i:=0; i<len(dbData); i++ {
		wg.Add(1)
		go dbCallSharedMemory(i)
	}
	wg.Wait() // goes forward here only when counter is 0
	fmt.Printf("\nTotal execution time: %v\n", time.Since(t0))
	fmt.Printf("\nThe results are: %v\n", results)
}
func dbCallSharedMemory(i int) {
	var delay float32 = 400 // removed randomness here
	time.Sleep(time.Duration(delay)*time.Millisecond)
	fmt.Println("The result from the database is:", dbData[i])
	results = append(results, dbData[i])
	wg.Done()
}

// Mutex
// Locks code so it's mutually exclusive to one goroutine
// shared memory won't overwrite each other
var m = sync.Mutex{}
var results2 = []string{}
var wg3 = sync.WaitGroup{}
func mutexRun() {
	fmt.Printf("\n=== Mutex Run ===\n")
	t0 := time.Now()
	for i:=0; i<len(dbData); i++ {
		wg3.Add(1)
		go dbCallMutex(i)
	}
	wg3.Wait() // goes forward here only when counter is 0
	fmt.Printf("\nTotal execution time: %v\n", time.Since(t0))
	fmt.Printf("\nThe results are: %v\n", results2)
}
func dbCallMutex(i int) {
	var delay float32 = 400
	// if you lock here, every goroutine will Sleep 
	// m.Lock()
	time.Sleep(time.Duration(delay)*time.Millisecond)
	fmt.Println("The result from the database is:", dbData[i])
	m.Lock()
	results2 = append(results2, dbData[i])
	m.Unlock()
	wg3.Done()
}

// Read, Write granular Mutex
var m2 = sync.RWMutex{}
var results3 = []string{}
var wg4 = sync.WaitGroup{}
func readMutexRun() {
	fmt.Printf("\n=== Read Mutex Run ===\n")
	t0 := time.Now()
	for i:=0; i<len(dbData); i++ {
		wg4.Add(1)
		go dbCallReadMutex(i)
	}
	wg4.Wait() // goes forward here only when counter is 0
	fmt.Printf("\nTotal execution time: %v\n", time.Since(t0))
	fmt.Printf("\nThe results are: %v\n", results3)
}
func dbCallReadMutex(i int) {
	var delay float32 = 400
	time.Sleep(time.Duration(delay)*time.Millisecond)
	fmt.Println("The result from the database is:", dbData[i])
	save(dbData[i])
	log()
	wg4.Done()
}
func save(result string) {
	m2.Lock() // full lock
	results = append(results, result)
	m2.Unlock()
}
func log() {
	m2.RLock()
	fmt.Printf("\nThe current results are: %v\n", results)
	m2.RUnlock()
}
