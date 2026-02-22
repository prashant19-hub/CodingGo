package main
import "fmt"

func main() {
	daysofweek := 1
	switch daysofweek {
	case 1:
		{
			fmt.Println("Monday")
		}
	case 2:
		{
		fmt.Println("Tuesday")
		}
	case 3:
		{
		fmt.Println("Wednesday")
		}
	case 4:
		{
		fmt.Println("Thursday")
		}
	case 5:
		{
		fmt.Println("Friday")
	}
	case 6:
		{
		fmt.Println("Saturday")
		}
	case 7:
		{
		fmt.Println("Sunday")
		}
	case 31:
		{
		fmt.Println("Invalid day of month")
		}
	default:
		fmt.Println("Invalid day of week")
	}
	clgbranch := "CSE"
	switch clgbranch {
	case "CSE":	
		fmt.Println("Computer Science and Engineering")
	case "ECE":	
		fmt.Println("Electronics and Communication Engineering")
	case "ME":	
		fmt.Println("Mechanical Engineering")
	default:
		fmt.Println("Invalid branch")
	}
}