package greedy

func GasStation(gas []int, cost []int) int {
	// Local effective fuel at the current index
	localFuel := 0
	// Global fuel which is fuel left after journey
	globalFuel := 0
	// Starting index of the journey
	index := 0
	// Process the array
	for i := 0; i < len(gas); i++ {
		globalFuel += gas[i] - cost[i]
		localFuel += gas[i] - cost[i]
		if localFuel < 0 {
			localFuel = 0
			index = i + 1
		}
	}
	if globalFuel >= 0 {
		return index
	}
	return -1
}
