package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type connection struct {
	a 		string
	command string
	b 		string
	t 		string
}

func main() {

	values, connections := readFile("input.txt")

	// sort.Slice(connections, func(i, j int) bool {
	// 	return connections[i].a > connections[j].a
	// })

	// fmt.Println("Connections:", connections)

	sValue := processConnections(values, connections)

	// convert a binary represented string to an integer
	i, _ := strconv.ParseInt(sValue, 2, 64)
	fmt.Println(i)

}

func processConnections(values map[string]int, connections []connection) string {

	completedConnection := make(map[string]bool)
	for _, c := range connections {
		completedConnection[fmt.Sprintf("%v", c)] = false;
	}

	connected := 0

	for connected < len(connections) {

		for _,connection := range connections {
			if !completedConnection[fmt.Sprintf("%v", connection)] {
				if _, okA := values[connection.a]; okA {
					if _, okB := values[connection.b]; okB {
						if connection.command == "AND" {
							values[connection.t] = values[connection.a] & values[connection.b]
						} else if connection.command == "OR" {
							values[connection.t] = values[connection.a] | values[connection.b]
						} else if connection.command == "XOR" {
							values[connection.t] = values[connection.a] ^ values[connection.b]
						}
						completedConnection[fmt.Sprintf("%v", connection)] = true
						connected++
					}
				}
			}
		}
	}

	// find all the values in the map of values which begin with z
	// then print them out in order from the largest z (z9) to the smallest z (z0)
	var zValues []string
	for k := range values {
		if strings.HasPrefix(k, "z") {
			zValues = append(zValues, k)
		}
	}
	sort.Slice(zValues, func(i, j int) bool {
		return zValues[i] > zValues[j]
	})
	sValue := ""
	for _, k := range zValues {
		sValue = sValue + strconv.Itoa(values[k])
	}
	return sValue
}

func readFile(filename string) (map[string]int, []connection) {
	startingValues := make(map[string]int)
	connections := []connection{}

	file, err := os.Open(filename)
	if err != nil {
		fmt.Errorf("error opening file: %w", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		s := strings.Split(line, ":")
		num, _ := strconv.Atoi(strings.TrimSpace(s[1]))
		startingValues[s[0]] = num
	}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		s := strings.Split(line, " ")
		connections = append(connections, connection{a: s[0], command: s[1], b: s[2], t: s[4]})
	}

	if err := scanner.Err(); err != nil {
		fmt.Errorf("error reading file: %w", err)
		os.Exit(1)
	}

	return startingValues, connections

}