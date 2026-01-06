package main

import (
	"fmt"
	"time"
)

func main() {
	limit := 5
	count := 0
	start := time.Now()

	for i := 1; i <= 10; i++ {
		if time.Since(start).Seconds() > 10 {
			count = 0
			start = time.Now()
		}

		if count >= limit {
			fmt.Println("Rate limit exceeded")
		} else {
			count++
			fmt.Println("Request accepted:", count)
		}
		time.Sleep(time.Second)
	}
}
