package main

import "testing"

var route = Route{1, 3, 4, 2}
var matchingRoute = FoundRoutes{
	{1, 3, 4, 2},
}
var nonMatchingRoute = Route{}

func TestBacktrack(t *testing.T) {
	var result = backtrack(route)
	var expect = Route{1, 3, 4}

	for i, v := range result {
		if v != expect[i] {
			t.Errorf("Backtracked route was incorrect, got %v but expected %v", v, expect[i])
		}
	}
}
