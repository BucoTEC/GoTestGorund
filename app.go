package main

import (
	"fmt"
	"test/helper"
)

func main(){
	helper.Help()
	data := Greet("adnan")
	fmt.Println(data)
}
