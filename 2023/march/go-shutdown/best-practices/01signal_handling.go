package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, syscall.SIGINT, syscall.SIGTERM)
	for {
		select {
		case sig := <-sigint:
			fmt.Printf("Shutdown request (signal: %v)\n", sig)
			fmt.Println("Closing connections")
			return
		case <-time.After(1 * time.Second):
			fmt.Println("Executing something")
		}
	}
}
