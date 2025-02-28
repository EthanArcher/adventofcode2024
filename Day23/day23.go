package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {

    connections, err := readConnectionsFromFile("input.txt")
    if err != nil {
        fmt.Println("Error reading connections:", err)
        return
    }
	
    network := buildNetwork(connections)
    uniqueTriples := findUniqueTriples(network)

    acc := countTriplesContainingT(uniqueTriples)
    fmt.Println(acc)

	longestChain := findLongestChain(connections, network)
    alpha := sortAlphabetically(longestChain)
    fmt.Println(strings.Join(alpha, ","))
}

// buildNetwork constructs the network map from the connections.
func buildNetwork(connections [][]string) map[string][]string {
    network := make(map[string][]string)
    for _, c := range connections {
        network[c[0]] = addToNetworkIfNotThere(c[1], network[c[0]])
        network[c[1]] = addToNetworkIfNotThere(c[0], network[c[1]])
    }
    return network
}

// findUniqueTriples finds unique triples in the network.
func findUniqueTriples(network map[string][]string) [][]string {
    uniqueTriples := [][]string{}
    for a, aConnections := range network {
        for _, b := range aConnections {
            for _, c := range network[b] {
                if connectedThrough(c, a, network) != "" {
                    uniqueTriples = addToTriples(a, b, c, uniqueTriples)
                }
            }
        }
    }
    return uniqueTriples
}

// countTriplesContainingT counts the triples that contain a connection starting with 't'.
func countTriplesContainingT(uniqueTriples [][]string) int {
    acc := 0
    for _, ut := range uniqueTriples {
        if containsT(ut) {
            acc++
        }
    }
    return acc
}

// findLongestChain finds the longest chain of connections.
func findLongestChain(connections [][]string, network map[string][]string) []string {
    for c, cons := range network {
        for i := range connections {
            if containsAll(connections[i], cons) {
                connections[i] = append(connections[i], c)
            }
        }
    }

    longestChain := []string{}
    for _, c := range connections {
        if len(c) > len(longestChain) {
            longestChain = c
        }
    }
    return longestChain
}

// sortAlphabetically sorts the connections alphabetically.
func sortAlphabetically(connections []string) []string {
    sort.Strings(connections)
    return connections
}

// containsT checks if any connection starts with 't'.
func containsT(connections []string) bool {
    for _, c := range connections {
        if strings.HasPrefix(c, "t") {
            return true
        }
    }
    return false
}

// containsAll checks if all required connections are present in the connections.
func containsAll(requiredConnections []string, connections []string) bool {
    for _, rc := range requiredConnections {
        if !contains(connections, rc) {
            return false
        }
    }
    return true
}


// addToTriples adds a unique triple to the list of unique triples.
func addToTriples(a string, b string, c string, uniqueTriples [][]string) [][]string {
    combination := []string{a, b, c}
    sort.Strings(combination)
    for _, t := range uniqueTriples {
        if contains(t, combination[0]) && contains(t, combination[1]) && contains(t, combination[2]) {
            return uniqueTriples
        }
    }
    return append(uniqueTriples, combination)
}

// contains checks if a string is present in a slice.
func contains(slice []string, item string) bool {
    for _, v := range slice {
        if v == item {
            return true
        }
    }
    return false
}

// connectedThrough checks if there is a connection through a given node.
func connectedThrough(b string, a string, network map[string][]string) string {
    for _, c := range network[b] {
        if c == a {
            return c
        }
    }
    return ""
}

// addToNetworkIfNotThere adds a connection to the network if it is not already present.
func addToNetworkIfNotThere(connection string, network []string) []string {
    if !contains(network, connection) {
        return append(network, connection)
    }
    return network
}

// readConnectionsFromFile reads connections from a file.
func readConnectionsFromFile(filename string) ([][]string, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, fmt.Errorf("error opening file: %w", err)
    }
    defer file.Close()

    var connections [][]string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        cons := strings.Split(scanner.Text(), "-")
        connections = append(connections, cons)
    }

    if err := scanner.Err(); err != nil {
        return nil, fmt.Errorf("error reading file: %w", err)
    }

    return connections, nil
}