package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
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
	contador := 0
	for {
		index := strings.Index(strings.ToLower(nuevaFrase[contador:]), palabra) // es en donde buscar la palabra, la palabra a buscar
		
		if index == -1 {
			break
		}
		suma:= index + contador
		nuevaFrase = nuevaFrase[:suma] + doyFormato(nuevaFrase[suma:suma+len(palabra)]) + nuevaFrase[suma+len(palabra):]
		contador = suma+len(palabra)
	}

	return nuevaFrase
}

func doyFormato(palabraOriginal string) string {
	var nuevaPalabra strings.Builder
	for i := 0; i < len(palabraOriginal); i++ {
		nuevaPalabra.WriteRune(unicode.SimpleFold(rune(palabraOriginal[i])))
	}
	return nuevaPalabra.String()
}
