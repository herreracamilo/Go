package main

import (
	"fmt"
	"log"
	"sort"
	"math"
)

func main(){
	var vectorTemperaturas [10]float64
	alta:= 0.0
	normal:= 0.0
	baja:= 0.0
	altaC:=0
	normalC:=0
	bajaC:=0
	llenarVector((*[10]float64)(&vectorTemperaturas))
	fmt.Println(vectorTemperaturas)
	dividirGrupos((*[10]float64)(&vectorTemperaturas),&alta,&normal,&baja,&altaC,&normalC,&bajaC)
	fmt.Print("El porcentaje de pacientes con alta temperatura es: ")
	fmt.Println((altaC *10) / 10)
	fmt.Print("El porcentaje de pacientes con temperatura normal es: ")
	fmt.Println((normalC *10) / 10)
	fmt.Print("El porcentaje de pacientes con baja temperatura es: ")
	fmt.Println((bajaC *10) / 10) 
	
	// para sacar el promedio entero entre la temp minima y maxima, ordeno el vector saco el minimo y el maximo y lo divido entre 2
	tempSlices:= vectorTemperaturas[:]
	sort.Float64s(tempSlices) // para ordenarlo con el sort necesito un slice
	fmt.Println(tempSlices)
	temperaturaMinima:= tempSlices[0]
	temperaturaMaxima:=tempSlices[9]
	promedioTemperaturas:= ((temperaturaMinima + temperaturaMaxima) / 2)
	promedioTemperaturasINT:= int(promedioTemperaturas)
	fmt.Print("El promedio entero de temperaturas entre la maxima y la minima es: ")
	fmt.Println(promedioTemperaturasINT) 
	fmt.Println(math.Round(promedioTemperaturas)) // preguntar si uso esta o convertir a entero (uso otra variable más)
	
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

func dividirGrupos(vectorTemperaturas *[10]float64, alta,normal,baja *float64,altaC,normalC,bajaC *int){
	i:=0
	for {
		switch {
		case vectorTemperaturas[i] > 37.5:
			*alta+= vectorTemperaturas[i]
			*altaC+=1
		case vectorTemperaturas[i] >= 36.00 && vectorTemperaturas[i] <= 37.5:
			*normal+=vectorTemperaturas[i]
			*normalC+=1
		case vectorTemperaturas[i] < 36.00:
			*baja+=vectorTemperaturas[i]
			*bajaC+=1
		}
		i++
		if(i >= 10){
			break
		}
	}
}