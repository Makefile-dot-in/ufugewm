package main

import (
	"log"
)

func HandleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func InitError()  {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}
