package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
	"os/exec"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <x>")
		fmt.Println("<x> should be a number between 0 and 9.")
		os.Exit(1)
	}

	x, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Invalid argument. <x> should be a number between 0 and 9.")
		os.Exit(1)
	}

	if x < 0 || x > 9 {
		fmt.Println("Invalid argument. <x> should be a number between 0 and 9.")
		os.Exit(1)
	}

	fmt.Println("Monitor app has been started!")

	ticker := time.NewTicker(10 * time.Second)
	for range ticker.C {
		if time.Now().Minute()%10 == x {
			cmd := exec.Command("osascript", "-e", `display notification "Pay Attention!" with title "Reminder"`)
			err := cmd.Run()
			if err != nil {
				fmt.Println(err)
			}

			time.Sleep(600 * time.Second)
		}
	}
}
