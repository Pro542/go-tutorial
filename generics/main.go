package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	var intSlice = []int{5, 4, 2}
	var float32Slice = []float32{54.2, 12.1, 34.5}

	fmt.Printf("int sum: %v\n", sumSlice[int](intSlice))
	fmt.Printf("float32 sum: %v\n", sumSlice[float32](float32Slice))

	var strSlice = []string{"foo", "bar", "baz"}
	var emptySlice = []bool{}

	// isEmpty[string](strSlice) type here is optional, sometimes not
	fmt.Printf("strSlice isEmpty: %v\n", isEmpty(strSlice))
	fmt.Printf("emptySlice isEmpty: %v\n", isEmpty(emptySlice))

	// where type is not optional (compiler can't infer otherwise)
	var contacts []contactInfo = loadJSON[contactInfo]("./contactInfo.json")
	fmt.Printf("\n%+v", contacts)

	var purchases []purchaseInfo = loadJSON[purchaseInfo]("./purchaseInfo.json")
	fmt.Printf("\n%+v", purchases)

	carInstances()
}

// T is generic, returns T type as well
func sumSlice[T int | float32](slice []T) T {
	var sum T
	for _, v := range slice {
		sum += v
	}
	return sum
}

// any
func isEmpty[T any](slice []T) bool {
	return len(slice) == 0
}

// type needed to unmarshal JSON
type contactInfo struct {
	Name  string
	Email string
}
type purchaseInfo struct {
	Name   string
	Price  float32
	Amount int
}

func loadJSON[T contactInfo | purchaseInfo](filePath string) []T {
	data, _ := os.ReadFile(filePath)

	var loaded = []T{}
	json.Unmarshal(data, &loaded)

	return loaded
}

// Using generics with struct types
type gasEngine struct {
	gallons float32
	mpg     float32
}

type electricEngine struct {
	kwh   float32
	mpkwh float32
}

type car[T gasEngine | electricEngine] struct {
	carMake  string
	carModel string
	engine   T
}

func carInstances() {
	var gasCar = car[gasEngine]{
		carMake:  "Honda",
		carModel: "Civic",
		engine: gasEngine{
			gallons: 12.4,
			mpg:     40,
		},
	}

	var electricCar = car[electricEngine]{
		carMake:  "Tesla",
		carModel: "Model 3",
		engine: electricEngine{
			kwh:   57.5,
			mpkwh: 4.17,
		},
	}

	fmt.Println(gasCar)
	fmt.Println(electricCar)
}
