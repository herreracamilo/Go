/*
	Realice un programa que reciba una frase e imprima en pantalla la

misma frase reemplazando las ocurrencias de “jueves” por “martes”
respetando las letras minúsculas o mayúsculas de la palabra original en
su posición correspondiente. Por ejemplo, se reemplaza “Jueves” por
“Martes” o “jueveS” por “marteS”.
*/
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
    var frase string
	read(&frase)
	nuevaFrase := crearFrase(frase, "jueves", "martes")
	fmt.Println(nuevaFrase)

}

func read(frase *string) {
    fmt.Println("Ingrese una frase con la palabra 'jueves': ")
    reader := bufio.NewReader(os.Stdin)
    input, err := reader.ReadString('\n')
    if err != nil {
        log.Fatal(err)
    }
    *frase = input // asigno el valor leído a la variable frase
}

func crearFrase(frase, palabraOriginal, palabraNueva string) string {
	nuevaFrase := frase
	for {
		index := strings.Index(strings.ToLower(nuevaFrase), palabraOriginal)
		if index == -1 {
			break
		}
		nuevaFrase = nuevaFrase[:index] + doyFormato(nuevaFrase[index:index+len(palabraOriginal)], palabraNueva) + nuevaFrase[index+len(palabraOriginal):]
	}
	return nuevaFrase
}

func doyFormato(palabraOriginal, palabraNueva string) string {
	var nuevaPalabra strings.Builder
	for i := 0; i < len(palabraOriginal); i++ { // len te devuelve la cantidad de caracteres
		if 'A' <= palabraOriginal[i] && palabraOriginal[i] <= 'Z' {
			nuevaPalabra.WriteByte(palabraNueva[i] - ('a' - 'A')) // la resta es para que se ponga en mayuscula
		} else {
			nuevaPalabra.WriteByte(palabraNueva[i]) // si no la tengo que poner en mayuscula, la escribo normal
		}
	}
	return nuevaPalabra.String()
}
