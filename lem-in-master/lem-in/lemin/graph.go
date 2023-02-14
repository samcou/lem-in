package lemin

//create a graph structure with nodes containing
//id, room, and neighbours

type Room struct {
	Id         int
	Name       string
	Neighbours []*Room
}

type Graph struct {
	Rooms map[int]*Room
}

// newGraph creates a new Graph object and initializes its rooms map
func NewGraph() *Graph {
	return &Graph{
		Rooms: make(map[int]*Room),
	}
}

// addRoom adds a new room to the graph with the given id and name
func (g *Graph) AddRoom(id int, name string) *Room {
	// create a new Room object with the given id and name
	r := &Room{
		Id:   id,
		Name: name,
	}

	// add the room to the graph's rooms map
	g.Rooms[id] = r

	// return the room
	return r
}

// addLink creates a two-way link between two rooms in the graph with the given ids
func (g *Graph) AddLink(id1, id2 int) {
	// retrieve the rooms with the given ids from the graph's rooms map
	r1 := g.Rooms[id1]
	r2 := g.Rooms[id2]

	// add each room to the other's neighbours list to create a two-way link
	r1.Neighbours = append(r1.Neighbours, r2)
	r2.Neighbours = append(r2.Neighbours, r1)
}
