package graph

import "container/list"

func CourseSchedule(numCourses int, prerequisites [][]int) bool {
	// Adjacency list for the courses
	adjacencyList := make([][]int, numCourses)
	// Array to store indegree of every node
	indegrees := make([]int, numCourses)
	// Populate both adjacency list and indegrees
	for _, prerequisite := range prerequisites {
		adjacencyList[prerequisite[1]] = append(adjacencyList[prerequisite[1]], prerequisite[0])
		indegrees[prerequisite[0]]++
	}
	// Queue to store nodes with zero indegree
	zeroIndegreeNodes := list.New()
	for i := 0; i < numCourses; i++ {
		if indegrees[i] == 0 {
			zeroIndegreeNodes.PushBack(i)
		}
	}
	// Count for number of courses one can take
	count := 0
	// Process all nodes in queue
	for zeroIndegreeNodes.Len() > 0 {
		count++
		node := zeroIndegreeNodes.Remove(zeroIndegreeNodes.Front()).(int)
		for _, neighbor := range adjacencyList[node] {
			indegrees[neighbor]--
			if indegrees[neighbor] == 0 {
				zeroIndegreeNodes.PushBack(neighbor)
			}
		}
	}
	return numCourses == count
}
