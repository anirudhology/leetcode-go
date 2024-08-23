package dynamic_programming_test

import (
	"testing"

	"github.com/anirudhology/leetcode-go/problems/dynamic_programming"
	"github.com/stretchr/testify/assert"
)

func TestCanPartition(t *testing.T) {
	assert := assert.New(t)

	assert.True(dynamic_programming.PartitionEqualSubsetSum([]int{1, 5, 11, 5}))
	assert.False(dynamic_programming.PartitionEqualSubsetSum([]int{1, 2, 3, 5}))
	assert.False(dynamic_programming.PartitionEqualSubsetSum([]int{1, 2, 5}))
	assert.True(dynamic_programming.PartitionEqualSubsetSum([]int{1, 2, 3, 4}))
	assert.True(dynamic_programming.PartitionEqualSubsetSum([]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}))
	assert.False(dynamic_programming.PartitionEqualSubsetSum([]int{}))
	assert.False(dynamic_programming.PartitionEqualSubsetSum(nil))
	assert.False(dynamic_programming.PartitionEqualSubsetSum([]int{3}))
}
