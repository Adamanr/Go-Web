package main

import (
	Server "GoWeb/server"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	Server.Start()
}

func processingSignals() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	exitChan := make(chan int)
	go func() {
		for {
			sig := <-sigs
			switch sig {
			case syscall.SIGHUP:
				fmt.Println("Signal hang up triggered")
			case syscall.SIGINT:
				fmt.Println("Signal interrupt triggered")
			case syscall.SIGTERM:
				fmt.Println("Signal terminate triggered")
				exitChan <- 0
			case syscall.SIGQUIT:
				fmt.Println("Signal quit triggered")
				exitChan <- 0
			default:
				fmt.Println("Unknown signal triggered")
				exitChan <- 1
			}
		}
	}()
	exitCode := <-exitChan
	os.Exit(exitCode)
}
