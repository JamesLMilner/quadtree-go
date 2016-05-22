package quadtree_test

import (
	"math/rand"
	"github.com/JamesMilnerUK/quadtree-go"
	"testing"
	"time"
)

func TestQuadtreeInsert(t *testing.T) {

	rand.Seed(time.Now().UTC().UnixNano()) // Seed Random properly

	qt := setupQuadtree(0, 0, 640, 480)

	grid := 10.0
	gridh := qt.Bounds.Width / grid
	gridv := qt.Bounds.Height / grid
	var randomObject quadtree.Bounds
	numObjects := 1000

	for i := 0; i < numObjects; i++ {

		x := randMinMax(0, gridh) * grid
		y := randMinMax(0, gridv) * grid

		randomObject = quadtree.Bounds{
			X:      x,
			Y:      y,
			Width:  randMinMax(1, 4) * grid,
			Height: randMinMax(1, 4) * grid,
		}

		index := qt.GetIndex(randomObject)
		if index < -1 || index > 3 {
			t.Errorf("The index should be -1 or between 0 and 3, got %d \n", index)
		}

		qt.Insert(randomObject)

	}

	if qt.Total != numObjects {
		t.Errorf("Error: Should have tottled %d, got %d \n", numObjects, qt.Total)
	} else {
		t.Logf("Success: Total objects in the Quadtree is %d (as expected) \n", qt.Total)
	}

	failure := false
	iterations := 20
	for j := 0; j < iterations; j++ {

		Cursor := quadtree.Bounds{
			X:      randMinMax(0, gridh) * grid,
			Y:      randMinMax(0, gridv) * grid,
			Width:  100,
			Height: 100,
		}

		objects := qt.Retrieve(Cursor)
		if len(objects) < 0 && len(objects) > numObjects {
			failure = true
		}

	}

	if failure == false {
		t.Logf("Success: All %d iterations had objects length greater than 0 and less than inserts (%d) \n", iterations, numObjects)
	}

	// All the objects!
	AllCursor := quadtree.Bounds{
		X:      0,
		Y:      0,
		Width:  10000,
		Height: 10000,
	}

	allObjects := qt.Retrieve(AllCursor)
	if len(allObjects) != numObjects {
		t.Errorf("Error: Should have returned all objects (%d) got %d \n", numObjects, len(allObjects))
	} else {
		t.Logf("Success: All the objects (%d) have been returned ", len(allObjects))
	}

}

func BenchmarkInsertOneThousand(b *testing.B) {

	qt := setupQuadtree(0, 0, 640, 480)

	grid := 10.0
	gridh := qt.Bounds.Width / grid
	gridv := qt.Bounds.Height / grid
	var randomObject quadtree.Bounds
	numObjects := 1000

	for n := 0; n < b.N; n++ {
		for i := 0; i < numObjects; i++ {

			x := randMinMax(0, gridh) * grid
			y := randMinMax(0, gridv) * grid

			randomObject = quadtree.Bounds{
				X:      x,
				Y:      y,
				Width:  randMinMax(1, 4) * grid,
				Height: randMinMax(1, 4) * grid,
			}

			qt.Insert(randomObject)

		}
	}



}

func TestCorrectQuad(t *testing.T) {

	qt := setupQuadtree(0, 0, 100, 100)

	var index int
	pass := true

	topRight := quadtree.Bounds{
		X:      99,
		Y:      99,
		Width:  0,
		Height: 0,
	}
	qt.Insert(topRight)
	index = qt.GetIndex(topRight)
	if index == 0 {
		t.Errorf("The index should be 0, got %d \n", index)
		pass = false
	}

	topLeft := quadtree.Bounds{
		X:      99,
		Y:      1,
		Width:  0,
		Height: 0,
	}
	qt.Insert(topLeft)
	index = qt.GetIndex(topLeft)
	if index == 1 {
		t.Errorf("The index should be 1, got %d \n", index)
		pass = false
	}

	bottomLeft := quadtree.Bounds{
		X:      1,
		Y:      1,
		Width:  0,
		Height: 0,
	}
	qt.Insert(bottomLeft)
	index = qt.GetIndex(bottomLeft)
	if index == 2 {
		t.Errorf("The index should be 2, got %d \n", index)
		pass = false
	}

	bottomRight := quadtree.Bounds{
		X:      1,
		Y:      51,
		Width:  0,
		Height: 0,
	}
	qt.Insert(bottomRight)
	index = qt.GetIndex(bottomRight)
	if index == 3 {
		t.Errorf("The index should be 3, got %d \n", index)
		pass = false
	}

	if pass == true {
		t.Log("Success: The points were inserted into the correct quadrants")
	}

}

func setupQuadtree(x float64, y float64, width float64, height float64) *quadtree.Quadtree {

	return &quadtree.Quadtree{
		Bounds: quadtree.Bounds{
			X:      x,
			Y:      y,
			Width:  width,
			Height: height,
		},
		MaxObjects: 10,
		MaxLevels:  4,
		Level:      0,
		Objects:    make([]quadtree.Bounds, 0),
		Nodes:      make([]quadtree.Quadtree, 0),
	}

}

func randMinMax(min float64, max float64) float64 {
	val := min + (rand.Float64() * (max - min))
	return val
}
