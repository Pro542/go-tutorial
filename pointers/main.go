package main

import (
	"fmt"
	"runtime"
)

func main() {
	var m runtime.MemStats

	// A - Copy by value
	runtime.ReadMemStats(&m)
	beforeA := m.TotalAlloc

	var thing1 = [5]float64{1, 2, 3, 4, 5}
	fmt.Printf("\nThe memory location of the thing1 array is: %p", &thing1)
	var result [5]float64 = square_copy(thing1)
	fmt.Printf("\nThe result is: %v", result)
	fmt.Printf("\nThe value of thing1 is: %v", thing1)

	runtime.ReadMemStats(&m)
	afterA := m.TotalAlloc
	usedByA := afterA - beforeA

	runtime.ReadMemStats(&m)
	beforeB := m.TotalAlloc

	// B - Copy by reference
	var thing3 = [5]float64{1, 2, 3, 4, 5}
	fmt.Printf("\nThe memory location of the thing3 array is: %p", &thing3)
	var result2 [5]float64 = square_ref(&thing3)
	fmt.Printf("\nThe result2 is: %v", result2)
	fmt.Printf("\nThe value of thing3 is: %v", thing3)

	runtime.ReadMemStats(&m)
	afterB := m.TotalAlloc
	usedByB := afterB - beforeB

	fmt.Printf("\n\nsquare_copy allocated %d bytes", usedByA)
	fmt.Printf("\nsquare_ref allocated %d bytes", usedByB)
	fmt.Printf("\nThat's a ~%dx difference!", usedByA/usedByB)
}

// this one copies the value, uses lots of memory
func square_copy(thing2 [5]float64) [5]float64 {
	fmt.Printf("\nThe memory location of the thing2 array is: %p", &thing2)
	for i := range thing2 {
		thing2[i] = thing2[i]*thing2[i]
	}
	return thing2
}

// this one uses the same value, saving memory
// (pay attention to where `*` and `&` are used)
func square_ref(thing4 *[5]float64) [5]float64 {
	fmt.Printf("\nThe memory location of the thing4 array is: %p", thing4)

	// we don't need dereferencing here because compiler does it automatically...
	for i := range thing4 {
		thing4[i] = thing4[i]*thing4[i]
	}

	// ...unless the exact values are needed.
	return *thing4


	// Go dereferences automatically only for selection operations: indexing, field access,
	// method calls. Not when the whole value is needed
}
