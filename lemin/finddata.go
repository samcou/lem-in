package lemin

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadData() ([]string, error) {
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

func PrintOutput(data []string) {
	antNbr, allRooms, allLinks := FilterData(data)
	if antNbr <= 0 {
		fmt.Println("ERROR: Invalid data format. Invalid number of ants. Must be > 0")
		return
	}

	graph := NewGraph()

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
		fmt.Println("ERROR: Invalid data format. No path found or text file is formatted incorrectly")
		return
	}

	validPaths := FindCompatiblePaths(paths)
	bestPath := PathAssign(paths, validPaths, antNbr)

	path := os.Args[1]
	bytes, err := os.ReadFile(path)

	if err != nil {
		fmt.Println(err)
		return
	}

	content := string(bytes)
	fmt.Println(content)
	fmt.Println()
	PrintAntSteps(paths, bestPath)
}
