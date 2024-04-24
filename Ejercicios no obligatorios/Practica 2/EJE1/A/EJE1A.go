package main

import (
	"fmt"
	"log"
	"sort"
	"math"
)

func main(){
	var vectorTemperaturas [10]float64
	var vector3Temperaturas[3]float64
	altaC:=0
	normalC:=0
	bajaC:=0
	llenarVectorA((*[10]float64)(&vectorTemperaturas))
	fmt.Println(vectorTemperaturas)
	dividirGruposA((*[10]float64)(&vectorTemperaturas),(*[3]float64)(&vector3Temperaturas),&altaC,&normalC,&bajaC)
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
	fmt.Println(vector3Temperaturas)
	
}


func readA(temperatura *float64) {
    _,err := fmt.Scanln(temperatura)
    if err != nil {
        log.Fatal(err)
    }
}

func llenarVectorA(vectorTemperaturas *[10]float64){
	var temperatura float64
	for i:= 0; i<10; i++{
		readA(&temperatura)
		vectorTemperaturas[i] = temperatura
	}
}

func dividirGruposA(vectorTemperaturas *[10]float64, vector3Temperaturas *[3]float64,altaC,normalC,bajaC *int){
	
	for i:=0; i < 10; i++ {
		switch {
		case vectorTemperaturas[i] > 37.5:
			vector3Temperaturas[0]+= vectorTemperaturas[i]
			*altaC+=1
		case vectorTemperaturas[i] >= 36.00 && vectorTemperaturas[i] <= 37.5:
			vector3Temperaturas[1]+=vectorTemperaturas[i]
			*normalC+=1
		case vectorTemperaturas[i] < 36.00:
			vector3Temperaturas[2]+=vectorTemperaturas[i]
			*bajaC+=1
		}
	}
}