package lemin

// dfs is a recursive function that implements depth-first search
func (g *Graph) dfs(currentRoom, end *Room, visited []bool, path []*Room, paths *[][]*Room) {
	// Mark the current room as visited
	visited[currentRoom.Id] = true

	// Add the current room to the path
	path = append(path, currentRoom)

	// If the current room is the end room, add the path to the list of all paths
	if currentRoom == end {
		pathCopy := make([]*Room, len(path))
		copy(pathCopy, path)
		*paths = append(*paths, pathCopy)
	} else {
		// Recursively search the neighbours of the current room
		for _, neighbour := range currentRoom.Neighbours {
			if !visited[neighbour.Id] {
				g.dfs(neighbour, end, visited, path, paths)
			}
		}
	}

	// Backtrack by removing the current room from the path
	path = path[:len(path)-1]

	// Mark the current room as not visited
	visited[currentRoom.Id] = false
}

// findPaths is a function that finds all paths from a start room to an end room in a graph
func (g *Graph) FindPaths(start, end *Room) [][]*Room {
	// Create a list to keep track of the rooms that have been visited
	visited := make([]bool, len(g.Rooms))

	// Initialize an empty path
	path := []*Room{}

	// Initialize an empty list of paths
	paths := [][]*Room{}

	// Call the depth-first search function to find all paths from the start room to the end room
	g.dfs(start, end, visited, path, &paths)

	// Return the list of all paths
	return paths
}
