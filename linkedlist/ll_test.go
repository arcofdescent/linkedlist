package linkedlist

import "testing"

func TestCreate(t *testing.T) {
	li := CreateNew()

	if len(li.GetItems()) != 0 {
		t.Error("new list should be empty")
	}
}

func TestAppend(t *testing.T) {
	expected := []int{2, 4, 0}

	l := CreateNew()
	l.Append(2)
	l.Append(4)
	l.Append(0)

	got := l.GetItems()

	for i := range got {
		if got[i] != expected[i] {
			t.Errorf("%d != %d for index %d", got[i], expected[i], i)
		}
	}
}

func TestPrepend(t *testing.T) {
	expected := []int{6, 0, 2, 4}

	l := CreateNew()
	l.Append(2)
	l.Append(4)
	l.Prepend(0)
	l.Prepend(6)

	got := l.GetItems()

	for i := range got {
		if got[i] != expected[i] {
			t.Errorf("%d != %d for index %d", got[i], expected[i], i)
		}
	}
}

func TestGetFirst(t *testing.T) {
	expected := 42

	l := CreateNew()
	l.Append(56)
	l.Prepend(7)
	l.Append(9)
	l.Prepend(42)
	l.Append(78)

	got := l.GetFirst()
	if got != expected {
		t.Errorf("%d != %d", got, expected)
	}
}

func TestGetLast(t *testing.T) {
	expected := 78

	l := CreateNew()
	l.Append(56)
	l.Prepend(7)
	l.Append(9)
	l.Append(78)
	l.Prepend(42)

	got := l.GetLast()
	if got != expected {
		t.Errorf("%d != %d", got, expected)
	}
}

func TestGetAndRemoveFirst(t *testing.T) {
	expected := 42
	expected_r := 7

	l := CreateNew()
	l.Append(56)
	l.Prepend(7)
	l.Append(9)
	l.Prepend(42)
	l.Append(78)

	got := l.GetAndRemoveFirst()
	if got != expected {
		t.Errorf("%d != %d", got, expected)
	}

	got = l.GetFirst()
	if got != expected_r {
		t.Errorf("%d != %d", got, expected_r)
	}
}

func TestGetAndRemoveLast(t *testing.T) {
	expected := 78
	expected_r := 9

	l := CreateNew()
	l.Append(56)
	l.Prepend(7)
	l.Append(9)
	l.Prepend(42)
	l.Append(78)

	got := l.GetAndRemoveLast()
	if got != expected {
		t.Errorf("%d != %d", got, expected)
	}

	got = l.GetLast()
	if got != expected_r {
		t.Errorf("%d != %d", got, expected_r)
	}
}
