package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
	//"sync"
)

func productor(number chan<-int, wg *sync.WaitGroup)  {
	defer wg.Done()
	
	for i := 0; i < 3; i++ {
		num:= rand.Intn(101)
		time.Sleep(time.Duration(rand.Intn(1000))* time.Microsecond)
		number<- num
	}
		

}

func consumidor(id int,number <-chan int,wg *sync.WaitGroup)  {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		num:= <-number
		fmt.Printf("El consumidor %d consumio: %d\n",id,num)
	}
}

func main()  {
	// inicio el random
	rand.Seed(time.Now().UnixNano())

	// creo el grupo de espera
	var wg sync.WaitGroup

	// creo el canal para los 6 nueros
	numbers:= make (chan int)

	// creo el for para los productores
	for i := 1; i <= 2; i++ {
		wg.Add(1)
		go productor(numbers,&wg)
	}

	// creo el for para los consumidores
	for i := 1; i <= 2; i++ {
		wg.Add(1)
		go consumidor(i,numbers,&wg)
	}

	wg.Wait()
}