package utils 

import (
	"github.com/cenkalti/backoff/v4"
	"time"
)

type (
	Config struct {
		MaxElapsedTime time.Duration
		Retries        int
	}
)

func Retry(MaxElapsedTime time.Duration, retries int, fn func() error) error {
	exponentialBackOff := backoff.NewExponentialBackOff()
	exponentialBackOff.MaxElapsedTime = MaxElapsedTime

	return backoff.Retry(fn, backoff.WithMaxRetries(exponentialBackOff, uint64(retries-1)))
}
