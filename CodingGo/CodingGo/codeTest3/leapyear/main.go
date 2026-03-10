package main

import "fmt"

func main() {
    var year int
    fmt.Print("Enter a year: ")
    fmt.Scanln(&year)

    if isLeapYear(year) {
        fmt.Printf("%d is a leap year.\n", year)
    } else {
        fmt.Printf("%d is not a leap year.\n", year)
    }
}

func isLeapYear(year int) bool {
    return year%400 == 0 || (year%4 == 0 && year%100 != 0)
}
