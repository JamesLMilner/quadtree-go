package quadtree

// Quadtree - The quadtree data structure
type Quadtree struct {
	Bounds     Bounds
	MaxObjects int // Maximum objects a node can hold before splitting into 4 subnodes
	MaxLevels  int // Total max levels inside root Quadtree
	Level      int // Depth level, required for subnodes
	Objects    []Bounds
	Nodes      []Quadtree
	Total      int
}

// Bounds - A bounding box with a x y origin and width and height
type Bounds struct {
	X      float64
	Y      float64
	Width  float64
	Height float64
}

// Split the node into 4 subnodes
func (qt *Quadtree) Split() {

	nextLevel := qt.Level + 1
	subWidth := qt.Bounds.Width / 2
	subHeight := qt.Bounds.Height / 2
	x := qt.Bounds.X
	y := qt.Bounds.Y

	//top right node (0)
	qt.Nodes = append(qt.Nodes, Quadtree{
		Bounds: Bounds{
			X:      x + subWidth,
			Y:      y,
			Width:  subWidth,
			Height: subHeight,
		},
		MaxObjects: qt.MaxObjects,
		MaxLevels:  qt.MaxLevels,
		Level:      nextLevel,
		Objects:    make([]Bounds, 0),
		Nodes:      make([]Quadtree, 0, 4),
	})

	//top left node (1)
	qt.Nodes = append(qt.Nodes, Quadtree{
		Bounds: Bounds{
			X:      x,
			Y:      y,
			Width:  subWidth,
			Height: subHeight,
		},
		MaxObjects: qt.MaxObjects,
		MaxLevels:  qt.MaxLevels,
		Level:      nextLevel,
		Objects:    make([]Bounds, 0),
		Nodes:      make([]Quadtree, 0, 4),
	})

	//bottom left node (2)
	qt.Nodes = append(qt.Nodes, Quadtree{
		Bounds: Bounds{
			X:      x,
			Y:      y + subHeight,
			Width:  subWidth,
			Height: subHeight,
		},
		MaxObjects: qt.MaxObjects,
		MaxLevels:  qt.MaxLevels,
		Level:      nextLevel,
		Objects:    make([]Bounds, 0),
		Nodes:      make([]Quadtree, 0, 4),
	})

	//bottom right node (3)
	qt.Nodes = append(qt.Nodes, Quadtree{
		Bounds: Bounds{
			X:      x + subWidth,
			Y:      y + subHeight,
			Width:  subWidth,
			Height: subHeight,
		},
		MaxObjects: qt.MaxObjects,
		MaxLevels:  qt.MaxLevels,
		Level:      nextLevel,
		Objects:    make([]Bounds, 0),
		Nodes:      make([]Quadtree, 0, 4),
	})

}

// GetIndex - Determine which node the object belongs to
func (qt *Quadtree) GetIndex(pRect Bounds) int {

	index := -1 // index of the subnode (0-3), or -1 if pRect cannot completely fit within a subnode and is part of the parent node

	verticalMidpoint := qt.Bounds.X + (qt.Bounds.Width / 2)
	horizontalMidpoint := qt.Bounds.Y + (qt.Bounds.Height / 2)

	//pRect can completely fit within the top quadrants
	topQuadrant := (pRect.Y < horizontalMidpoint) && (pRect.Y+pRect.Height < horizontalMidpoint)

	//pRect can completely fit within the bottom quadrants
	bottomQuadrant := (pRect.Y > horizontalMidpoint)

	//pRect can completely fit within the left quadrants
	if (pRect.X < verticalMidpoint) && (pRect.X+pRect.Width < verticalMidpoint) {

		if topQuadrant {
			index = 1
		} else if bottomQuadrant {
			index = 2
		}

	} else if pRect.X > verticalMidpoint {
		//pRect can completely fit within the right quadrants

		if topQuadrant {
			index = 0
		} else if bottomQuadrant {
			index = 3
		}

	}

	return index

}

// Insert - Insert the object into the node. If the node exceeds the capacity,
// it will split and add all objects to their corresponding subnodes.
func (qt *Quadtree) Insert(pRect Bounds) {

	qt.Total++

	i := 0
	var index int

	// If we have subnodes within the Quadtree
	if len(qt.Nodes)-1 > 0 == true {

		index = qt.GetIndex(pRect)

		if index != -1 {
			qt.Nodes[index].Insert(pRect)
			return
		}
	}

	// If we don't subnodes within the Quadtree
	qt.Objects = append(qt.Objects, pRect)

	// If total objects is greater than max objects and level is less than max levels
	if (len(qt.Objects) > qt.MaxObjects) && (qt.Level < qt.MaxLevels) {

		// Split if we don't already have subnodes
		if len(qt.Nodes)-1 > 0 == false {
			qt.Split()
		}

		// Add all objects to there corresponding subNodes
		for i < len(qt.Objects) {

			index = qt.GetIndex(qt.Objects[i])

			if index != -1 {

				splice := qt.Objects[i]                                  // Get the object out of the slice
				qt.Objects = append(qt.Objects[:i], qt.Objects[i+1:]...) // Remove the object from the slice

				qt.Nodes[index].Insert(splice)

			} else {

				i++

			}

		}

	}

}

// Retrieve - Return all objects that could collide with the given object
func (qt *Quadtree) Retrieve(pRect Bounds) []Bounds {

	index := qt.GetIndex(pRect)

	returnObjects := qt.Objects // Array with all detected objects

	//if we have subnodes ...
	if len(qt.Nodes)-1 > 0 == true {

		//if pRect fits into a subnode ..
		if index != -1 {
			returnObjects = append(returnObjects, qt.Nodes[index].Retrieve(pRect)...)

		} else {

			//if pRect does not fit into a subnode, check it against all subnodes
			for i := 0; i < len(qt.Nodes); i++ {
				returnObjects = append(returnObjects, qt.Nodes[i].Retrieve(pRect)...)
			}

		}
	}

	return returnObjects

}

//Clear - Clear the Quadtree
func (qt *Quadtree) Clear() {

	qt.Objects = []Bounds{}

	if len(qt.Nodes)-1 > 0 {
		for i := 0; i < len(qt.Nodes); i++ {
			qt.Nodes[i].Clear()
		}
	}

	qt.Nodes = []Quadtree{}

}
