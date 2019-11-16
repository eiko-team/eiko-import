package formating

import (
	"fmt"
	"reflect"
	"testing"
)

func Example_split() {
	fmt.Printf("%q", split("item1, item2(test1, test2), item3", ','))

	// Output:
	// ["item1" "item2(test1, test2)" "item3"]
}

func Test_split(t *testing.T) {
	type args struct {
		s   string
		sep rune
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"single", args{"test", '€'}, []string{"test"}},
		{"double", args{"t€st", '€'}, []string{"t", "st"}},
		{"list", args{"t, to, tok, toke, token", ','}, []string{"t", "to", "tok", "toke", "token"}},
		{"parenthese", args{"t, to(test), tok", ','}, []string{"t", "to(test)", "tok"}},
		{"no split in parenthese", args{"t, to(test1, test2), tok", ','}, []string{"t", "to(test1, test2)", "tok"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := split(tt.args.s, tt.args.sep)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("split() = \n%q, want \n%q", got, tt.want)
			}
		})
	}
}

func Example_betterList() {
	fmt.Printf("%q", betterList("item1, item2(test1, test2), item3."))

	// Output:
	// ["item1" "item2(test1, test2)" "item3"]
}

func Test_betterList(t *testing.T) {
	type args struct {
		list string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"simple", args{"t."}, []string{"t"}},
		{"list", args{"t, tok."}, []string{"t", "tok"}},
		{"goal", args{"t, to(test1, test2), tok."}, []string{"t", "to(test1, test2)", "tok"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := betterList(tt.args.list); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("betterList() = %v, want %v", got, tt.want)
			}
		})
	}
}
