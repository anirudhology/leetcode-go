package heaps

func TaskScheduler(tasks []byte, n int) int {
	// Special case
	if len(tasks) == 0 {
		return 0
	}
	// Array to store frequencies of all the tasks
	frequencies := make([]int, 26)
	for _, task := range tasks {
		frequencies[task-'A']++
	}
	// Maximum frequency and tasks with maximum frequency
	maxFrequency, maxFrequencyTasks := 0, 0
	for _, frequency := range frequencies {
		if frequency > maxFrequency {
			maxFrequency = frequency
			maxFrequencyTasks = 1
		} else if maxFrequency == frequency {
			maxFrequencyTasks++
		}
	}
	// Maximum time taken to execute all tasks
	time := (n+1)*(maxFrequency-1) + maxFrequencyTasks
	if time > len(tasks) {
		return time
	} else {
		return len(tasks)
	}
}
