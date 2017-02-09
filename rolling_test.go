package rolling

import (
	"runtime"
	"testing"
)

func TestGet(t *testing.T) {
	w := New(Max, 5, 10)
	w.PushPoint(Point{Value: 1})
	w.PushPoint(Point{Value: 2})
	w.PushPoint(Point{Value: 3})

	testEquals(t, w, 0, []int64{0, 0, 1, 2, 3})

	w.PushPoint(Point{Value: 4})
	testEquals(t, w, 0, []int64{0, 1, 2, 3, 4})

	w.PushPoint(Point{Value: 5})
	testEquals(t, w, 0, []int64{1, 2, 3, 4, 5})

	w.PushPoint(Point{Value: 6})
	testEquals(t, w, 0, []int64{2, 3, 4, 5, 6})

	for i := 7; i < 12; i++ {
		w.PushPoint(Point{Value: int64(i)})
	}

	testEquals(t, w, 0, []int64{7, 8, 9, 10, 11})
	testEquals(t, w, 1, []int64{0, 0, 0, 0, 0, 0, 0, 0, 5, 10})
}

func testEquals(t *testing.T, w *Window, level int, b []int64) bool {
	_, _, line, _ := runtime.Caller(1)
	a := make([]int64, 0)
	for _, p := range w.Get(level) {
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
