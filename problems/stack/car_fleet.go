package stack

import "sort"

func CarFleet(target int, position []int, speed []int) int {
	// Special case
	if target < 0 || len(position) == 0 || len(speed) == 0 {
		return 0
	}
	// Total number of cars
	n := len(position)
	// Array to store cars' positions and times taken to reach target
	cars := make([][2]float64, n)
	for i := 0; i < n; i++ {
		cars[i][0] = float64(position[i])
		cars[i][1] = float64(target-position[i]) / float64(speed[i])
	}
	// Sort this array in descending order by position
	sort.Slice(cars, func(i, j int) bool {
		return cars[i][0] > cars[j][0]
	})
	// Last time registered
	lastTime := 0.0
	// Fleet counter
	fleets := 0
	// Process all the cars
	for i := 0; i < n; i++ {
		if cars[i][1] > lastTime {
			fleets++
			lastTime = cars[i][1]
		}
	}
	return fleets
}
