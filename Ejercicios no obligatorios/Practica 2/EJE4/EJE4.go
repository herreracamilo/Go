package main

import (
	"fmt"
	"log"
	"math"
	"sort"
)

const N = 5


func main()  {
	var arregloX[N]float64
	var arregloY[N]float64
	var arregloZ[N]float64
	llenarVectores((&arregloX),(&arregloY),(&arregloZ))
	min, max := maxYmin((&arregloY))
	R:= (sumatoria(&arregloX)-productoria(&arregloZ)*(max*min))
	fmt.Print("El resultado del problema es --> ")
	fmt.Print(R)


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

func maxYmin(arregloY *[N]float64) (float64,float64 )  {
	sliceX:= arregloY[:]
	sort.Float64s(sliceX)
	min:= sliceX[0]
	max:= sliceX[len(sliceX)-1]
	return min, max
}


