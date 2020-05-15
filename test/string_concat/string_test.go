package stringconcat

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

var (
	str1 = "asdfasdfasdfasdfasdfasdfasdfasdf"
	str2 = "zxcvzxcvzxcvzxcvzxcvzxcvzxcvzxcv"
)

func BenchmarkSum(b *testing.B) {
	var out string

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		out = "asdfasdfasdfasdfasdfasdfasdfasdf" + "zxcvzxcvzxcvzxcvzxcvzxcvzxcvzxcv"
	}
	b.StopTimer()

	b.Log(out)
}

func BenchmarkSum2(b *testing.B) {
	var out string

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		out = str1 + str2
	}
	b.StopTimer()

	b.Log(out)
}

func BenchmarkFmt(b *testing.B) {
	var out string

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		out = fmt.Sprintf("%s%s", str1, str2)
	}
	b.StopTimer()

	b.Log(out)
}

func BenchmarkJoin(b *testing.B) {
	var out string

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		out = strings.Join([]string{str1, str2}, "")
	}
	b.StopTimer()

	b.Log(out)
}

func BenchmarkBuf(b *testing.B) {
	var buf bytes.Buffer
	var out string

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf.WriteString(str1)
		buf.WriteString(str2)
		out = buf.String()
		buf.Reset()
	}
	b.StopTimer()

	b.Log(out)
}

func BenchmarkBuild(b *testing.B) {
	var build strings.Builder
	var out string

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		build.WriteString(str1)
		build.WriteString(str2)
		out = build.String()
		build.Reset()
	}
	b.StopTimer()

	b.Log(out)
}
