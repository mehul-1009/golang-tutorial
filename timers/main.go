package main

import (
	"fmt"
	"time"
)

func main() {

	timer1 := time.NewTimer(2 * time.Second)

	<-timer1.C
	fmt.Println("Timer 1 fired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	time.Sleep(2 * time.Second)
}

// | Part      | Description                             | Output              |
// | --------- | --------------------------------------- | ------------------- |
// | Timer 1   | Waits 2 seconds, then fires             | ✅ `Timer 1 fired`   |
// | Timer 2   | Created for 1 second, but stopped early | ✅ `Timer 2 stopped` |
// | Goroutine | Waits for Timer 2, but never unblocks   | ⛔ No output         |
