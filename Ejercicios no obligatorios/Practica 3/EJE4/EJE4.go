package main
import (
"fmt"
)

func pxng(m chan string, str string, done chan bool) {
	m <- str
	done <-true
}
func main() {
	messages := make(chan string)
	done:= make(chan bool)
	for i := 0; i < 5; i++ {
		go pxng(messages, "PING",done)
		go pxng(messages, "PONG",done)
	}
	
	for i := 0; i < 10; i++ {
		fmt.Println(<-messages)
		<-done 
	}
}	