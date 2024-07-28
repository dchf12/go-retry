package retry

import (
	"errors"
	"time"
)

type RetryableFunc func() error

type Options struct {
	MaxRetries int
	Delay      time.Duration
}

func Retry(fn RetryableFunc, opts Options) error {
	var err error
	for range opts.MaxRetries {
		err = fn()
		if err == nil {
			return nil
		}
		time.Sleep(opts.Delay)
	}
	return errors.New("all retries failed")
}
