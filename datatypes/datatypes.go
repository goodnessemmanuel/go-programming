package datatypes

import "fmt"

//Long declaration
var y = 42 //package scope decalared variable. similar to var z and n below
//Decalared the variable with the Identifier "z" and is of type string
// both z and n are referred to as string literrals
var z string = "I am a variable of type string"
var n string = `"I am also a variable of type string"`

var status bool = true


//End long variable declaration

func Datatypes() {
	fmt.Println(y)
	fmt.Printf("%T\n", y) //print the variable type
	fmt.Println(z)
	fmt.Printf("%T\n", z)
	fmt.Println(n)
	fmt.Printf("%T\n", n)
    fmt.Println(status)
	fmt.Printf("%T\n", status)
}
