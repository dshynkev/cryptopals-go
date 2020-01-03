package test

import (
	"reflect"
	"testing"
)

func Test(t *testing.T, want interface{}, got interface{}) {
	if !reflect.DeepEqual(want, got) {
		FailTest(t, want, got)
	}
}

func FailTest(t *testing.T, want interface{}, got interface{}) {
	t.Fatalf("= want =\n%+v\n= got =\n%+v\n", want, got)
}
