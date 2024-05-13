package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)
func main()  {
	var letrasMayus,letrasMinus,especiales int
	var numerosMap = map[int]int{
		0: 0,
		1: 0,
		2: 0,
		3: 0,
		4: 0,
		5: 0,
		6: 0,
		7: 0,
		8: 0,
		9: 0,
	}
	read := bufio.NewReader(os.Stdin)

	fmt.Println("ingrese la secuencia de caracteres")
	input, _ := read.ReadString('\r')

	secuencia:= []rune(input[:len(input)-1])
	filtro(secuencia,&letrasMayus,&letrasMinus,&especiales,numerosMap)
	fmt.Println("mayusculas")
	fmt.Println(letrasMayus)
	fmt.Println("minusculas")
	fmt.Println(letrasMinus)
	fmt.Println("map de numeros")
	fmt.Println(numerosMap)
	fmt.Println("especiales")
	fmt.Println(especiales)
}


	func filtro(s []rune,lMayu,lMin,e *int, numerosMap map[int]int)  {
		for _,r :=range s{
			num := int(r - '0')
			if(unicode.IsUpper(r)){
				*lMayu++
			}else if(unicode.IsLower(r)){
				*lMin++
			}else if(unicode.IsNumber(r)){
				numerosMap[num]++
			}else {
				*e++
			}
		}
}

/*
Cómo funciona el num:= int(r-'0') 
'0' tiene un valor ASCII de 48 entonces cuando recibe el ASCII del numero 5 por ejemplo
que tiene un vlaor de 53 hace 53 (del cinco) - 48 (del cero) = 5 entonces ya me da el 5 para buscarlo en el map
*/