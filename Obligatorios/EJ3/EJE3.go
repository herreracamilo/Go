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
	var palabra string
	readFrase(&frase)
	readPalabra(&palabra)
	nuevaFrase := crearFrase(frase, palabra)
	fmt.Println(nuevaFrase)

}

func readFrase(frase *string) {
	fmt.Println("ingrese la frase en la cual reemplazar")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	*frase = input // asigno el valor leído a la variable frase
}

func readPalabra(palabra *string) {
	fmt.Println("ingrese la palabra")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	*palabra = input // asigno el valor leído a la variable frase
}

func crearFrase(frase, palabra string) string {
	nuevaFrase := frase
	for {
		index := strings.Index(strings.ToLower(nuevaFrase), palabra)
		if index == -1 {
			break
		}
		nuevaFrase = nuevaFrase[:index] + doyFormato(nuevaFrase[index:index+len(palabra)], palabra) + nuevaFrase[index+len(palabra):]
	}

	return nuevaFrase
}

func doyFormato(palabraOriginal, palabraNueva string) string {
	var nuevaPalabra strings.Builder
	for i := 0; i < len(palabraOriginal); i++ {
		if 'A' <= palabraOriginal[i] && palabraOriginal[i] <= 'Z' {
			// si la letra en palabraOriginal es mayúscula, la convierte a minúscula
			nuevaPalabra.WriteByte(palabraNueva[i] + ('a' - 'A'))
		} else if 'a' <= palabraOriginal[i] && palabraOriginal[i] <= 'z' {
			// si la letra en palabraOriginal es minúscula, la convierte a mayúscula
			nuevaPalabra.WriteByte(palabraNueva[i] - ('a' - 'A'))
		} else {
			// si la letra no es una letra del alfabeto, la agrega tal cual
			nuevaPalabra.WriteByte(palabraNueva[i])
		}
	}
	return nuevaPalabra.String()
}
