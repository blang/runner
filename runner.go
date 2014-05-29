package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	log.Printf("Started Runner: %s", os.Args)
	f, err := os.OpenFile("./runner.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0660)
	if err != nil {
		log.Printf("Could not write log file, %s", err)
	} else {
		defer f.Close()
		f.WriteString(fmt.Sprintf(time.Now().String()+": Started Runner: %s\n", os.Args))
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGINT, syscall.SIGTERM)
	<-c
	log.Println("Stopped Runner")
}
