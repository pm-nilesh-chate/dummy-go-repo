package p1_test

import (
	"testing"

	"github.com/chatenilesh/dummy-go-repo/p1"
)

func TestFoo(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "happy path",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(_ *testing.T) {
			p1.Foo()
		})
	}
}
