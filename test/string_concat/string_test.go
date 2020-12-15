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
	str3 = "asdfasdfasdfasdfasdfasdfasdfasdf"
	str4 = "zxcvzxcvzxcvzxcvzxcvzxcvzxcvzxcv"
	str5 = "asdfasdfasdfasdfasdfasdfasdfasdf"
	str6 = "zxcvzxcvzxcvzxcvzxcvzxcvzxcvzxcv"
	str7 = "asdfasdfasdfasdfasdfasdfasdfasdf"
	str8 = "zxcvzxcvzxcvzxcvzxcvzxcvzxcvzxcv"
)

func BenchmarkSum(b *testing.B) {
	var out string

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		out = ("asdfasdfasdfasdfasdfasdfasdfasdf" +
			"asdfasdfasdfasdfasdfasdfasdfasdf" + "zxcvzxcvzxcvzxcvzxcvzxcvzxcvzxcv" +
			"asdfasdfasdfasdfasdfasdfasdfasdf" + "zxcvzxcvzxcvzxcvzxcvzxcvzxcvzxcv" +
			"asdfasdfasdfasdfasdfasdfasdfasdf" + "zxcvzxcvzxcvzxcvzxcvzxcvzxcvzxcv" +
			"zxcvzxcvzxcvzxcvzxcvzxcvzxcvzxcv")
	}
	b.StopTimer()

	_ = out
	// b.Log(out)
}

func BenchmarkSum2(b *testing.B) {
	var out string

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		out = (str1 +
			str2 + str3 + str4 + str5 + str6 + str7 +
			str8)
	}
	b.StopTimer()

	_ = out
	// b.Log(out)
}

func BenchmarkFmt(b *testing.B) {
	var out string

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		out = fmt.Sprintf("%s%s%s%s%s%s%s%s", str1,
			str2, str3, str4, str5, str6, str7,
			str8)
	}
	b.StopTimer()

	_ = out
	// b.Log(out)
}

func BenchmarkJoin(b *testing.B) {
	var out string

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		out = strings.Join([]string{str1,
			str2, str3, str4, str5, str6, str7,
			str8}, "")
	}
	b.StopTimer()

	_ = out
	// b.Log(out)
}

func BenchmarkBuf(b *testing.B) {
	var buf bytes.Buffer
	var out string

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf.WriteString(str1)
		buf.WriteString(str2)
		buf.WriteString(str3)
		buf.WriteString(str4)
		buf.WriteString(str5)
		buf.WriteString(str6)
		buf.WriteString(str7)
		buf.WriteString(str8)
		out = buf.String()
		buf.Reset()
	}
	b.StopTimer()

	_ = out
	// b.Log(out)
}

func BenchmarkBuild(b *testing.B) {
	var build strings.Builder
	var out string

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		build.WriteString(str1)
		build.WriteString(str2)
		build.WriteString(str3)
		build.WriteString(str4)
		build.WriteString(str5)
		build.WriteString(str6)
		build.WriteString(str7)
		build.WriteString(str8)
		out = build.String()
		build.Reset()
	}
	b.StopTimer()

	_ = out
	// b.Log(out)
}
