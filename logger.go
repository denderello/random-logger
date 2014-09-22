package main

import (
	"log"
	"math/rand"
	"time"
)

func main() {
	for {
		log.Printf("Random log entry %v", rand.Int())
		time.Sleep(time.Second)
	}
}
