package main

func KeepProcessAlive() {
	var ch chan int
	<-ch
}

func main() {
	// execute some things
	KeepProcessAlive()
}
