package main

import (
	"fmt"
	"log"
)

func main(){
	var temperatura float64
	var vectorTemperaturas [10]float64
	alta:= 0.0
	normal:= 0.0
	baja:= 0.0
	llenarVector(temperatura, (*[10]float64)(&vectorTemperaturas))
	fmt.Println(vectorTemperaturas)
	dividirGrupos((*[10]float64)(&vectorTemperaturas),&alta,&normal,&baja)
	fmt.Println(alta)
	fmt.Println(normal)
	fmt.Println(baja)

}


func read(temperatura *float64) {
    _,err := fmt.Scanln(temperatura)
    if err != nil {
        log.Fatal(err)
    }
}

func llenarVector(temperatura float64, vectorTemperaturas *[10]float64){
	i:= 0
	for {
		read(&temperatura)
		vectorTemperaturas[i] = temperatura
		i++
		if(i >= 10){
			break
		}
	}
}

func dividirGrupos(vectorTemperaturas *[10]float64, alta,normal,baja *float64){
	i:=0
	for {
		switch {
		case vectorTemperaturas[i] > 37.5:
			*alta+= vectorTemperaturas[i]
		case vectorTemperaturas[i] >= 36.00 && vectorTemperaturas[i] <= 37.5:
			*normal+=vectorTemperaturas[i]
		case vectorTemperaturas[i] < 36.00:
			*baja+=vectorTemperaturas[i]
		}
		i++
		if(i >= 10){
			break
		}
	}
}