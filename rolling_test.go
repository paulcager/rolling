package rolling

import (
	"testing"
)

func TestGet(t *testing.T) {
	w := New(Max, 5, 10)
	w.Push(Point{Value: 1})
	w.Push(Point{Value: 2})
	w.Push(Point{Value: 3})
	t.Log("First bit")
	testEquals(t, w.rows[0].points, []int64{1, 2, 3, 0, 0})

	t.Log("Second bit")
	testEquals(t, w.Get(0), []int64{1, 2, 3, 0, 0})

	w.Push(Point{Value: 4})
	testEquals(t, w.Get(0), []int64{1, 2, 3, 4, 0})

	w.Push(Point{Value: 5})
	testEquals(t, w.Get(0), []int64{1, 2, 3, 4, 5})

	w.Push(Point{Value: 6})
	testEquals(t, w.Get(0), []int64{2, 3, 4, 5, 6})

}

func testEquals(t *testing.T, a []Point, b []int64) bool {
	if len(a) != len(b) {
		t.Error(a, b)
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i].Value != b[i] {
			t.Error(a, b)
			return false
		}
	}
	return true
}
