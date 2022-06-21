package main

import (
	"fmt"
	"log"
	"test/helper"
)

func main(){
	helper.Help()
	data, err := Greet("")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data)
}
