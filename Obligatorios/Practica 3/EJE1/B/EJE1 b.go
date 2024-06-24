package main

import (
	"fmt"
	"log"
	"container/list"
	"runtime"
	"sync"
	"time"
)

func main() {
	var n int
	fmt.Print("Enter the value of N: ")
	read(&n)

	numGoroutines := runtime.NumCPU() // Usar el número de CPUs disponibles
	canal := make(chan *list.List, numGoroutines) // Canal con buffer para los resultados
	var wg sync.WaitGroup

	start := time.Now() // Inicio de la medición de tiempo

	// Dividir el rango de números entre las goroutines
	rangeSize := n / numGoroutines
	for i := 0; i < numGoroutines; i++ {
		start := i*rangeSize + 1 // Inicio del rango
		end := (i + 1) * rangeSize // Fin del rango
		if i == numGoroutines-1 { // Ajuste para la última goroutine
			end = n
		}
		wg.Add(1)
		go addPrimos(start, end, canal, &wg)
	}

	/*
				^
				|
	Ejemplo n = 100 y numGoroutines = 4:
    Rango 1: start = 1, end = 25
    Rango 2: start = 26, end = 50
    Rango 3: start = 51, end = 75
    Rango 4: start = 76, end = 100
	*/

	// Esperar a que todas las goroutines terminen
	wg.Wait()
	close(canal) // Cerrar el canal después de que todas las goroutines hayan terminado

	// Recopilar y combinar resultados
	primos := list.New()
	for l := range canal {
		primos.PushBackList(l)
	}

	duration := time.Since(start) // Fin de la medición de tiempo

	printPrimos(primos) // Imprimir todos los números primos encontrados
	fmt.Printf("Tiempo con %d goroutines: %v\n", numGoroutines, duration)
}

func read(n *int) {
	_, err := fmt.Scanln(n)
	if err != nil {
		log.Fatal(err)
	}
}

func esPrimo(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func addPrimos(start, end int, c chan *list.List, wg *sync.WaitGroup) {
	defer wg.Done()
	primos := list.New() // Crear una nueva lista doblemente enlazada
	for i := start; i <= end; i++ {
		if esPrimo(i) {
			primos.PushBack(i)
		}
	}
	c <- primos // Enviar los primos por el canal
}

func printPrimos(primos *list.List) {
	for e := primos.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, " ")
	}
	fmt.Println()
}
