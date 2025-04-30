package main

import (
	"hashtables/ht"
	"log"
)

func main() {
	ht := ht.NewHashTable(10)
	logger := log.Default()

	ht.Set(123, "Hello")
	val, err := ht.Get(123)
	if err != nil {
		logger.Println(err.Error())
	}
	ht.Set(123, "World")
	val2, err := ht.Get(33)
	if err != nil {
		logger.Println(err.Error())
		return
	}

	logger.Printf("value of key was: %s", val)
	logger.Printf("value of key is: %s", val2)
}
