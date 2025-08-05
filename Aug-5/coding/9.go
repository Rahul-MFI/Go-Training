package coding

import (
	"fmt"
)

func getWeekday(day int) string {
	switch day {
	case 1:
		return "Monday"
	case 2:
		return "Tuesday"
	case 3:
		return "Wednesday"
	case 4:
		return "Thursday"
	case 5:
		return "Friday"
	case 6:
		return "Saturday"
	case 7:
		return "Sunday"
	default:
		return "Invalid day"
	}
}

func Nine() {
	dayNumber := 1
	fmt.Printf("Day %d is %s\n", dayNumber, getWeekday(dayNumber))
}
