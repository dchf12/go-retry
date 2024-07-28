package retry

import (
	"time"
)

type RetryableFunc func() error

type Options struct {
	MaxRetries int
	Delay      time.Duration
}

type Option func(*Options)

func WithMaxRetries(maxRetries int) Option {
	return func(opts *Options) {
		opts.MaxRetries = maxRetries
	}
}

func WithDelay(delay time.Duration) Option {
	return func(opts *Options) {
		opts.Delay = delay
	}
}

func Retry(fn RetryableFunc, options ...Option) error {
	opts := Options{
		MaxRetries: 3,
		Delay:      1 * time.Second,
	}
	for _, opt := range options {
		opt(&opts)
	}
	var err error
	for range opts.MaxRetries {
		err = fn()
		if err == nil {
			return nil
		}
		time.Sleep(opts.Delay)
	}
	return err
}
