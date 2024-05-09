package main

import (
	"fmt"

	t "vehicle-routing-problem/types"
	u "vehicle-routing-problem/util"
)

const FULL_DAY = float64(720)
var drivers []t.Driver

func main() {
	routes := u.ImportCoordsToSlice()

	for u.NumLeftToDeliver(routes) > 0 {
		// starting from base with a new driver
		var driver = t.Driver{MinutesElapsed: 0, LoadsTaken: []int{}}
		currentPt := t.Point{X:0, Y: 0, DistFromBase: 0}
		enRoute := true

		// let driver continue delivering until time is up
		for enRoute {
			idx := u.FindNearestUndeliveredNeighbor(currentPt, routes)
			if routes[idx].Complete {
				enRoute = false
				continue
			}
			next := routes[idx]

			// complete next route if driver has enough time
			if driver.MinutesElapsed + u.Dist(currentPt, next.A) + next.RouteLength + next.B.DistFromBase < FULL_DAY {
				driver.MinutesElapsed += u.Dist(currentPt, next.A) + next.RouteLength
				driver.LoadsTaken = append(driver.LoadsTaken, idx+1)
				routes[idx].Complete = true
				currentPt = next.B
			} else {
				driver.MinutesElapsed += currentPt.DistFromBase
				enRoute = false
			}
		}

		drivers = append(drivers, driver)
	}

	// print out driver lists
	for _, dr := range drivers {
		fmt.Printf("[%s]\n", u.IntsToStr(dr.LoadsTaken))
	}
}
