package math

import (
	"math"
	"time"
)

// A Point represents a two dimensional Cartesian coordinate
type Point struct {
	X float64
	Y float64
}

func secondHandPoint(t time.Time) Point {
	return angleToPoint(secondsInRadians(t))
}

func secondsInRadians(t time.Time) float64 {
	return (math.Pi / (30 / (float64(t.Second()))))
}

func minuteHandPoint(t time.Time) Point {
	return angleToPoint(minutesInRadians(t))
}

func minutesInRadians(t time.Time) float64 {
	return (math.Pi / (30 / (float64(t.Minute())))) + (secondsInRadians(t) / 60)
}

func angleToPoint(angle float64) Point {
	return Point{X: math.Sin(angle), Y: math.Cos(angle)}
}
