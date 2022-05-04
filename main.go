package main

import "fmt"

func main() {
	fmt.Println("Welcome to go!")
	fmt.Println(stringManipulate())
}

var testString string = `"this is a sample for testing"'`

func stringManipulate() string {
	maniPulateTest := []byte(testString) //string converted to byte sequence

	//print the UTF-8 code point for maniPulateTest
	for i := 0; i < len(maniPulateTest); i++ {
		fmt.Printf("%d utf code point -> %#U \n", maniPulateTest[i], maniPulateTest[i])
	}
	fmt.Println()
	return fmt.Sprintf("%v", maniPulateTest)
}