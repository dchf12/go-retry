package retry

import (
	"errors"
	"testing"
)

func TestRetry(t *testing.T) {
	t.Parallel()
	retryFunc := func() error {
		return errors.New("failure")
	}
	err := Retry(retryFunc, WithDelay(2), WithMaxRetries(2))
	if err == nil {
		t.Error("expected an error")
	}
}
