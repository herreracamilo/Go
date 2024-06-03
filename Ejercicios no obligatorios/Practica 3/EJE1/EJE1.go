package main

import (
	"fmt"
)

type Map[K comparable, V any] map[K]V


func main()  {
	ints:= map[string]int{
		"primero": 88,
		"segundo": 99,
		"tercero":77,
	}

	floats:= map[string]float64{
		"primero": 73.65,
		"segundo": 91.33,
		"tercero": 777.34,
	}

	intToStringMap := map[int]string{
		1: "Uno",
		2: "Dos",
		3: "Tres",
	}

	stringToFloatMap := map[string]float64{
		"A": 1.1,
		"B": 2.2,
		"C": 3.3,
	}

	fmt.Printf("Generic Sums: %v and %v\n",
	suma[string, int](ints),
	suma[string,float64](floats),
)
	printMap(stringToFloatMap)
	printMap(intToStringMap)
}


func suma [K comparable, V int | float64] (m map[K]V) V {
	var r V
	for _, v:= range m{
		r+=v
	}
	return r
}

func printMap[K comparable, V any](m map[K]V) {
	for key, value := range m {
		fmt.Printf("%v: %v\n", key, value)
	}
	fmt.Println()
}