package lemin

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// find most optimal paths which can be simultaneously traversed
var requiredSteps int

// findCompatiblePaths is a function that takes a 2D array of paths and returns a 2D array of compatible paths.
func FindCompatiblePaths(paths [][]*Room) [][]int {
	// Initialize a 2D slice to store the compatible paths.
	var compatiblePaths [][]int
	// Loop through each path in the array, and compare it to every subsequent path in the array.
	for i, path1 := range paths {
		// Add the index of the current path to the compatiblePaths slice as a new array containing only that index.
		compatiblePaths = append(compatiblePaths, []int{i})
		// Create a map to keep track of the rooms in the current path.
		roomMap := make(map[int]struct{})
		// Loop through each room in the current path and add it to the roomMap.
		for _, room := range path1[1 : len(path1)-1] {
			roomMap[room.Id] = struct{}{}
		}
		// Loop through each subsequent path and compare it to the current path.
		for j, path2 := range paths[i+1:] {
			// Assume that the two paths are compatible.
			isCompatible := true
			// Loop through each room in the current path and check if it appears in the roomMap of the other path.
			for _, room := range path2[1 : len(path2)-1] {
				if _, ok := roomMap[room.Id]; ok {
					// If a room appears in both paths, the paths are not compatible.
					isCompatible = false
					break
				}
			}
			// If the paths are compatible, add the index of the other path to the compatiblePaths slice for the current path.
			if isCompatible {
				compatiblePaths[i] = append(compatiblePaths[i], i+1+j)
				// Loop through each room in the other path and add it to the roomMap.
				for _, room := range path2[1 : len(path2)-1] {
					roomMap[room.Id] = struct{}{}
				}
			}
		}
	}
	// Return the compatiblePaths slice.
	return compatiblePaths
}

// pathAssign is a function that takes the 2D array of paths and the compatible paths, along with the number of ants, and assigns a path to each ant.
func PathAssign(paths [][]*Room, validPaths [][]int, antNbr int) []string {
	// Initialize variables to keep track of the best assigned path and its maximum step length.
	var bestAssignedPath []string
	bestMaxStepLength := math.MaxInt32
	// Loop through each valid path.
	for _, validPath := range validPaths {
		// Initialize a slice to store the step lengths of each path in the current valid path.
		var stepLength []int
		// Initialize a slice to store the assigned path for each ant.
		var assignedPath []string
		// Loop through each index in the current valid path and add the step length of the corresponding path to the stepLength slice.
		for _, pathIndex := range validPath {
			path := paths[pathIndex]
			stepLength = append(stepLength, len(path)-1)
		}
		// Loop through each ant.
		for i := 1; i <= antNbr; i++ {
			// Find the path in the valid path with the shortest step length and assign the ant to that path
			minStepsIndex := 0
			for j, steps := range stepLength {
				if steps <= stepLength[minStepsIndex] {
					minStepsIndex = j
				}
			}
			assignedPath = append(assignedPath, fmt.Sprintf("%d-%d", i, validPath[minStepsIndex]))
			stepLength[minStepsIndex]++
		}
		// Calculate the maximum step length in the assigned path.
		maxStepLength := 0
		for _, steps := range stepLength {
			if steps > maxStepLength {
				maxStepLength = steps
			}
		}
		// If the maximum step length in the assigned path is less than the best maximum step length so far, update the bestAssignedPath and bestMaxStepLength.
		if maxStepLength < bestMaxStepLength {
			bestAssignedPath = assignedPath
			bestMaxStepLength = maxStepLength
		}
	}
	// Store the required number of steps as the best maximum step length.
	requiredSteps = bestMaxStepLength
	// Return the best assigned path.
	return bestAssignedPath

}

// printAntSteps is a function that takes the filtered paths and the assigned paths and prints the steps taken by each ant.
func PrintAntSteps(filteredPaths [][]*Room, pathStrings []string) {
	// Initialize a 2D slice to store the steps taken by each ant in order.
	var antSteps [][]string
	// Calculate the number of turns required to complete the path.
	arrayLen := requiredSteps - 1
	// Initialize a slice to store the steps taken by each ant in order.
	orderedSteps := make([][]string, arrayLen)
	// Loop through each assigned path.
	for _, antPath := range pathStrings {
		// Initialize a slice to store the steps taken by the current ant.
		var steps []string
		// Split the antPath string into its ant number and path index components.
		parts := strings.SplitN(antPath, "-", 2)
		antStr := parts[0]
		antPath, _ := strconv.Atoi(string(parts[1]))
		// Loop through each room in the path and add a step string to the steps slice for each room.
		for i := 1; i < len(filteredPaths[antPath]); i++ {
			path := filteredPaths[antPath][i]
			temp := "L" + antStr + "-" + path.Name
			steps = append(steps, temp)
		}
		// Add the steps slice to the antSteps slice.
		antSteps = append(antSteps, steps)
	}
	// Loop through each step in each ant's path and add it to the orderedSteps slice in order.
	for i := 0; i < len(antSteps); i++ {
		slice := antSteps[i]
		var row int
		for j := 0; j < len(slice); j++ {
			temp := slice[j]
			if j == 0 {
				// Split the step string to get the room name and use getRow to find the first row in orderedSteps that does not contain the room name.
				parts := strings.SplitN(temp, "-", 2)
				row = getRow(orderedSteps, "-"+parts[1])
			}
			// Add the step string to the row in orderedSteps.
			orderedSteps[j+row] = append(orderedSteps[j+row], temp)
		}
		row = 0
	}
	// Loop through each step in the orderedSteps slice and print it.
	for i, printline := range orderedSteps {
		fmt.Println(strings.Trim(fmt.Sprint(printline), "[]"))
		if i == len(orderedSteps)-1 {
			fmt.Println()
			fmt.Printf("Number of turns: %v\n", i+1)
		}
	}
}

// getRow is a helper function that takes a 2D slice and a value to search for, and returns the index of the first row in the slice that does not contain the value.
func getRow(tocheck [][]string, value string) int {
	// Loop through each row in the slice.
	for i, row := range tocheck {
		found := false
		// Loop through each item in the current row and check if it contains the value.
		for _, item := range row {
			if strings.Contains(item, value) {
				found = true
				break
			}
		}
		// If the value is not found in the current row, return its index.
		if !found {
			return i
		}
	}
	// If the value is found in every row, return 0.
	return 0
}
