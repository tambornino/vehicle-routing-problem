package util

import (
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
	t "vehicle-routing-problem/types"
)

func ImportCoordsToSlice() []t.Route {
	filePath := os.Args[1]
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()

	var idx int
	var a, b t.Point
	var routes []t.Route

	for {
		_, err := fmt.Fscanf(f, "%d (%f,%f) (%f,%f)", &idx, &a.X, &a.Y, &b.X, &b.Y)

		// ignore non-matching lines and break loop at EOF
		if err != nil {
			if err.Error() == "unexpected newline" || err.Error() == "input does not match format" || err.Error() == "expected integer" {
				continue
			}
			if err == io.EOF || err.Error() == "unexpected EOF" {
				break // done reading file
			}
			fmt.Println(err.Error())
			os.Exit(1)
		}

		route := t.Route{
			A: t.Point{a.X, a.Y, Dist(a, t.Point{0,0, 0})},
			B: t.Point{b.X, b.Y, Dist(b, t.Point{0,0,0})},
			RouteLength: Dist(a, b),
			Complete: false,
		}

		routes = append(routes, route)
	}

	return routes
}

func Dist(p1, p2 t.Point) float64 {
	return math.Sqrt(math.Pow(p2.X - p1.X, 2) + math.Pow(p2.Y - p1.Y, 2))
}

func NumLeftToDeliver(rts []t.Route) int {
	var num int
	for _, r := range rts {
		if !r.Complete {
			num++
		}
	}
	return num
}

func FindNearestUndeliveredNeighbor(p t.Point, rts []t.Route) int {
	var idx int
	var min float64

	for i, r := range rts {
		if !r.Complete && (min == 0 || Dist(p, r.A) < min) {
			min = Dist(p, r.A)
			idx = i
		}
	}

	return idx
}

func IntsToStr(a []int) string {
	b := make([]string, len(a))
	for i, v := range a {
		b[i] = strconv.Itoa(v)
	}

	return strings.Trim(strings.Join(b, ","), " ")
}
