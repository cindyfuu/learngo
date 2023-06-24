package main

// init in 顺序: imported packages -> pacckage level variable -> init func in main -> main func
//------------import-----------
import (
	"fmt"
	"learnpackage/simpleinterest"

	// _"learnpackage/simpleinterest" use the even silencer to run init but not use the package
	"log"
	// can't import a package that is not used but can use var _ = package.Function() to temporarily avoid this
)

// -------------package level----------
var p, r, t = 5000.0, 10.0, 1.0

/*
* init function to check if p, r and t are greater than zero
 */
func init() {
	println("Main package initialized")
	if p < 0 {
		log.Fatal("Principal is less than zero") //log.Fatal terminate program execution
	}
	if r < 0 {
		log.Fatal("Rate of interest is less than zero")
	}
	if t < 0 {
		log.Fatal("Duration is less than zero")
	}
}

func main() {
	fmt.Println("Simple interest calculation")
	si := simpleinterest.Calculate(p, r, t)
	// any captiized function or variable are exported in go
	// has to be capitalized to access a function outside of current package
	fmt.Println("Simple interest is", si)

	slice := []string{"hello", "world"}
	test(slice...) // slice is directly passed into variadic function test without parse into another slice type
	fmt.Println(slice)
}

// ------------------learn variadic function---------------
func test(s ...string) {
	//1. accept any number of arguments type string
	//2. convert the variable number of arguements to the slice of the type of variadic parameter
	//3. calling test([]string{"hello", "world"}) would fail since test is a variadic func that accepts a variable number of strings
	// it will convert the number of strings passed in as slice of type string. []string{"hello", "world"} will
	// be turned to []string{[]string{"hello", "world"}} but []string{"hello", "world"} is not an string but a slice
	s[0] = "go"
	s = append(s, "playground")
	fmt.Println(s)

}
