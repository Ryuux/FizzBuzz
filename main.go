package main

import "fmt"

const (
    Fizz = "Fizz"
    Buzz = "Buzz"
)

func main() {
    fmt.Printf("%-10s %-10s %s\n", "Number", "FizzBuzz", "Result")
    fmt.Println("-------------------------------------")

    count := 0
    for i := 1; i <= 100; i++ {
        fizz := i%3 == 0
        buzz := i%5 == 0

        var result string
        if fizz && buzz {
            result = Fizz + Buzz
        } else if fizz {
            result = Fizz
        } else if buzz {
            result = Buzz
        } else {
            result = fmt.Sprintf("%d", i)
        }

        fmt.Printf("%-10d %-10s %s\n", i, result, " ")
        count++
    }
}

