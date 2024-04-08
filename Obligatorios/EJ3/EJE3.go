/* Realice un programa que reciba una palabra como argumento y lee de
la entrada una frase. Luego, el programa debe imprimir la frase que
leyó con cada una de las ocurrencias de la palabra con las mayúsculas
y minúsculas invertidas. Por ejemplo, si la frase es:
“Parece peqUEño, pero no es tan pequeÑo el PEQUEÑO”
y la palabra es “PEQUEÑO” entonces el programa imprimirá:
“Parece PEQueÑO, pero no es tan PEQUEñO el pequeño”
Tenga en cuenta que la palabra a buscar puede ser ingresada con
mayúsculas y minúsculas mezcladas. */

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
	readFrase(&frase) // leo la frase 
	readPalabra(&palabra) // leo la palabra la cual darle formato
	nuevaFrase := crearFrase(frase, palabra) // le asigno a nueva frase la funcion que recibe la frase original y la palabra a darle formato
	fmt.Println(nuevaFrase) // imprimo la nueva frase

}

// funcion que lee desde teclado la frase
func readFrase(frase *string) {
	fmt.Println("ingrese la frase en la cual reemplazar:")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	*frase = strings.TrimSpace(input) // saco los espacios en blanco alrededor de la frase puede dar errores
}

// funcion que lee desde teclado la palabra
func readPalabra(palabra *string) {
	fmt.Println("ingrese la palabra:")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	*palabra = strings.TrimSpace(input) // saco los espacios en blanco alrededor de la palabra
}

func crearFrase(frase, palabra string) string {
	var nuevaFrase strings.Builder // el strings.Builder lo uso para construir el nuevo string a partir del uso de runas y no de bytes para respetar los caracteres especiales
	indice := 0	// creo un indice para saber cuanto ocupa la palabra y darle el formato correcto				
	for {
		index := strings.Index(strings.ToLower(frase[indice:]), strings.ToLower(palabra)) // busco una ocurrencia de la palabra desde donde comienza la frase (indice=0) para adelante
		if index == -1 { // si no hay ocurrencias salgo
			break
		}
		nuevaFrase.WriteString(frase[indice:indice+index]) // escribo la frase en el mismo formato desde antes de la palabra o sea de donde no hay palabra hasta la primera letra de la palabra
		nuevaFrase.WriteString(doyFormato(frase[indice+index : indice+index+len(palabra)])) // aca le doy formato a la palabra desde donde comienza hasta donde termina
		indice = indice + index + len(palabra)   // actualizamos el índice --> indice + index + la cantidad de caracteres que tenga la palabra                                      
	}
	nuevaFrase.WriteString(frase[indice:]) // agrego el resto de la frase que queda después de la última ocurrencia de la palabra

	return nuevaFrase.String() // devuelvo el string
}


func doyFormato(palabra string) string {
	var nuevaPalabra strings.Builder // uso strings.Builder para crear la nueva frase a partir de runas
	for _, r:= range palabra { // hago un for para recorrer caracter por caracter. como no necesito el indice del for para nada se pone _ en go
							   // go al usar range itera los string en runas, por eso r toma este tipo
		nuevaPalabra.WriteRune(unicode.SimpleFold(r)) // con el simplefold no necesito corroborar si es min o mayus ya que te da el inverso si o si
	}
	return nuevaPalabra.String()
}

