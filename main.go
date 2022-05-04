package main

import "fmt"

func main() {
	fmt.Println("Welcome to go!")
	//fmt.Println(stringManipulate())
	demonstrateBitshifting()
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

func demonstrateBitshifting() {
	x := 8
	fmt.Println("Before shifting bit, decimal and its binary equivalent is:")
	fmt.Printf("%d(decimal)  => %b (binary) \n", x, x) //%b prints the equivalent binary digit of the given variable
	fmt.Println("After shifting bit, decimal and its binary equivalent is:")
	y := x << 1 //left shifting
	fmt.Printf("%d(decimal)  => %b (binary) \n", y, y)
}
