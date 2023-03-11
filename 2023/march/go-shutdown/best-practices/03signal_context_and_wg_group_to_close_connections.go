package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	done := make(chan struct{})
	ctx, cancel := context.WithCancel(context.Background())

	// creating DB connections

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, syscall.SIGINT, syscall.SIGTERM)
		<-sigint
		cancel()

		wg := sync.WaitGroup{}
		wg.Add(1)

		go func() {
			defer wg.Done()
			closeDB()
		}()

		ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		shutdownServer(ctx)
		wg.Wait()

		done <- struct{}{}
	}()
	go run(ctx)

	// starting a server

	<-done
}

func run(ctx context.Context) {
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

func closeDB() {
	fmt.Println("close DB connection ...")
}

func shutdownServer(_ context.Context) {
	fmt.Println("Stopping http server ...")
}
