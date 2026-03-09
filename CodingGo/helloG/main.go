package main
mport "fmt"

type Period struct {
    Time   string
    Class  string
    Teacher string
}

type Schedule map[string][]Period // Day -> list of periods

func main() {
    sch := Schedule{
        "Monday": {
            {"9:00-10:00", "Math", "Mr. Rao"},
            {"10:00-11:00", "Science", "Ms. Patel"},
        },
        "Tuesday": {
            {"9:00-10:00", "English", "Mrs. Singh"},
            {"10:00-11:00", "History", "Mr. Khan"},
        },
        // Add more days...
    }

    displaySchedule(sch)
    queryPeriod(sch, "Monday", "Math")
}

func displaySchedule(sch Schedule) {
    fmt.Println("Weekly School Schedule:")
    for day, periods := range sch {
        fmt.Printf("\n%s:\n", day)
        for _, p := range periods {
            fmt.Printf("  %s: %s (%s)\n", p.Time, p.Class, p.Teacher)
        }
    }
}

func queryPeriod(sch Schedule, day, class string) {
    if periods, ok := sch[day]; ok {
        for _, p := range periods {
            if p.Class == class {
                fmt.Printf("%s %s is at %s with %s\n", day, class, p.Time, p.Teacher)
                return
            }
        }
    }
    fmt.Printf("%s has no %s\n", day, class)
}
