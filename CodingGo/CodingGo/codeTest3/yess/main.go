package main

import (
    "fmt"
    "time"
)

func main() {
    now := time.Now()
    printCalendar(now.Year(), int(now.Month()))
}

func printCalendar(year, month int) {
    t := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
    daysInMonth := daysIn(t.Year(), t.Month())
    
    fmt.Printf("   %s %d\n", t.Month().String()[:3], t.Year())
    fmt.Println("Su Mo Tu We Th Fr Sa")
    
    // Start padding
    for i := 0; i < int(t.Weekday()); i++ {
        fmt.Print("   ")
    }
    
    // Days
    for day := 1; day <= daysInMonth; day++ {
        fmt.Printf("%2d ", day)
        if (day+int(t.Weekday()))%7 == 0 {
            fmt.Println()
        }
    }
    fmt.Println()
}

func daysIn(year int, month time.Month) int {
    return time.Date(year, month+1, 0, 0, 0, 0, 0, time.UTC).Day()
}
