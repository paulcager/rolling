package rolling

import (
	"sync/atomic"
	"time"
)

// A point is an int64 value at a certain time.
type Point struct {
	Time  time.Time
	Value int64
}

// An Aggregator is a function that takes a slice of points and returns a single point, for example
// Sum or Max.
type Aggregator func(points []Point) Point

// A Window is a rolling window over the data.
// Got to be a better name than this.
type Window struct {
	rows       []row
	aggregator Aggregator
	subs       atomic.Value // of []chan[]Point
}

type row struct {
	points []Point
	next   int
}

// New creates a new Window. The
func New(aggregator Aggregator, widths ...int) *Window {
	w := Window{
		aggregator: aggregator,
		rows:       make([]row, len(widths)),
	}

	for i := range widths {
		if widths[i] <= 0 {
			panic("Widths must be >= 0")
		}
		w.rows[i] = row{points: make([]Point, widths[i])}
	}

	return &w
}

func (w *Window) Push(value int64) {
	w.PushPoint(Point{time.Now(), value})
}

func (w *Window) PushPoint(p Point) {
	for level := 0; level < len(w.rows); level++ {
		r := &w.rows[level]
		r.points[r.next] = p
		r.next++
		if r.next < len(r.points) {
			return
		}
		r.next = 0
		p = w.aggregator(r.points)
	}
}

func (w *Window) Flush(level int) {
	r := &w.rows[level]
	zero := Point{Time: time.Now()}
	for r.next > 0 {
		w.PushPoint(zero)
	}
}

func (w *Window) Get(level int) []Point {
	r := &w.rows[level]
	ret := make([]Point, len(r.points))
	copy(ret, r.points[r.next:])
	copy(ret[len(r.points)-r.next:], r.points)
	return ret
}

func (w *Window) Close() {

}

func (w *Window) Subscribe(level int) chan []Point {
	panic("TODO")
}

func Max(points []Point) Point {
	// We know it is at least one element long.
	max := points[0]
	for i := 1; i < len(points); i++ {
		if points[i].Value > max.Value {
			max = points[i]
		}
	}

	return max
}

func Avg(points []Point) Point {
	if len(points) == 0 {
		return Point{}
	}
	sum := Sum(points)
	sum.Time = points[0].Time
	sum.Value = sum.Value / int64(len(points))
	return sum
}

func Sum(points []Point) Point {
	sum := points[0]
	for i := 1; i < len(points); i++ {
		sum.Value += points[i].Value
	}

	return sum
}
