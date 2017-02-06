package rolling

import (
	"time"
)

type Point struct {
	Time  time.Time
	Value int64
}

type Aggregator func(point []Point) Point

// Got to be a better name than this.
type Window struct {
	rows []row
	aggregator Aggregator
}

type row struct {
	points []Point
	next int
}

func New(aggregator Aggregator, widths []int) *Window {
	w := Window{
		aggregator: aggregator,
		rows: make([]row, len(widths)),
	}
	
	for i := range widths {
		if widths[i] <= 0 {
			panic("Widths must be >= 0")
		}
		w.rows[i] = row{points:make([]Point, widths[i])}
	}
	
	return &w
}

func (w *Window) Push(p Point) {

}

func (w *Window) Flush(level int) {

}

func (w *Window) Get(level int) []Point {

}

func (w *Window) Close() {

}

func (w *Window) Subscribe(level int) chan []Point {

}
