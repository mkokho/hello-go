package main

import (
	"log"
	"time"
	"fmt"
)

type Box struct {
	A string
	X int
}

func main() {
	expected()
	time.Sleep(200 * time.Millisecond)
	fmt.Println("-----------")
	weird()

	time.Sleep(1 * time.Second)
}

func expected() {
	events := []*Box{
		{"1", 10},
		{"2", 20},
		{"3", 99},
	}

	ch := make(chan *Box, 10)
	go func() {
		for x := range ch {
			log.Printf("In : %v", x)
		}
	}()

	for _, e := range events {
		x := e
		log.Printf("Out: %v", x)
		ch <- x
	}
	close(ch)
}

func weird() {
	events := []Box{
		{"1", 10},
		{"2", 20},
		{"3", 99},
	}

	ch := make(chan *Box, 10)
	go func() {
		for x := range ch {
			log.Printf("In : %v", x)
		}
	}()

	for i, _ := range events {
		x := &events[i]
		log.Printf("Out: %v", x)
		ch <- x
	}
}
