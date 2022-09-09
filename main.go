package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Print("\033[?47h") // save screen
	fmt.Print("\033[H")    // move to top of screen
	fmt.Print("\033[?25l") // make cursor invisible

	for i := 0; i < 500; i++ {
		fmt.Print("   ʙɪɴᴀʀʏ ᴄʟᴏᴄᴋ\n\n")
		fmt.Print(getClock())
		time.Sleep(50 * time.Millisecond)
		fmt.Print("\033[10A")
	}

	fmt.Print("\033[?25h") // make cursor visible
	fmt.Print("\033[?47l") // save screen
}

func getClock() string {
	var digits [6]int

	now := time.Now()

	seconds := now.Second()
	minutes := now.Minute()
	hours := now.Hour()

	// hour
	digits[0] = hours / 10
	digits[1] = hours % 10

	// minute
	digits[2] = minutes / 10
	digits[3] = minutes % 10

	// second
	digits[4] = seconds / 10
	digits[5] = seconds % 10

	on := "🔵"
	off := "⚪"

	clock := ""

	// first row
	for i := 0; i < len(digits); i++ {
		if i == 0 || i == 2 || i == 4 {
			clock += "   "
			continue
		}

		if digits[i] == 8 || digits[i] == 9 {
			clock += on
		} else {
			clock += off
		}

		clock += " "
	}

	clock += "\n\n"

	// second row
	for i := 0; i < len(digits); i++ {
		if i == 0 {
			clock += "   "
			continue
		}

		if digits[i] == 4 || digits[i] == 5 || digits[i] == 6 || digits[i] == 7 {
			clock += on
		} else {
			clock += off
		}

		clock += " "
	}

	clock += "\n\n"

	// third row
	for i := 0; i < len(digits); i++ {
		if digits[i] == 2 || digits[i] == 3 || digits[i] == 6 || digits[i] == 7 {
			clock += on
		} else {
			clock += off
		}

		clock += " "
	}

	clock += "\n\n"

	// fourth row
	for i := 0; i < len(digits); i++ {
		if digits[i] == 1 || digits[i] == 3 || digits[i] == 5 || digits[i] == 7 || digits[i] == 9 {
			clock += on
		} else {
			clock += off
		}

		clock += " "
	}

	clock += "\n\n"

	return clock
}
