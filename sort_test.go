package sort_test

import (
	"reflect"
	"testing"

	"github.com/karlhepler/sort"
)

func TestSelection(t *testing.T) {
	for i, tc := range testCases() {
		if output := sort.Selection(tc.input); !reflect.DeepEqual(tc.output, output) {
			t.Errorf("%d. Expected %v; Received %v\n", i, tc.output, output)
		}
	}
}

func TestInsertion(t *testing.T) {
	for i, tc := range testCases() {
		if output := sort.Insertion(tc.input); !reflect.DeepEqual(tc.output, output) {
			t.Errorf("%d. Expected %v; Received %v\n", i, tc.output, output)
		}
	}
}

func TestBubble(t *testing.T) {
	for i, tc := range testCases() {
		if output := sort.Bubble(tc.input); !reflect.DeepEqual(tc.output, output) {
			t.Errorf("%d. Expected %v; Received %v\n", i, tc.output, output)
		}
	}
}

func TestMerge(t *testing.T) {
	for i, tc := range testCases() {
		if output := sort.Merge(tc.input); !reflect.DeepEqual(tc.output, output) {
			t.Errorf("%d. Expected %v; Received %v\n", i, tc.output, output)
		}
	}
}

func TestQuick(t *testing.T) {
	for i, tc := range testCases() {
		if output := sort.Quick(tc.input, 0, len(tc.input)-1); !reflect.DeepEqual(tc.output, output) {
			t.Errorf("%d. Expected %v; Received %v\n", i, tc.output, output)
		}
	}
}

func testCases() []struct {
	input  []int
	output []int
} {
	return []struct {
		input  []int
		output []int
	}{
		{
			[]int{},
			[]int{},
		},
		{
			[]int{42},
			[]int{42},
		},
		{
			[]int{2, 1},
			[]int{1, 2},
		},
		{
			[]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			[]int{5, 3, 7, 1, 2, 0, 9, 5, 5, 5, 2},
			[]int{0, 1, 2, 2, 3, 5, 5, 5, 5, 7, 9},
		},
		{
			[]int{2, 3, 8, 6, 1, 5, 4, 7},
			[]int{1, 2, 3, 4, 5, 6, 7, 8},
		},
	}
}
