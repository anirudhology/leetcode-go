package stack

func DailyTemperatures(temperatures []int) []int {
    // Special case
    if len(temperatures) == 0 {
        return temperatures
    }
    // Length of the array
    n := len(temperatures)
    // Array to store the final output
    answer := make([]int, n)
    // Stack to store the indices of the elements
    stack := make([]int, 0, n)
    // Top of the stack
    top := -1
    // Process the array from right to left
    for i := n - 1; i >= 0; i-- {
        // Since we are moving from right to left, the next
        // greater element will be at the top of the stack
        // and we can consider it
        for top > -1 && temperatures[i] >= temperatures[stack[top]] {
            top--
            stack = stack[:top+1] // pop from stack
        }
        // If the stack is not empty at this time, it means we have next greater
        // element represented by the top, and we can find the distance between
        // current index and top
        if top > -1 {
            answer[i] = stack[top] - i;
        }
        // Insert current index
        top++
        stack = append(stack, i)
    }
    return answer
}