package types

type Point struct {
	X float64
	Y float64
	DistFromBase float64
}

type Route struct {
	A Point
	B Point
	RouteLength float64
	Complete bool
}

type Driver struct {
	MinutesElapsed float64
	LoadsTaken []int
}
