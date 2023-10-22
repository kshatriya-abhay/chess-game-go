package util

import (
	"fmt"
	"testing"
)

func AssertTrue(t *testing.T, condition bool, errorMessage string, args ...any) {
	if !condition {
		t.Errorf(fmt.Sprintf(errorMessage, args...))
	}
}