# :four_leaf_clover: Quadtree (Go)

[![godoc reference](godoc.png)](https://godoc.org/github.com/JamesMilnerUK/quadtree-go)

A quadtree implementation in Go, featuring insertion and retrieval of bounding boxes and points.

# What is it?

> "A quadtree is a tree data structure in which each internal node has exactly four children. Quadtrees are most often used to partition a two-dimensional space by recursively subdividing it into four quadrants or regions. The regions may be square or rectangular, or may have arbitrary shapes. This data structure was named a quadtree by Raphael Finkel and J.L. Bentley in 1974." - Wikipedia (2016)

<br>

<p align="center">
<img src="https://upload.wikimedia.org/wikipedia/commons/8/8b/Point_quadtree.svg">
</p>

<br>


# Setup

Usage
```go
    qt := &quadtree.Quadtree{
		Bounds: quadtree.Bounds{
			X:      0,
			Y:      0,
			Width:  100,
			Height: 100,
		},
		MaxObjects: 10,
		MaxLevels:  4,
		Level:      0,
		Objects:    make([]quadtree.Bounds, 0),
		Nodes:      make([]quadtree.Quadtree, 0),
	}

    // Insert some boxes
    qt.Insert(Bounds{
		X:      1,
		Y:      1,
		Width:  10,
		Height: 10,
	})
	qt.Insert(Bounds{
		X:      5,
		Y:      5,
		Width:  10,
		Height: 10,
	})
	qt.Insert(Bounds{
		X:      10,
		Y:      10,
		Width:  10,
		Height: 10,
	})
	qt.Insert(Bounds{
		X:      15,
		Y:      15,
		Width:  10,
		Height: 10,
	})

	//Get all the intersections
	intersections := qt.RetrieveIntersections(Bounds{
		X:      5,
		Y:      5,
		Width:  2.5,
		Height: 2.5,
	})

	// Clear the Quadtree
	qt.Clear()


```
# Acknowledgements

This package is based on Timo Hausmann [JavaScript Quadtree implementation](https://github.com/timohausmann/quadtree-js). This is in turn based on [this post](http://gamedev.tutsplus.com/tutorials/implementation/quick-tip-use-quadtrees-to-detect-likely-collisions-in-2d-space/).

# License
MIT

Original JavaScript code: Copyright © 2012 Timo Hausmann

Go code: Copyright © 2016 James Milner
