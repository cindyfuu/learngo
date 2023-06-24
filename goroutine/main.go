package main

import (
	"fmt"
	"time"
)

func numbers() {  
    for i := 1; i <= 5; i++ {
        time.Sleep(250 * time.Millisecond)
        fmt.Printf("%d ", i)
    }
}
func alphabets() {  
    for i := 'a'; i <= 'e'; i++ {
        time.Sleep(400 * time.Millisecond)
        fmt.Printf("%c ", i)
    }
}

func calcSquares(number int, squareop chan int) {  
    sum := 0
    for number != 0 {
        digit := number % 10
        sum += digit * digit
        number /= 10
    }
    squareop <- sum
}

func calcCubes(number int, cubeop chan int) {  
    sum := 0 
    for number != 0 {
        digit := number % 10
        sum += digit * digit * digit
        number /= 10
    }
    cubeop <- sum
} 

func main() {  
    fmt.Println("part 1: goroutine")
    go numbers()
    go alphabets()
    time.Sleep(3000 * time.Millisecond) //time.Sleep(1 * time.Second)
    fmt.Println("main terminated")

    fmt.Println("part 2: channel")
    var a chan int 
    // assign channel to nil but it can not be used until properly intialized 
    if a == nil {
        fmt.Println("channel a is nil, going to define it")
        a = make(chan int) // initialize empty data strucutre using make for map, slice, channel
        fmt.Printf("Type of a is %T", a)
    }

    fmt.Println("part 3: channel blocks the read and writes in default")
    //read in main is blocked untill data is write in the goroutine 
    // read from channel a : data := <- a  
    //write in main is blocked until data is read in the goroutine 
    // write from channel a :a <- data
    number := 589
    sqrch := make(chan int)
    cubech := make(chan int)
    go calcSquares(number, sqrch)
    go calcCubes(number, cubech)
    squares, cubes := <-sqrch, <-cubech
    fmt.Println("Final output", squares + cubes)
    // send-only channel: chan<- TYPE
    // receive-only channel: <-chan TYPE


    fmt.Println("part 4: buffer channel")
    // send multiple value to a channel without blockin g
    ch := make(chan string, 2)
    ch <- "naveen"
    ch <- "paul"
    fmt.Println(<- ch)
    fmt.Println(<- ch)
    //very interesting example
    ch := make(chan int, 2)
    go write(ch)
    time.Sleep(2 * time.Second)
    for v := range ch {
        fmt.Println("read value", v,"from ch")
        time.Sleep(2 * time.Second)
    }
    //use range 
    ch := make(chan int, 5)
    ch <- 5
    ch <- 6
    close(ch)
    for n := range ch {
        fmt.Println("Received:", n)
    }
    

}

func write(ch chan int) {  
    for i := 0; i < 5; i++ {
        ch <- i
        fmt.Println("successfully wrote", i, "to ch")
    }
    close(ch)
}