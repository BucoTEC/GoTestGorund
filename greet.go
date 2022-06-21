package main

import (
	"errors"
	"fmt"
)

func Greet(name string) (string, error){
	if name == ""{
		return "",errors.New("missing name param")
	}
	hy := fmt.Sprintf("heeloooo %v.", name)
	return hy,nil
	
}

