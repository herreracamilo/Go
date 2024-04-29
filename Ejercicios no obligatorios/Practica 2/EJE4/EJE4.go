package main

import (
	"fmt"
	"log"
	"math"
)

const N = 5

func main()  {
	var arregloX[N]float64
	var arregloY[N]float64
	var arregloZ[N]float64
	llenarVectores((*[5]float64)(&arregloX),(*[5]float64)(&arregloY),(*[5]float64)(&arregloZ))
	fmt.Println(arregloX)
	fmt.Println(arregloY)
	fmt.Println(arregloZ)
	fmt.Println(sumatoria((&arregloX)))
	fmt.Println(productoria((&arregloZ)))
}

func read(valor *float64) {
    _,err := fmt.Scanln(valor)
    if err != nil {
        log.Fatal(err)
    }
}

func llenarVectores(arregloX,arregloY,arregloZ *[N]float64) {
	var valor float64
	for i:=0 ; i<N ; i++{
		read(&valor)
		arregloX[i]= valor;
		read(&valor)
		arregloY[i]= valor;
		read(&valor)
		arregloZ[i]= valor;
	}
}

func sumatoria(arregloX *[N]float64) float64 {
	var result float64
	for i := 0; i < N; i++ {
		result+= 1/(arregloX[i])
	}
	return result
}

func productoria(arregloZ *[N]float64) float64  {
	var result float64
	for i := 0; i < N; i++ {
		result+= math.Pow(arregloZ[i],3.00)
	}
	return result
}


