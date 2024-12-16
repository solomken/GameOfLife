package main

import (
	"reflect"
	"testing"
)

func Test_initUniverse(t *testing.T) {
	got := initUniverse(4, 4)
	want := [][]string{{"O", " ", "O", "O"}, {"O", "O", " ", "O"}, {"O", " ", "O", " "}, {"O", "O", " ", " "}}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q want %q", got, want)
	}
}
