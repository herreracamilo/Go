package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)
func main()  {
	var letrasMayus,letrasMinus,numeros,especiales int
	read := bufio.NewReader(os.Stdin)

	fmt.Println("ingrese la secuencia de caracteres")
	input, _ := read.ReadString('\r')

	secuencia:= []rune(input[:len(input)-1])
	filtro(secuencia,&letrasMayus,&letrasMinus,&numeros,&especiales)
	fmt.Println(letrasMayus)
	fmt.Println(letrasMinus)
	fmt.Println(numeros)
	fmt.Println(especiales)
}

func filtro(s []rune,lMayu,lMin,n,e *int)  {
	for _,r :=range s{
		if(unicode.IsUpper(r)){
			*lMayu++
		}else if(unicode.IsLower(r)){
			*lMin++
		}else if(unicode.IsNumber(r)){
			*n++
		}else {
			*e++
		}
	}
}