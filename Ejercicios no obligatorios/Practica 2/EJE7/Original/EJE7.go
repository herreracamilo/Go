package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)
func main()  {
	var letras,numeros,especiales int
	read := bufio.NewReader(os.Stdin)

	fmt.Println("ingrese la secuencia de caracteres")
	input, _ := read.ReadString('\r')

	secuencia:= []rune(input[:len(input)-1])
	filtro(secuencia,&letras,&numeros,&especiales)
	fmt.Println(letras)
	fmt.Println(numeros)
	fmt.Println(especiales)
}

func filtro(s []rune,l,n,e *int)  {
	for _,r :=range s{
		if(unicode.IsLetter(r)){
			*l++
		}else if(unicode.IsNumber(r)){
			*n++
		}else {
			*e++
		}
	}
}