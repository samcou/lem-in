package lemin

import (
	"regexp"
	"strconv"
	"strings"
)

// filterData takes in an array of strings, data, and returns three values:
// antNbr, allRooms, and links.
// antNbr is an integer representing the number of ants.
// allRooms is an array of strings representing the names of all rooms, including the start and end rooms.
// links is an array of strings representing links between rooms.
func FilterData(data []string) (antNbr int, allRooms, links []string) {
	// Convert the first element of data, representing the number of ants, into an integer.
	antNbr, _ = strconv.Atoi(data[0])

	// Find the index of the line containing "##start".
	startIndex := -1
	for i, line := range data {
		if strings.Contains(line, "##start") {
			startIndex = i
			break
		}
	}

	// Extract the room name associated with "##start".
	var startRoom []string
	if startIndex != -1 {
		temp := data[startIndex+1]
		parts := strings.Split(temp, " ")
		startRoom = append(startRoom, parts[0])
	}

	// Find the index of the line containing "##end".
	endIndex := -1
	for i, line := range data {
		if strings.Contains(line, "##end") {
			endIndex = i
			break
		}
	}

	// Extract the room name associated with "##end".
	var endRoom string
	if endIndex != -1 {
		temp := data[endIndex+1]
		parts := strings.Split(temp, " ")
		endRoom = parts[0]
	}

	// Extract the names of all rooms, excluding the start and end rooms.
	var rooms []string
	for i, line := range data {
		if i == startIndex+1 || i == endIndex+1 {
			continue
		}
		if strings.Contains(line, " ") {
			parts := strings.Split(line, " ")
			rooms = append(rooms, parts[0])
		}
	}

	// Combine the start room, all rooms, and end room into a single array.
	allRooms = append(startRoom, rooms...)
	allRooms = append(allRooms, endRoom)

	// Extract links between rooms using a regular expression.
	regex := regexp.MustCompile(`\b\w+-\w+\b`)
	for _, line := range data {
		matches := regex.FindAllString(line, -1)
		links = append(links, matches...)
	}

	return antNbr, allRooms, links
}
