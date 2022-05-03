package datatypes

import (
	"fmt"
	"runtime"
)

//Long declaration
var y = 42 //package scope decalared variable. similar to var z and n below
//Decalared the variable with the Identifier "z" and is of type string
// both z and n are referred to as string literrals
var z string = "I am a variable of type string"
var n string = `"I am also a variable of type string"`

var status bool = true

var zeroValueString string
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
//working with own type
type hotdog int
var derivativeOfHotDog hotdog
var num int
func OwnType()  {
    fmt.Println(derivativeOfHotDog)
    fmt.Printf("%T\n", derivativeOfHotDog)
    //you cannot do > num = derivativeOfHotDog because both are of different datatype
    num = int(derivativeOfHotDog)
    fmt.Println(num)

}

type myowntype int
var x myowntype
func Exercise()  {
    fmt.Println(x) //you get the zero value here since x is not assigned
    fmt.Printf("%T\n", x) // print type of x
    x = 42
    fmt.Println(x)
    fmt.Printf("%T\n", x)
    k := int(x)
    fmt.Println(k)
    fmt.Printf("%T\n" , k)
    
    fmt.Println("System spec inbuilt go command")
    fmt.Println(runtime.GOOS)
    fmt.Println(runtime.GOARCH)
}