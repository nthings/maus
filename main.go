package main

import (
	"github.com/go-vgo/robotgo"
)

func main() {
  robotgo.MoveMouseSmooth(100, 200, 1.0, 100.0)
  robotgo.MoveMouseSmooth(200, 200, 1.0, 100.0)
}
