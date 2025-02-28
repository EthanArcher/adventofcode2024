package main

import (
	"reflect"
	"testing"
)

func TestAddConnectionToNetworks(t *testing.T) {
	tests := []struct {
		name        string
		connection  []string
		initialLans [][]string
		expectedLans [][]string
	}{
		{
			name:        "Add to empty lans",
			connection:  []string{"A", "B"},
			initialLans: [][]string{},
			expectedLans: [][]string{{"A", "B"}},
		},
		{
			name:        "Add to existing lan",
			connection:  []string{"B", "C"},
			initialLans: [][]string{{"A", "B"}},
			expectedLans: [][]string{{"A", "B", "C"}},
		},
		{
			name:        "Add new lan",
			connection:  []string{"D", "E"},
			initialLans: [][]string{{"A", "B"}},
			expectedLans: [][]string{{"A", "B"}, {"D", "E"}},
		},
		{
			name:        "Merge lans",
			connection:  []string{"C", "D"},
			initialLans: [][]string{{"A", "B", "C"}, {"D", "E"}},
			expectedLans: [][]string{{"A", "B", "C", "D"}, {"D", "E", "C"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lans := tt.initialLans
			addConnectionToNetworks(tt.connection, &lans)
			if !reflect.DeepEqual(lans, tt.expectedLans) {
				t.Errorf("got %v, want %v", lans, tt.expectedLans)
			}
		})
	}
}
