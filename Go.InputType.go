package main

import (
    "fmt"
    
    
)

func printType(varType interface{}){
	switch varVal := varType.(type){
		case int:
		    fmt.Printf("%v doubled is %v\n", varVal, varVal*2)
		case string:
		    fmt.Printf("Length of string %q is %v bytes\n", varVal, len(varVal))
		default:
		    fmt.Printf("I'm not sure what type your variable is\n")
	}
	
}

func main() {
    printType("hello")
    printType(7)
    printType(7.8)
}