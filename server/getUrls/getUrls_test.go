package getUrls

import (
	"testing"
	safeStack "vmWare/server/safeStack"
	val "vmWare/server/values"
)

func TestGetURLInfo(t *testing.T) {
	t.Log("->")

	stack := &safeStack.SafeStack{}
	err := GetURLInfo(stack, val.GOOGLE)
	got := stack.ReturnSize()
	want := 5

	if got != want || err != nil {
		t.Errorf("got %d, wanted %d", got, want)
	} else {
		t.Logf("%d == %d", got, want)
	}

}

func TestGetAllURLS(t *testing.T) {
	t.Log("->")

	stack := &safeStack.SafeStack{}

	// Download data, and check number of entries
	err := GetAllURLS(stack, val.DUCKDUCKGO, val.GOOGLE, val.WIKIPEDIA)

	got := stack.ReturnSize()
	want := 15

	if got != want || err != nil {
		t.Errorf("got %d, wanted %d", got, want)
	} else {
		t.Logf("%d == %d", got, want)
	}
}
