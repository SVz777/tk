//go:build go1.16 && linux

package goid_test

import (
	"fmt"
	"testing"

	"github.com/SVz777/tk/goid"
)

func BenchmarkSlowGetGoID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = goid.SlowGetGoID()
	}
}

func BenchmarkGetGoID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = goid.GetGoID()
	}
}

func TestGetGoID(t *testing.T) {
	fmt.Println(goid.GetGoID(), goid.SlowGetGoID())
}
