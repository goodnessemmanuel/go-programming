package datatypes

import "fmt"
//Long declaration
var y = 42;
//Decalared the variable with the Identifier "z" and is of type string
var z string = "I am a variable of type string";
//End long variable declaration

 func Datatypes() string {
    fmt.Println(y)
    fmt.Printf("%T\n", y) //print the variable type
    fmt.Println(z)
    fmt.Printf("%T\n", z)
    return z
}