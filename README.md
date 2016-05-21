# :four_leaf_clover: Quadtree - Go

> "A quadtree is a tree data structure in which each internal node has exactly four children. Quadtrees are most often used to partition a two-dimensional space by recursively subdividing it into four quadrants or regions. The regions may be square or rectangular, or may have arbitrary shapes. This data structure was named a quadtree by Raphael Finkel and J.L. Bentley in 1974." - Wikipedia (2016)

<br>

<p align="center">
<img src="https://upload.wikimedia.org/wikipedia/commons/8/8b/Point_quadtree.svg">
</p>

<br>

A Quadtree based on [this JavaScript implementation](https://github.com/timohausmann/quadtree-js) from Timo Hausmann. This is in turn based on [this post](http://gamedev.tutsplus.com/tutorials/implementation/quick-tip-use-quadtrees-to-detect-likely-collisions-in-2d-space/).

# Setup

Installation

    go install

Test

    go test

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

    // Methods
    topRight := quadtree.Bounds{
		X:      99,
		Y:      99,
		Width:  0, // A point essentially (0 width and height)
		Height: 0,
	}
	qt.Insert(topRight) // Insert the object into the quadtree
	index = qt.GetIndex(topRight) // Will return 0
```

# To Do

* More unit tests
* Benchmarking, profiling, performance improvements

# License
MIT

Original JavaScript code: Copyright © 2012 Timo Hausmann

Go code: Copyright © 2016 James Milner
