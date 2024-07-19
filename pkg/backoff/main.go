package backoff

import (
	"fmt"
	"time"

	"golang.org/x/exp/rand"
)

func backoff(attempt int) time.Duration {
	min := 1 * time.Second
	max := 1 * time.Hour

	// Wait time based on attempt number and a factor of randomness
	sleep := min * (1 << attempt)
	if sleep > max {
		sleep = max
	}
	jitter := time.Duration(rand.Int63n(int64(min)))
	return sleep + jitter
}

func RetryWithBackoff(operation func() error, errorMessage string, successMessage string) error {
	for attempt := 0; ; attempt++ {
		if err := operation(); err == nil {
			fmt.Println(successMessage)
			return nil
		} else {
			wait := backoff(attempt)
			fmt.Printf("%s %v seconds\n", errorMessage, wait)
			time.Sleep(wait)
		}
	}
}
