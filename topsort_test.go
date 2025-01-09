// Copyright 2013 Steven Le. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package topsort

import (
	"strings"
	"testing"
)

func TestTopSort(t *testing.T) {
	graph := initGraph()

	// a -> b -> c
	graph.AddEdge("a", "b")
	graph.AddEdge("b", "c")

	results, err := graph.TopSort("a")
	if err != nil {
		t.Error(err)
		return
	}
	if results[0] != "c" || results[1] != "b" || results[2] != "a" {
		t.Errorf("Wrong sort order: %v", results)
	}
}

func TestTopSort2(t *testing.T) {
	graph := initGraph()

	// a -> c
	// a -> b
	// b -> c
	graph.AddEdge("a", "c")
	graph.AddEdge("a", "b")
	graph.AddEdge("b", "c")

	results, err := graph.TopSort("a")
	if err != nil {
		t.Error(err)
		return
	}
	if results[0] != "c" || results[1] != "b" || results[2] != "a" {
		t.Errorf("Wrong sort order: %v", results)
	}
}

func TestTopSort3(t *testing.T) {
	graph := initGraph()

	// a -> b
	// a -> d
	// d -> c
	// c -> b
	graph.AddEdge("a", "b")
	graph.AddEdge("a", "d")
	graph.AddEdge("d", "c")
	graph.AddEdge("c", "b")

	results, err := graph.TopSort("a")
	if err != nil {
		t.Error(err)
		return
	}
	if len(results) != 4 {
		t.Errorf("Wrong number of results: %v", results)
		return
	}
	expected := [4]string{"b", "c", "d", "a"}
	for i := 0; i < 4; i++ {
		if results[i] != expected[i] {
			t.Errorf("Wrong sort order: %v", results)
			break
		}
	}
}

func TestTopSortCycleError(t *testing.T) {
	graph := initGraph()

	// a -> b
	// b -> a
	graph.AddEdge("a", "b")
	graph.AddEdge("b", "a")

	_, err := graph.TopSort("a")
	if err == nil {
		t.Errorf("Expected cycle error")
		return
	}
	if !strings.Contains(err.Error(), "a -> b -> a") {
		t.Errorf("Error doesn't print cycle: %q", err)
	}
}

func TestTopSortCycleError2(t *testing.T) {
	graph := initGraph()

	// a -> b
	// b -> c
	// c -> a
	graph.AddEdge("a", "b")
	graph.AddEdge("b", "c")
	graph.AddEdge("c", "a")

	_, err := graph.TopSort("a")
	if err == nil {
		t.Errorf("Expected cycle error")
		return
	}
	if !strings.Contains(err.Error(), "a -> b -> c -> a") {
		t.Errorf("Error doesn't print cycle: %q", err)
	}
}

func TestTopSortCycleError3(t *testing.T) {
	graph := initGraph()

	// a -> b
	// b -> c
	// c -> b
	graph.AddEdge("a", "b")
	graph.AddEdge("b", "c")
	graph.AddEdge("c", "b")

	_, err := graph.TopSort("a")
	if err == nil {
		t.Errorf("Expected cycle error")
		return
	}
	if !strings.Contains(err.Error(), "b -> c -> b") {
		t.Errorf("Error doesn't print cycle: %q", err)
	}
}

func TestStableSort(t *testing.T) {
	graph := initGraph()
	graph.AddEdge("a", "b")
	graph.AddEdge("a", "c")
	graph.AddEdge("a", "d")
	graph.AddEdge("b", "x")
	graph.AddEdge("c", "y")
	graph.AddEdge("d", "z")

	expected := []string{"x", "b", "y", "c", "z", "d", "a"}
	actual, err := graph.StableTopSort("a")
	if err != nil {
		t.Error(err)
		return
	}

	deepEqual(t, expected, actual)

	actual2, err := graph.StableTopSort("a")
	if err != nil {
		t.Error(err)
		return
	}

	if !deepEqual(t, actual, actual2) {
		t.Fatal("sort wasn't stable")
	}

}

func TestNonStable(t *testing.T) {
	graph := initGraph()
	graph.AddEdge("a", "b")
	graph.AddEdge("a", "c")
	graph.AddEdge("a", "d")
	graph.AddEdge("b", "x")
	graph.AddEdge("c", "y")
	graph.AddEdge("d", "z")

	a1, _ := graph.TopSort("a")
	a2, _ := graph.TopSort("a")

	if !deepEqual(t, a1, a2) {
		t.Log("not equal, but still fine")
	}
}

func deepEqual(t *testing.T, expected, actual []string) bool {
	if len(actual) != len(expected) {
		t.Logf("actual len %d, expected len %d", len(actual), len(expected))
		return false
	}

	for i := range actual {
		if actual[i] != expected[i] {
			t.Logf("actual %q != expected %q", actual[i], expected[i])
			return false
		}
	}

	return true
}

func initGraph() *Graph[string] {
	graph := NewGraph[string]()
	return graph
}
