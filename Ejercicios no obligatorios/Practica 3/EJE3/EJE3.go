// respuestas teoricas https://justpaste.it/5rs13

package main
import (
	"fmt"
	//"time"
)

func main() {
done:= make(chan bool)
fmt.Println("Inicia Goroutine del main")
go hello(done)
<-done
//time.Sleep(time.Second) // primera solucion con time.sleep
fmt.Println("Termina Goroutine del main")
}

func hello(done chan bool) {
fmt.Println("Inicia Goroutine de hello")
for i := 0; i < 3; i++ {
fmt.Println(i, " Hello world")
}
fmt.Println("Termina Goroutine de hello")

done <- true // aviso que terminé
}

