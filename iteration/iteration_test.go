package iteration

import (
	"testing"
	"strings"
)

func TestRepeat(t *testing.T) {
	t.Run("Test 5 iterations", func(t *testing.T){
		for i := 0; i <=10; i++ {
			repeated := Repeat("a", i)
			expected := strings.Repeat("a", i)
			if repeated != expected {
				t.Errorf("expected %q but got %q", expected, repeated)
			}
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
