package utils

import (
	"testing"
)

func TestMinOne(t *testing.T){
	t.Log("->")
    got := Min(1, 5)
    want := 1

    if got != want {
		  t.Errorf("got %d, wanted %d", got, want)
    } else {
      t.Logf("%d == %d", got, want)
    }
}

func TestMinTwo(t *testing.T){
	t.Log("->")
    got := Min(-10, 5)
    want := -10

    if got != want {
		  t.Errorf("got %d, wanted %d", got, want)
    } else {
      t.Logf("%d == %d", got, want)
    }
}