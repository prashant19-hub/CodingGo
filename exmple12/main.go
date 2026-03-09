package main

import "fmt"

func isLeapYear(year int) bool {
    return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

func main() {
    testYears := []int{2024, 2025, 2000, 1900, 2028}
    
    for _, year := range testYears {
        if isLeapYear(year) {
            fmt.Printf("%d is a leap year\n", year)
        } else {
            fmt.Printf("%d is not a leap year\n", year)
        }
    }
}
