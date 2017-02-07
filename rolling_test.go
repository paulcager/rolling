package rolling

import (
	"runtime"
	"testing"
)

func TestGet(t *testing.T) {
	w := New(Max, 5, 10)
	w.Push(Point{Value: 1})
	w.Push(Point{Value: 2})
	w.Push(Point{Value: 3})

	testEquals(t, w, []int64{0, 0, 1, 2, 3})

	w.Push(Point{Value: 4})
	testEquals(t, w, []int64{0, 1, 2, 3, 4})

	w.Push(Point{Value: 5})
	testEquals(t, w, []int64{1, 2, 3, 4, 5})

	w.Push(Point{Value: 6})
	testEquals(t, w, []int64{2, 3, 4, 5, 6})

	for i := 7; i < 12; i++ {
		w.Push(Point{Value: int64(i)})
	}

	testEquals(t, w, []int64{7, 8, 9, 10, 11})
}

func testEquals(t *testing.T, w *Window, b []int64) bool {
	_, _, line, _ := runtime.Caller(1)
	a := make([]int64, 0)
	for _, p := range w.Get(0) {
		a = append(a, p.Value)
	}
	if len(a) != len(b) {
		t.Error("Line", line, w, b)
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			t.Errorf("Line %v: Expected %v, got %v", line, b, a)
			t.Errorf("W: %v", w.rows[0])
			return false
		}
	}
	return true
}
