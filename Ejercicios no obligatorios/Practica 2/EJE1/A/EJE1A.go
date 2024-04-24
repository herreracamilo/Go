package main

import (
	"fmt"
	"log"
	"sort"
	"math"
)

func main(){
	var vectorTemperaturas [10]float64
	var mapTemperaturas = map[string]float64{
		"high":      0,
        "normal":    0,
        "low":       0,
        "incorrect": 0,
	}
	altaC:=0
	normalC:=0
	bajaC:=0
	llenarVector((*[10]float64)(&vectorTemperaturas))
	fmt.Println(vectorTemperaturas)
	dividirGrupos((*[10]float64)(&vectorTemperaturas),mapTemperaturas,&altaC,&normalC,&bajaC)
	fmt.Println("El porcentaje de pacientes con alta temperatura es:",((altaC *10) / 10),"%")
	fmt.Println("El porcentaje de pacientes con temperatura normal es:",((normalC *10) / 10),"%")
	fmt.Println("El porcentaje de pacientes con baja temperatura es:",((bajaC *10) / 10),"%")

	
	// para sacar el promedio entero entre la temp minima y maxima, ordeno el vector saco el minimo y el maximo y lo divido entre 2
	tempSlices:= vectorTemperaturas[:]
	sort.Float64s(tempSlices) // para ordenarlo con el sort necesito un slice
	fmt.Println()
	fmt.Println("slices ordenado")
	fmt.Println(tempSlices)
	temperaturaMinima:= tempSlices[0]
	temperaturaMaxima:=tempSlices[9]
	promedioTemperaturas:= ((temperaturaMinima + temperaturaMaxima) / 2)
	promedioTemperaturasINT:= int(promedioTemperaturas)
	fmt.Println()
	fmt.Print("El promedio entero de temperaturas entre la maxima y la minima es: ", promedioTemperaturasINT)
	fmt.Println()
	fmt.Println("El promedio entero de temperaturas entre la maxima y la minima es: ",math.Round(promedioTemperaturas)) // preguntar si uso esta o convertir a entero (uso otra variable más)

	// A usando un arreglo o un Map de tres posiciones donde se acumulan los valores de cada grupo
	fmt.Println()
	fmt.Println(mapTemperaturas)
	
}


func read(temperatura *float64) {
    _,err := fmt.Scanln(temperatura)
    if err != nil {
        log.Fatal(err)
    }
}

func llenarVector(vectorTemperaturas *[10]float64){
	var temperatura float64
	for i:= 0; i<10; i++{
		read(&temperatura)
		vectorTemperaturas[i] = temperatura
	}
}

func dividirGrupos(vectorTemperaturas *[10]float64, mapTemperaturas map[string]float64,altaC,normalC,bajaC *int){
	
	for i:=0; i < 10; i++ {
		switch {
		case vectorTemperaturas[i] > 37.5 && vectorTemperaturas[i] < 50.00 :
			mapTemperaturas["high"]+= vectorTemperaturas[i]
			(*altaC)+=1
		case vectorTemperaturas[i] >= 36.00 && vectorTemperaturas[i] <= 37.5:
			mapTemperaturas["normal"]+=vectorTemperaturas[i]
			(*normalC)+=1
		case vectorTemperaturas[i] >= 20.00 && vectorTemperaturas[i] < 36.00:
			mapTemperaturas["low"]+=vectorTemperaturas[i]
			(*bajaC)+=1
		case vectorTemperaturas[i] < 20.00 || vectorTemperaturas[i] > 50:
			mapTemperaturas["incorrect"]+=vectorTemperaturas[i]
		}
	}
}