package pinchers

import (
	"reflect"
	"testing"
)

func TestParser(t *testing.T) {
	s1 := " hello=world "
	s2 := "'something'=\"yep else\" something2=\"hello  =  = something\""
	s3 := " "
	s4 := "hello \"what's\" up"

	p1 := make(RuntimeVarMap)
	FillRuntimeVarMap(&p1, s1)

	p2 := ToMap(s2)
	p3 := ToMap(s3)
	p4 := ToMap(s4)

	c1 := map[string]RuntimeVar{"hello": RuntimeVar{"world"}}
	c2 := map[string]RuntimeVar{"something": RuntimeVar{"yep else"}, "something2": RuntimeVar{"hello  =  = something"}}
	c3 := map[string]RuntimeVar{}
	c4 := map[string]RuntimeVar{"hello": RuntimeVar{""}, "what's": RuntimeVar{""}, "up": RuntimeVar{""}}

	if !reflect.DeepEqual(p1, c1) {
		t.Error(s1, p1, c1)
	}
	if !reflect.DeepEqual(p2, c2) {
		t.Error(s2, p2, c2)
	}
	if !reflect.DeepEqual(p3, c3) {
		t.Error(s3, p3, c3)
	}
	if !reflect.DeepEqual(p4, c4) {
		t.Error(s4, p4, c4)
	}
}

func TestParseFactPincher(t *testing.T) {
	test := map[string]string{
		"test": "a=1 b=2",
	}

	pincher, err := ToFactPincher(test)
	if err != nil {
		t.Fatalf("ParseFactPincher: %v", err)
	}

	if pincher == nil {
		t.Fatalf("Pincher is nil")
	}
}
