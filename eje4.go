/* escriba un programa que imprima en la salida estándar la suma
de los primeros números positivos pares menores o iguales a
250. Cambiar el programa para que itere en el sentido contrario
pero obtener el mismo resultado. Cambiar el programa para que
en lugar de usar un literal como tope se use una constante. Si lo
desea, investigue la herramienta gofmt y pruebe sobre el código
escrito.
Sub-objetivo: Uso de E/S de valores numéricos en Go,
estructuras de control básicas, constantes y variables.*/

package main

import (
	"fmt"
)

func main() {
	const max = 250
	fmt.Println(sumaAscendente(max))
	fmt.Println(sumaDescendente(max))

}

func sumaAscendente(max int) int {
	suma := 0
	for i := 0; i <= max; i += 2 { //con el +=2 sumo de a 2 para que solo sean pares
		suma += i
	}
	return suma
}

func sumaDescendente(max int) int {
	suma := 0
	for i := max; i >= 0; i -= 2 {
		suma += i
	}
	return suma
}
