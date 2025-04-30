package main

import (
	"hashtables/ht"
	"log"
)

func main() {
	ht := ht.NewHashTable(10)
	logger := log.Default()
	val, err := ht.Get(3)
	if err != nil {
		logger.Println(err.Error())
	}
	logger.Printf("value of key is: %s", val)
}
