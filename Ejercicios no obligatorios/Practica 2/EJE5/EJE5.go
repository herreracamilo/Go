package main

import (
	"fmt"
	"log"
)

type Vector [10]float64

func main()  {
	var v1 Vector
	var v2 Vector
	var f float64
	Initialize(&v1,f)
	Initialize(&v2,f)
	fmt.Println(v1)
	fmt.Println(v2)
	suma:=Sum(v1,v2)
	fmt.Println(suma)
	SumInPlace(&v1,v2)
	fmt.Println(v1)

}

func read(valor *float64) {
    _,err := fmt.Scanln(valor)
    if err != nil {
        log.Fatal(err)
    }
}

func Initialize(v *Vector ,f float64){
	for i := 0; i < len(v); i++ {
		read(&f)
		v[i] = f
	}

}


func Sum(v1 , v2 Vector) Vector{
	var sumado Vector
	for i := 0; i < len(v1); i++ {
		sumado[i] = (v1[i] + v2[i])
	}
	return sumado
}


func SumInPlace(v1 *Vector, v2 Vector){
	for i := 0; i < len(v1); i++ {
		v1[i] = (v1[i] + v2[i])
	}
}
