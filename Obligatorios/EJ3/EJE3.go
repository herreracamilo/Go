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
	fmt.Println("ingrese la frase en la cual reemplazar:")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	*frase = strings.TrimSpace(input) // saco los espacios en blanco alrededor de la frase puede dar errores
}

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
	var nuevaFrase strings.Builder // el strings.Builder lo uso para construir eficientemente cadenas de caracteres mediante la concatenación secuencial de partes de texto
	indice := 0					
	for {
		index := strings.Index(strings.ToLower(frase[indice:]), strings.ToLower(palabra))
		if index == -1 {
			break
		}
		nuevaFrase.WriteString(frase[indice:indice+index])                     // parte antes de la palabra o sea de donde no hay palabra hasta la primera letra de la palabra
		nuevaFrase.WriteString(doyFormato(frase[indice+index : indice+index+len(palabra)])) // la palabra con el formato deseado
		indice = indice + index + len(palabra)   // actualizamos el índice --> indice + index + la cantidad de caracteres que tenga la palabra                                      
	}

	
	nuevaFrase.WriteString(frase[indice:]) // agrego el resto de la frase que queda después de la última ocurrencia de la palabra

	return nuevaFrase.String()
}



func doyFormato(palabra string) string {
	var nuevaPalabra strings.Builder
	for i := 0; i < len(palabra); i++ { // hago un for para recorrer caracter por caracter
		r := rune(palabra[i]) // transformo la palabra en runa
		if unicode.IsUpper(r) { // si esta en mayuscula
			nuevaPalabra.WriteRune(unicode.ToLower(r)) // entonces la hago minuscula
		} else if unicode.IsLower(r) { // si esta en miniscula
			nuevaPalabra.WriteRune(unicode.ToUpper(r)) // entonces la hago mayuscula
		} else {
			nuevaPalabra.WriteRune(r) // si es un caracter especial lo dejo asi
		}
	}
	return nuevaPalabra.String()
}


// PREGUNTAR ESTE
/*
func doyFormato(palabra string) string {
	var nuevaPalabra strings.Builder
	for i := 0; i < len(palabra); i++ { // hago un for para recorrer caracter por caracter
		r := rune(palabra[i]) // transformo la palabra en runa
		nuevaPalabra.WriteRune(unicode.SimpleFold(r)) // con el simplefold no necesito corroborar si es min o mayus ya que te da el inverso si o si
	}
	return nuevaPalabra.String()
}
*/