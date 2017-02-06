package rolling

import (
	"time"
)

type Point struct {
	Time  time.Time
	Value int64
}

// Got to be a better name than this.
type Window struct {
	x [][]Point
}

func New(widths []int) *Window {

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
