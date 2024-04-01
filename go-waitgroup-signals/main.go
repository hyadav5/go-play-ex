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

var shutdownSignals = []os.Signal{os.Interrupt, syscall.SIGTERM}

// This call will setup a signal handler and return the context which will get cancelled on receiving an OS signals.
// Also, this call will wait in a go routine, to wait for all the other go routines which are tied to the received wait group.
// This call is perfect example of an application which is running multiple goroutines and want all of them to be gracefully
// terminated on receiving the OS signal.
func SetupSignalHandler(wg *sync.WaitGroup) context.Context {
	ctx, cancel := context.WithCancel(context.Background())

	c := make(chan os.Signal, 2)
	signal.Notify(c, shutdownSignals...)
	go func() {
		<-c
		println("Closing down all the Data Collectors")
		cancel()
		wg.Wait()
		fmt.Println("Shutdown complete. Closing Sys Health Service")
		<-c
		os.Exit(1) // second signal. Exit directly.
	}()

	return ctx
}

func myRoutineOne(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	c := make(chan int)
	comeOut := false
	for {
		select {
		case <-c:
			return
		case <-ctx.Done():
			fmt.Println("Shutting down myRoutineOne() gracefully...")
			comeOut = true
			break
		}
		if comeOut {
			break
		}
	}
}

func myRoutineTwo(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	c := make(chan int)
	comeOut := false
	for {
		select {
		case <-c:
			return
		case <-ctx.Done():
			fmt.Println("Shutting down myRoutineTwo() gracefully...")
			comeOut = true
			break
		}
		if comeOut {
			time.Sleep(5 * time.Second)
			break
		}
	}
}

func main() {
	// Create a WaitGroup to keep track of running goroutines
	var wg sync.WaitGroup
	ctx := SetupSignalHandler(&wg)

	wg.Add(1)
	go myRoutineOne(ctx, &wg)

	wg.Add(1)
	go myRoutineTwo(ctx, &wg)

	for {
	}
}
