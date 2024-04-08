package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "oneline" {
		fmt.Println(oneLine())

		return
	}

	// set up signal for closing
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-quit
		onExit()
		os.Exit(0)
	}()

	fmt.Print("\033[?47h") // save screen
	fmt.Print("\033[H")    // move to top of screen
	fmt.Print("\033[?25l") // make cursor invisible

	for {
		fmt.Print("   ʙɪɴᴀʀʏ ᴄʟᴏᴄᴋ\n\n")
		fmt.Print(getClock())
		time.Sleep(50 * time.Millisecond)
		fmt.Print("\033[10A") // moves the cursor 10 cells in upward
	}
}

func onExit() {
	fmt.Print("\033[?47l") // load screen
	fmt.Print("\033[?25h") // make cursor visible
	fmt.Print("\033[1 q")  // make cursor blinks again
}

func oneLine() string {
	str := ""
	digits := getDigits()

	str += digitToBinaryString(digits[0])
	str += " "
	str += digitToBinaryString(digits[1])
	str += " : "
	str += digitToBinaryString(digits[2])
	str += " "
	str += digitToBinaryString(digits[3])
	str += " : "
	str += digitToBinaryString(digits[4])
	str += " "
	str += digitToBinaryString(digits[5])

	return str
}

func digitToBinaryString(digit int) string {
	on := "1"
	off := "0"

	switch digit {
	case 0:
		return off + off + off + off
	case 1:
		return off + off + off + on
	case 2:
		return off + off + on + off
	case 3:
		return off + off + on + on
	case 4:
		return off + on + off + off
	case 5:
		return off + on + off + on
	case 6:
		return off + on + on + off
	case 7:
		return off + on + on + on
	case 8:
		return on + off + off + off
	case 9:
		return on + off + off + on
	}

	return ""
}

func getDigits() [6]int {
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

	return digits
}

func getClock() string {
	digits := getDigits()

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
