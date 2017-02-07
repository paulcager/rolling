package rolling

import (
	"sync/atomic"
	"time"
)

type Point struct {
	Time  time.Time
	Value int64
}

type Aggregator func(points []Point) Point

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

func (w *Window) Push(p Point) {
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
		w.Push(zero)
	}
}

func (w *Window) Get(level int) []Point {
	r := &w.rows[level]
	ret := make([]Point, len(r.points))
	n := copy(ret, r.points[:r.next+1])
	copy(ret[n:], r.points[n:])
	return ret
}

func (w *Window) Close() {

}

func (w *Window) Subscribe(level int) chan []Point {
	panic("TODO")
}

func Max(points []Point) Point {
	// Must be at least one element long.
	max := points[0]
	for i := 1; i < len(points); i++ {
		if points[i].Value > max.Value {
			max = points[i]
		}
	}

	return max
}

func Sum(points []Point) Point {
	sum := points[0]
	for i := 1; i < len(points); i++ {
		sum.Value += points[i].Value
	}

	return sum
}
