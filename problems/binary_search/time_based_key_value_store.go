package binary_search

type Node struct {
	values     []string
	timestamps []int
}

type TimeMap struct {
	entries map[string]*Node
}

func TimeMapConstructor() TimeMap {
	return TimeMap{
		entries: make(map[string]*Node),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	if _, exists := this.entries[key]; !exists {
		this.entries[key] = &Node{}
	}
	node := this.entries[key]
	node.values = append(node.values, value)
	node.timestamps = append(node.timestamps, timestamp)
}

func (this *TimeMap) Get(key string, timestamp int) string {
	if _, exists := this.entries[key]; !exists {
		return ""
	}
	// Get both values and timestamps for the current key
	values := this.entries[key].values
	timestamps := this.entries[key].timestamps
	// Left and right pointers for the binary search
	left := 0
	right := len(timestamps) - 1
	// Final result
	result := ""
	// Process the array from both ends
	for left <= right {
		middle := left + int((right-left)/2)
		if timestamps[middle] <= timestamp {
			result = values[middle]
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return result
}
