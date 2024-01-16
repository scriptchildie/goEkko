package main

import (
	"fmt"
	"time"

	"golang.org/x/sys/windows"
)

func myGoroutine() {

	for {
		time.Sleep(1 * time.Second)
		fmt.Println("Hello9")
	}
}

func main() {

	go myGoroutine()
	go myGoroutine()
	go myGoroutine()
	go myGoroutine()

	fmt.Println(windows.GetCurrentThreadId())
	for {
		fmt.Println("\n[----]Calling ekko[----]")

		err := ekkoSleep(10000)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("[----]Exited ekko[----]\n\n")

	}
}
