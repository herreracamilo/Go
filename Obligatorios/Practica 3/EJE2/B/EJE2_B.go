package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Cliente struct {
	num int
}

type Caja struct {
	num  int
	cola chan Cliente
}

func main() {
	rand.Seed(time.Now().UnixNano())
	var wg sync.WaitGroup

	numCajas := 3
	numClientes := 10

	cajas := make([]*Caja, numCajas)

	// creo un canal para cada cola
	for i := 0; i < numCajas; i++ { 
		cajas[i] = &Caja{
			num:  i + 1,
			cola: make(chan Cliente, numClientes),
		}

		// Cada caja atiende clientes en su cola individual
		go func(c *Caja) {
			for cliente := range c.cola {
				c.atender(cliente, &wg)
			}
		}(cajas[i])
	}

	// le doy un cliente a las cajas forma round-robin
	for i := 0; i < numClientes; i++ {
		cliente := Cliente{num: i + 1}
		indiceCaja := i % numCajas
		wg.Add(1) // agregar wg antes de enviar al canal
		cajas[indiceCaja].cola <- cliente
	}

	// cierro las cajas
	for _, caja := range cajas {
		close(caja.cola)
	}

	wg.Wait()
	fmt.Println("Todos los clientes han sido atendidos")
}
func (c *Caja) atender(cliente Cliente, wg *sync.WaitGroup)  {
	defer wg.Done()// manda que terminó

	fmt.Printf("La caja %d está atendiendo al cliente %d\n", c.num, cliente.num)
	time.Sleep(time.Duration(rand.Intn(2))*time.Second)
	fmt.Printf("La caja %d terminó de atender\n", c.num)
}