package p1_test

import (
	"testing"

	"github.com/chatenilesh/dummy-go-repo/p1"
)

func TestFoo(t *testing.T) {
	t.Run("happy path", func(_ *testing.T) {
		p1.Foo()
	})
}
