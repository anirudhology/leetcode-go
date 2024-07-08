package strings_test

import (
	"reflect"
	"testing"

	"github.com/anirudhology/leetcode-go/problems/strings"
)

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		name string
		strs []string
		want [][]string
	}{
		{
			name: "Empty array",
			strs: []string{},
			want: [][]string{},
		},
		{
			name: "Single string",
			strs: []string{"abc"},
			want: [][]string{{"abc"}},
		},
		{
			name: "Anagrams",
			strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			want: [][]string{
				{"eat", "tea", "ate"},
				{"tan", "nat"},
				{"bat"},
			},
		},
		{
			name: "No anagrams",
			strs: []string{"abc", "def", "ghi"},
			want: [][]string{
				{"abc"},
				{"def"},
				{"ghi"},
			},
		},
		{
			name: "Multiple identical strings",
			strs: []string{"abc", "abc", "abc"},
			want: [][]string{
				{"abc", "abc", "abc"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strings.GroupAnagrams(tt.strs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GroupAnagrams() = %v, want %v", got, tt.want)
			}
		})
	}
}
