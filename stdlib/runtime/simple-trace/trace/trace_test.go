package trace_test

import (
	"testing"

	"github.com/radisvaliullin/test-golang/stdlib/runtime/simple-trace/trace"
)

func TestTrace(t *testing.T) {

	trace.Trace()

}

func BenchmarkTrace(b *testing.B) {
	for n := 0; n < b.N; n++ {
		trace.Trace()
	}
}
