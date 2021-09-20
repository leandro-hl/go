package shared

import (
	"fmt"
	"time"
)

func Goroutines() {
	announce := func(message string, delay time.Duration) {
		//go routine calling a function literal.
		go func() {
			time.Sleep(delay)
			fmt.Println(message)
		}()
	}

	announce("a mge", 1000)
}
