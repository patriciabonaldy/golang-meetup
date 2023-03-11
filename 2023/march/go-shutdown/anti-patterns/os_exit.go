package main

import (
	"fmt"
	"os"
	"time"
)

func anyMethod() {
	fmt.Println("Executing something")
	<-time.After(3 * time.Second)
	fmt.Println("Done")
}

func main() {
	var ch = make(chan int)
	go func() {
		<-ch
		os.Exit(1)
	}()
	go anyMethod()
	time.Sleep(1 * time.Second)
	ch <- 1
}
