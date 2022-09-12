package safeStack

import (
	"testing"
	"vmWare/server/urlStruct"
)

func Setup() SafeStack {

	testArray := []urlStruct.UrlInformation{
		{
			Url: "www.testOne.com",
			Views: 3.0,
			RelevanceScore: 100,
		},
		{
			Url: "www.testTwo.com",
			Views: 6.0,
			RelevanceScore: 200,
		},
		{
			Url: "www.testThree.com",
			Views: 1.0,
			RelevanceScore: 300,
		},
	}

	testStruct := urlStruct.UrlList{
		Count: 3,
		Data:  testArray,
	}

	testStack := SafeStack{
		stack: testStruct,
	}

	return testStack
}

func AdditionalData() urlStruct.UrlList {

	testArray := []urlStruct.UrlInformation{
		{
			Url: "www.testFour.com",
			Views: 1.0,
			RelevanceScore: 200,
		},
		{
			Url: "www.testFive.com",
			Views: 10.0,
			RelevanceScore: 400,
		},
	}

	testStruct := urlStruct.UrlList{
		Count: 3,
		Data:  testArray,
	}

	return testStruct
}

func TestUpdate(t *testing.T) {
	t.Log("->")

	testStack := Setup()
	newData := AdditionalData()
	testStack.Update(&newData)
	
	got := testStack.ReturnSize()
	want := 5

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	} else {
		t.Logf("%d == %d", got, want)
	}
}

func TestSortStackRelevanceScore(t *testing.T) {
	t.Log("->")

	testStack := Setup()
	testStack.SortStackRelevanceScore()
	
	got := testStack.stack.Data[0].RelevanceScore
	want := 300.0

	if got != want {
		t.Errorf("got %f, wanted %f", got, want)
	} else {
		t.Logf("%f == %f", got, want)
	}
}

func TestSortStackViews(t *testing.T) {
	t.Log("->")

	testStack := Setup()
	testStack.SortStackViews()
	
	got := testStack.stack.Data[0].Views
	want := 6.0

	if got != want {
		t.Errorf("got %f, wanted %f", got, want)
	} else {
		t.Logf("%f == %f", got, want)
	}
}

func TestReturnSubStack(t *testing.T) {
	t.Log("->")

	testStack := Setup()
	subStack := testStack.ReturnSubStack(2)
	
	got := len(subStack)
	want := 2

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	} else {
		t.Logf("%d == %d", got, want)
	}
}

func TestReturnSize(t *testing.T) {
	t.Log("->")

	testStack := Setup()
	
	got := testStack.ReturnSize()
	want := 3

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	} else {
		t.Logf("%d == %d", got, want)
	}
}

