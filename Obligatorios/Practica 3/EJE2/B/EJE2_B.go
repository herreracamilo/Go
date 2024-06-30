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
	num int
	cola chan Cliente
}

func main() {
	rand.Seed(time.Now().UnixNano()) // random
	start := time.Now() // Inicio de la medición de tiempo
	numCajas := 3
	numClientes := 15

	cajas := make([]*Caja, numCajas)
	var wg sync.WaitGroup

	// creo las cajas con sus colas
	for i := 1; i <= numCajas; i++ {
		cajas[i-1] = &Caja{
			num: i,
			cola: make(chan Cliente, numClientes/numCajas+1), // me aseguro que la cola tenga lugar suficiente
		}

		// cada caja atiende a los clientes en su propia cola
		go func(c *Caja) {
			for cliente := range c.cola {
				c.atender(cliente, &wg)
			}
		}(cajas[i-1])
	}

	// asignar los clientes clientes a las cajas
	for i := 1; i <= numClientes; i++ {
		wg.Add(1)
		cajas[(i-1)%numCajas].cola <- Cliente{num: i} // Round-Robin
	}

	// cierro todas las colas
	for _, caja := range cajas {
		close(caja.cola)
	}

	wg.Wait()
	duration := time.Since(start) // Fin de la medición de tiempo
	fmt.Println("Todos los clientes han sido atendidos")
	fmt.Printf("Tiempo con 1 goroutines: %v\n", duration)
}

func (c *Caja) atender(cliente Cliente, wg *sync.WaitGroup) {
	defer wg.Done() // manda que terminó

	fmt.Printf("La caja %d está atendiendo al cliente %d\n", c.num, cliente.num)
	time.Sleep(time.Duration(rand.Intn(2)) * time.Second)
	fmt.Printf("La caja %d terminó de atender\n", c.num)
}
