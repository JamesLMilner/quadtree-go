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

func (b *Bounds) isPoint() bool {

	if b.Width == 0 && b.Height == 0 {
		return true
	}

	return false

}

// Checks if a Bounds object intersects with another Bounds
func (b *Bounds) Intersects(a Bounds) bool {

	aMaxX := a.X + a.Width
	aMaxY := a.Y + a.Height
	bMaxX := b.X + b.Width
	bMaxY := b.Y + b.Height

	// a is left of b
	if aMaxX < b.X {
		return false
	}

	// a is right of b
	if a.X > bMaxX {
		return false
	}

	// a is above b
	if aMaxY < b.Y {
		return false
	}

	// a is below b
	if a.Y > bMaxY {
		return false
	}

	// The two overlap
	return true

}

// Retrieve the total number of sub-Quadtrees in a Quadtree
func (qt *Quadtree) TotalNodes() int {

	total := 0

	if len(qt.Nodes) > 0 {
		for i := 0; i < len(qt.Nodes); i++ {
			total += 1
			total += qt.Nodes[i].TotalNodes()
		}
	}

	return total

}

// Split the node into 4 subnodes
func (qt *Quadtree) Split() {

	if len(qt.Nodes) == 4 {
		return
	}

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
	if len(qt.Nodes) > 0 == true {

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
		if len(qt.Nodes) > 0 == false {
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

	// Array with all detected objects
	returnObjects := qt.Objects

	//if we have subnodes ...
	if len(qt.Nodes) > 0 {

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

// Retrieve - Return all points that collide
func (qt *Quadtree) RetrievePoints(find Bounds) []Bounds {

	var foundPoints []Bounds
	potentials := qt.Retrieve(find)
	for o := 0; o < len(potentials); o++ {

		// X and Ys are the same and it has no Width and Height (Point)
		xyMatch := potentials[o].X == float64(find.X) && potentials[o].Y == float64(find.Y)
		if xyMatch && potentials[o].isPoint() {
			foundPoints = append(foundPoints, find)
		}
	}

	return foundPoints

}

func (qt *Quadtree) RetrieveIntersections(find Bounds) []Bounds {

	var foundIntersections []Bounds

	potentials := qt.Retrieve(find)
	for o := 0; o < len(potentials); o++ {
		if potentials[o].Intersects(find) {
			foundIntersections = append(foundIntersections, potentials[o])
		}
	}

	return foundIntersections

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
	qt.Total = 0

}
