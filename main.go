package main

import (
	"fmt"

	datatypes "github.com/goodnessemmanuel/go-programming/datatypes"
)

func main() {
	datatype := datatypes.Datatypes()
	fmt.Println("Welcome to go!")
	fmt.Println(datatypes.Datatypes())//direct call
	fmt.Println(datatype) //Call via a short variable declaration 
}
