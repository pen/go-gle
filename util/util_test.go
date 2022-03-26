//nolint:varnamelen
package util_test

import (
	"os"
	"testing"

	"github.com/pen/go-gle/util"
)

func TestGetEncodedEnv(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		arg  string
		want string
	}{
		{
			name: "raw",
			arg:  `{"name":"Smith"}`,
			want: `{"name":"Smith"}`,
		},
		{
			name: "base64",
			arg:  `eyJuYW1lIjoiU21pdGgifQ==`,
			want: `{"name":"Smith"}`,
		},
		{
			name: "xz | base64",
			arg:  `/Td6WFoAAATm1rRGAgAhARwAAAAQz1jMAQAPeyJuYW1lIjoiU21pdGgifQBiI68qCK3i4gABKBDlC2xgH7bzfQEAAAAABFla`,
			want: `{"name":"Smith"}`,
		},
		{
			name: "bzip2 | base64",
			arg:  `QlpoOTFBWSZTWY0EjzwAAAcbgBAAABAIACJjBAogACIAGQQNA0Jc3KAFgKPP4u5IpwoSEaCR54A=`,
			want: `{"name":"Smith"}`,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			os.Setenv("TEST_KEY", tt.arg)
			got, _ := util.GetEncodedEnv("TEST_KEY")
			if string(got) != tt.want {
				t.Errorf("NewMultiEncodedStringReader() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStripQuotes(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		arg  string
		want string
	}{
		{arg: ``, want: ``},
		{arg: `'`, want: `'`},
		{arg: `"`, want: `"`},
		{arg: `''`, want: ``},
		{arg: `""`, want: ``},
		{arg: `'''`, want: `'`},
		{arg: `'single'`, want: `single`},
		{arg: `"double"`, want: `double`},
		{arg: "`backtick`", want: "`backtick`"},
		{arg: `"'one level'"`, want: `'one level'`},
		{arg: `"not match 1`, want: `"not match 1`},
		{arg: `'not match 2"`, want: `'not match 2"`},
	}

	for _, tt := range tests {
		tt := tt
		t.Run("["+tt.arg+"]", func(t *testing.T) {
			t.Parallel()

			got := util.StripQuotes(tt.arg)
			if got != tt.want {
				t.Errorf("StripQuotes() = [%v], want [%v]", got, tt.want)
			}
		})
	}
}
