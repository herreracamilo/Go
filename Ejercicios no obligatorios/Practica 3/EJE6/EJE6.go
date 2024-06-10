package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main()  {
	// inicio el random
	rand.Seed(time.Now().UnixNano())

	// creo el grupo de espera
	var wg sync.WaitGroup

	//creo los 3 canales
	ch1:= make(chan int)
	ch2:= make(chan int)
	ch3:= make(chan int)

	wg.Add(3)
	go sendValue("ch1",ch1,&wg)
	go sendValue("ch2",ch2,&wg)
	go sendValue("ch3",ch3,&wg)

	mapCount:= map[int]int{1: 0,2: 0,3:0}
	
	for i := 0; i < 15; i++ {
		select {
		case v, ok := <-ch1:
			if ok {
				fmt.Printf("Recibido del canal 1: %d\n", v)
				mapCount[1]++
			}
		case v, ok := <-ch2:
			if ok {
				fmt.Printf("Recibido del canal 2: %d\n", v)
				mapCount[2]++
			}
		case v, ok := <-ch3:
			if ok {
				fmt.Printf("Recibido del canal 3: %d\n", v)
				mapCount[3]++
			}
		}
	}
	wg.Wait()

	fmt.Println("Totales recibidos de cada canal:")
	fmt.Println(mapCount)
}

func sendValue(name string,ch chan int, wg *sync.WaitGroup){
	defer wg.Done()
	//creo un slices de 5 numeros del 1 al 100
	values := make([]int, 5)
	for i := 0; i < len(values); i++ {
		values[i] = rand.Intn(100) + 1 // generar un número aleatorio del 1 al 100
	}
	//recorro el slice y mando cada numero al canal
	for _,v:= range values {
		fmt.Printf("Enviando %d al canal %s\n", v, name)
		ch <- v
	}
}