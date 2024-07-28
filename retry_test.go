package retry

import (
	"errors"
	"testing"
	"time"
)

func TestRetry(t *testing.T) {
	retryFunc := func() error {
		return errors.New("failure")
	}
	opts := Options{
		MaxRetries: 3,
		Delay:      1 * time.Second,
	}
	err := Retry(retryFunc, opts)
	if err == nil {
		t.Error("expected an error")
	}
}
