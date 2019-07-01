package ibt_test

import (
	"fmt"
	"testing"
)

func TestTree1(t *testing.T) {
	t.Run(fmt.Sprintf("()"), func(t *testing.T) {
		t.Errorf("() = %v, want : %v")
	})
}
