/* Realice las modificaciones necesarias al ejercicio anterior para que en
lugar de reemplazar la palabra “jueves” por “martes” ahora se
reemplace “miércoles” por “automóvil”. Piense qué impacto tuvieron
esas modificaciones en el programa que había realizado.*/

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
	read(&frase)
	nuevaFrase := crearFrase(frase, "miércoles", "automóvil")
	fmt.Println(nuevaFrase)

}

func read(frase *string) {
    fmt.Println("Ingrese una frase con la palabra miercoles: ")
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
	runaOriginal := []rune(palabraOriginal)
	runaNueva := []rune(palabraNueva)
	for i, c:= range runaOriginal{
		if unicode.IsUpper(c) {
			nuevaPalabra.WriteRune(unicode.ToUpper(runaNueva[i])) // pongo en mayuscula la letra de automóvil en el indice que esta mayuscula en miercoles
		} else {
			nuevaPalabra.WriteRune(((runaNueva[i]))) // si no la tengo que poner en mayuscula, la escribo normal
		}
	}
	return nuevaPalabra.String()
}
