package main

import (
	"fmt"
	"test/helper"
)

func Greet(name string) string{
	helper.Help()
	hy := fmt.Sprintf("heeloooo %v.", name)
	return hy
	
}

