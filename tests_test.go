package main

import (
	"reflect"
	"testing"
)

func Test_initUniverse(t *testing.T) {
	tests := map[string]struct {
		want [][]string
	}{
		"empty universe": {want: [][]string{{"O", " ", "O", "O"}, {"O", "O", " ", "O"}, {"O", " ", "O", " "}, {"O", "O", " ", " "}}},
	}

	for name, tc := range tests {
		got := initUniverse(4, 4)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("%s: got: %v, want: %v", name, got, tc.want)
		}
	}
}
