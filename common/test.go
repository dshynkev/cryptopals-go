package common

import (
	"reflect"
	"testing"
)

func Test(t *testing.T, want interface{}, got interface{}) {
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("= want =\n%+v\n= got =\n%+v\n", want, got)
	}
}
