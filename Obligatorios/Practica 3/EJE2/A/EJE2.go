package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Cliente struct{
	num int
}

type Caja struct{
	num int
}

func main()  {
	rand.Seed(time.Now().UnixNano()) // random
	start := time.Now() // Inicio de la medición de tiempo
	numCajas:= 3;
	numClientes:=10;

	cola := make(chan Cliente,numClientes)

	var wg sync.WaitGroup

	for i := 1; i <= numCajas; i++ {
		caja:= &Caja{num: i}

		go func(c *Caja)  {
			for cliente:= range cola{
				
				c.atender(cliente,&wg)
			}
		}(caja)
	}

	for i := 1; i <= numClientes; i++ {
		wg.Add(1)
		cola <- Cliente{num: i}
	}
	close(cola)
	wg.Wait()
	duration := time.Since(start) // Fin de la medición de tiempo
	fmt.Println("Todos los clientes han sido atendidos")
	fmt.Printf("Tiempo con 1 goroutines: %v\n", duration)
}

func (c *Caja) atender(cliente Cliente, wg *sync.WaitGroup)  {
	defer wg.Done()// manda que terminó

	fmt.Printf("La caja %d está atendiendo al cliente %d\n", c.num, cliente.num)
	time.Sleep(time.Duration(rand.Intn(2))*time.Second)
	fmt.Printf("La caja %d terminó de atender\n", c.num)
}