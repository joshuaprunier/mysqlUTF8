package mysqlUTF8

import (
	"fmt"
	"testing"
)

func TestEncodeFilename(t *testing.T) {
	var testStrings = []struct {
		in      string
		want    string
		comment string
	}{
		{"ascii", "ascii", "ascii no encoding"},
		{"À", "@0G", "single unicode character"},
		{"¢", "@00a2", "hex encoding example"},
		{"×Ø", "@00d7@1K", "multiple unicode"},
		{"Àà", "@0G@0g", "upper and lower case unicode"},
		{"ΠmixedΩ", "@8Imixed@7P", "mixed unicode and ascii"},
		{"ẞⅩⓐ＊Ａ", "@1e9e@P9@@a@ff0a@A@", "higher unicode"},
		{"(╯°□°)╯︵ ┻━┻", "@0028@256f@00b0@25a1@00b0@0029@256f@fe35@0020@253b@2501@253b", "table flip"},
	}

	for _, s := range testStrings {
		got := EncodeFilename(s.in)

		if got != s.want {
			t.Errorf("EncodeFilename(%q) == %q, want %q", s.in, got, s.want)
		}
	}
}

func TestNeedsEncoding(t *testing.T) {
	var testStrings = []struct {
		in   string
		want bool
	}{
		{"ascii", false},
		{"_", false},
		{"(", true},
		{"À", true},
		{"¢", true},
		{"Àà", true},
		{"ΠmixedΩ", true},
	}

	for _, s := range testStrings {
		got := NeedsEncoding(s.in)

		if got != s.want {
			t.Errorf("NeedsEncoding(%q) == %t, want %t", s.in, got, s.want)
		}
	}
}

func ExampleEncodeFilename() {
	fmt.Println(EncodeFilename("test"))
	fmt.Println(EncodeFilename("¿"))
	fmt.Println(EncodeFilename("(╯°□°)╯︵ ┻━┻"))
	// Output:
	// test
	// @00bf
	// @0028@256f@00b0@25a1@00b0@0029@256f@fe35@0020@253b@2501@253b
}

func ExampleNeedsEncoding() {
	fmt.Println(NeedsEncoding("test"))
	fmt.Println(NeedsEncoding("¿"))
	fmt.Println(NeedsEncoding("¢ent"))
	// Output:
	// false
	// true
	// true
}
