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
	read(&frase) // leo desde teclado la frase que contenga miércoles
	nuevaFrase := crearFrase(frase, "miércoles", "automóvil") // le asigno a nuevafrase el crearfrase que devolverá el nuevo string y recibe la palabra a reemplazar y la palabra nueva
	fmt.Println(nuevaFrase) // imprimo la nueva frase

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

// creo la frase, recibo la frase introducida por el usuario, la palabra a buscar en la frase y la palabra nueva que se escribirá
func crearFrase(frase, palabraOriginal, palabraNueva string) string {
	nuevaFrase := frase // asigno a nuevafrase la frase original
	for {
		index := strings.Index(strings.ToLower(nuevaFrase), strings.ToLower(palabraOriginal)) // transformo la nueva palabra y la original todo a minuscula para buscar el indice de donde esta en la frase
		if index == -1 { // si no encuentro coincidencias salgo del bucle
			break
		}
		// nuevaFrase[:index] --> escribe la frase desde el principio hasta donde se encontró la palabra de la misma forma en la que está
		// doyFormato(nuevaFrase[index:index+len(palabraOriginal)], palabraNueva)  doy formato toma la palabra en la frase desde donde empieza hasta donde termina y le cambia el formato
		// nuevaFrase[index+len(palabraOriginal):] escribo desde donde termina la palabra para adelante la frase como sigue
		// todo esto se concatena para formar el nuevo string
		nuevaFrase = nuevaFrase[:index] + doyFormato(nuevaFrase[index:index+len(palabraOriginal)], palabraNueva) + nuevaFrase[index+len(palabraOriginal):]
	}
	return nuevaFrase
}

func doyFormato(palabraOriginal, palabraNueva string) string {
	var nuevaPalabra strings.Builder // uso strings.Builder para crear la nueva frase a partir de runas
	runaOriginal := []rune(palabraOriginal) // paso el string de la palabra original a runa para poder trabajar con unicode
	runaNueva := []rune(palabraNueva) // paso el string de la palabra nueva a runa para poder trabajar con unicode
	for i, c:= range runaOriginal{
		if unicode.IsUpper(c) { // pregunto si la letra de palabra original esta en mayuscula
			nuevaPalabra.WriteRune(unicode.ToUpper(runaNueva[i])) // si lo está transformo la letra de la palabra nueva que esta en la misma posicion que la palabra original en mayuscula
			                                                      // trabajar con runas te permite que respete las letras que contienen tilde
		} else {
			nuevaPalabra.WriteRune(((runaNueva[i]))) // si no la tengo que poner en mayuscula, la escribo normal
		}
	}
	return nuevaPalabra.String() // devuelvo el string
}
