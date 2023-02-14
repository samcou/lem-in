package main

import (
	"bufio"
	"fmt"
	"lemin/lemin"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	data, err := readData()
	if err != nil {
		fmt.Println(err)
		return
	}
	printOutput(data)
}

func readData() ([]string, error) {
	if len(os.Args) != 2 {
		return nil, fmt.Errorf("ERROR: Incorrect number of arguments.\ninput format: go run . example00.txt")
	}

	path := os.Args[1]
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("ERROR: Failed to open the file: %v", err)
	}
	defer file.Close()

	var data []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("ERROR: Failed to read the file: %v", err)
	}

	return data, scanner.Err()
}

func printOutput(data []string) {
	antNbr, allRooms, allLinks := lemin.FilterData(data)
	if antNbr <= 0 {
		fmt.Println("ERROR: Invalid number of ants. Must be > 0")
		return
	}

	graph := lemin.NewGraph()

	//add rooms to the graph
	roomIDs := make(map[string]int)
	for id, room := range allRooms {
		roomIDs[room] = id
		graph.AddRoom(id, room)
	}

	//add links to the graph
	for _, link := range allLinks {
		parts := strings.Split(link, "-")
		id1 := roomIDs[parts[0]]
		id2 := roomIDs[parts[1]]
		graph.AddLink(id1, id2)
	}

	//assign start and end points
	startRoom := graph.Rooms[0]
	endRoom := graph.Rooms[len(graph.Rooms)-1]

	paths := graph.FindPaths(startRoom, endRoom)
	if len(paths) == 0 {
		fmt.Println("ERROR: no path found or text file is formatted incorrectly")
		return
	}

	validPaths := lemin.FindCompatiblePaths(paths)
	bestPath := lemin.PathAssign(paths, validPaths, antNbr)

	// sanitize the input path
	path := filepath.Clean(os.Args[1])

	// check if the input path is absolute or relative to the current working directory
	if !filepath.IsAbs(path) {
		wd, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
			return
		}
		path = filepath.Join(wd, path)
	}

	bytes, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		return
	}

	content := string(bytes)
	fmt.Println(content)
	fmt.Println()
	lemin.PrintAntSteps(paths, bestPath)
}
