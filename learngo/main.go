package main

import (
	"fmt"
	"unsafe" // use for Sizeof
)

func main() {
	fmt.Println("Hello World!!")
	fmt.Println("The following is a little info about myself")
	// ---------------------declaring var-----------------------
	const name = "Cindy Fu"
	// constant value can not be assigned by function because it should be known at complie time not proggram running time
	// const name has no type, but it have a default string type if want using %T. Since it has no type, we can aissgn them into var in diff types
	var year int //initialize year to 0 value in the datatype
	year = 20
	var food, interest = "hotpot", "skiing" // drop type if variable is assigned; this format only store var with same type
	var (                                   // use to declar different type var
		lucky_num            = 6
		message              = "go!!!"
		random_float float64 = 3.1434
	)
	// short hand declaration
	class, year := "CSE 311, CSE 351, Dance109", 18 //用这种format： all var needs to be assigned

	fmt.Println("My name is ", name, ".I am ", year, " years old. My fav food is ", food, ". I love doing ", interest, ". My lucky number is ", lucky_num, ". This is a random message and float ", message, random_float, ". I am taking ", class, "this year.")
	// Pakage.function() is the format
	// Go is strict type!!!! can't change the type of a variable
	// ------------------------type conversion-------------------
	compute := lucky_num + int(random_float)
	var compute2 int = 15 / 15.0
	fmt.Println("Converting random_float into int to do computation", compute, compute2)
	// ------------------------format specifier---------------
	fmt.Printf("Type of year variable %T, size of year variable %d", year, unsafe.Sizeof(year))
	//-------------------------func-------------
	var next, nextnext int = calc_fib(0, 1)
	fmt.Println("the next two num in fibbonacchhi after 0 and 1 are", next, nextnext)
}

func calc_fib(a, b int) (next, nextnext int) { // if a returned value is named, it is the same as declaring var at 1st line of func
	next = a + b
	nextnext = b + next
	return // don't need to explicit say when return values are named
}

//-------------------different ways to compile and run module--------------------
// compile: go install will compile this program to go/bin/
// execute: ~/go/bin/learngo to execute the program
// export PATH=$PATH:~/go/bin use this to avoid typing the path each time but just typing learngo
// !! allow us to run program from any direc in the terminal
// go install will find the go.mod in the current directory(or recursively trace its parent directory)and find the package

// compile: go build to complie program in the current location
// execute: ./learngo to run program

// compile + execute: go run main.go
// compile the file to a temporary location
