package structs

import (
	"testing"
)

func Test_Equal_1(t *testing.T) {
	a := Tag{}
	b := Tag{}
	if !a.Equal(b) {
		t.Errorf("tags not equal: [%+v] != [%+v]", a, b)
	}
}

func Test_Equal_2(t *testing.T) {
	a := Tag{Name: "apa"}
	b := Tag{}
	if a.Equal(b) {
		t.Errorf("tags equal: [%+v] != [%+v]", a, b)
	}
}

func Test_Equal_3(t *testing.T) {
	a := Tag{MediaCount: 1}
	b := Tag{}
	if a.Equal(b) {
		t.Errorf("tags equal: [%+v] != [%+v]", a, b)
	}
}
