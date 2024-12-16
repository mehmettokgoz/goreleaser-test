package main

import (
	"fmt"
	"main/cmd"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	fmt.Printf("Starting development server...")
	service := cmd.NewService()
	service.Start()
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)
	<-exit
}
