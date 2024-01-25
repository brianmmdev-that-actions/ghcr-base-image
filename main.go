package main

import (
	"log"
	"os"
)

func main() {
	args := os.Args
	if len(args) > 1 {
		log.Printf("[FROM GHCR] Hello %v!!", args[1])
		return
	}
	log.Println("[FROM GHCR] Hello world!!")
}
