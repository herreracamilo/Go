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
	rand.Seed(time.Now().UnixNano())

	numCajas := 3
	numClientes := 15

	cajas := make([]*Caja, numCajas)
	var wg sync.WaitGroup
	var mutex sync.Mutex

	// crea las cajas con sus colas
	for i := 1; i <= numCajas; i++ {
		cajas[i-1] = &Caja{
			num:  i,
			cola: make(chan Cliente, numClientes), 
		}

		// cada caja atiende a los clientes en su propia cola
		go func(c *Caja) {
			for cliente := range c.cola {
				c.atender(cliente, &wg)
			}
		}(cajas[i-1])
	}

	// asigna los clientes a las cajas con la cola mas corta
	for i := 1; i <= numClientes; i++ {
		wg.Add(1)
		cliente := Cliente{num: i}

		mutex.Lock() //lockeo y busco la cola con menor cola
		indiceMin := 0
		longitudMin := len(cajas[0].cola)
		for j := 1; j < numCajas; j++ {
			if len(cajas[j].cola) < longitudMin {
				indiceMin = j
				longitudMin = len(cajas[j].cola)
			}
		}
		mutex.Unlock() // desbloqueo y mando a la caja con menos cola al cliente

		cajas[indiceMin].cola <- cliente
	}

	// cierro todas las colas despues de asignar los clientes
	for _, caja := range cajas {
		close(caja.cola)
	}

	wg.Wait()
	fmt.Println("Todos los clientes han sido atendidos")
}

func (c *Caja) atender(cliente Cliente, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("La caja %d está atendiendo al cliente %d\n", c.num, cliente.num)
	time.Sleep(time.Duration(rand.Intn(2)) * time.Second)
	fmt.Printf("La caja %d terminó de atender\n", c.num)
}
