package main

import (
	"fmt"
	"log"
	"container/list"
)

func main(){
	var n int
	read(&n)

	canal := make(chan *list.List) //creo el canal

	go addPrimos(n, canal)

	primos := <- canal  // aca declaro y recibo la lista a partir del canal

	printPrimos(primos)

}

func read(n *int) {
    _,err := fmt.Scanln(n)
    if err != nil {
        log.Fatal(err)
    }
}

func esPrimo(n int) bool{
	if(n<=1){
		return false
	}
	for i:= 2 ; i*i <= n; i++{
		if(n % i == 0){
			return false
		}
	}
	return true
}

func addPrimos(n int, c chan *list.List){
	primos := list.New() // creo una nueva lista doblemente enlazada
	for i:= 2; i<=n;i++{
		if(esPrimo(i)){
			primos.PushBack(i)
		}

	}
	c <- primos // mando los primos por el canal
	close(c)
}

func printPrimos(primos *list.List) {
	for e := primos.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, " ")
	}
	fmt.Println()
}
