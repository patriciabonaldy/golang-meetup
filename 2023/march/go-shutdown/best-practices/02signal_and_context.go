package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	done := make(chan struct{})
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		sigint := make(chan os.Signal, 1)

		signal.Notify(sigint, syscall.SIGINT, syscall.SIGTERM)
		<-sigint
		cancel()

		ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		ShutdownServer(ctx)
		done <- struct{}{}
	}()

	// starting a server

	// execute a goroutine
	go Run(ctx)

	<-done
}

func Run(ctx context.Context) {
	for {
		select {
		case <-time.After(1 * time.Second):
			fmt.Println("doing something...")
		case <-ctx.Done():
			fmt.Println("Break the loop")
			return
		}
	}
}

func ShutdownServer(_ context.Context) {
	fmt.Println("Stopping http server ...")
	<-time.After(1 * time.Second)
}
