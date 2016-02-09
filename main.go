package main

import (
	"github.com/josephspurrier/ball"
	"github.com/josephspurrier/ball/greenball"
	"github.com/josephspurrier/ball/redball"
	"github.com/josephspurrier/beachball"
)

import (
	"fmt"
)

func main() {
	// Default ball
	b := ball.New()
	fmt.Println(b.Color())

	// Red ball
	r := redball.New()
	fmt.Println(r.Color())

	// Green ball
	g := greenball.New()
	fmt.Println(g.Color())

	// Beachball which is actually the default ball
	e := beachball.New()
	fmt.Println(e.Color())
}
