package main

import (
	"testing"
)

func TestInArea(t *testing.T) {
	// Plaza San Martin
	area := NewArea("POLYGON ((-58.375252 -34.592306, -58.377019 -34.594361, -58.377732 -34.593935, -58.378679 -34.594031, -58.378607 -34.595325, -58.377828 -34.595235, -58.376067 -34.596529, -58.375294 -34.596385, -58.373371 -34.593456, -58.375252 -34.592306))",
		100)

	//Benchmark
	for i := 0; i < 1000000; i++ {

		//Puntos adentro
		if !area.InArea(Point{X: -58.3761802, Y: -34.5950633}) ||
			!area.InArea(Point{X: -58.3746935, Y: -34.595550}) || // muy cerca
			!area.InArea(Point{X: -58.3772645, Y: -34.593460}) { //tipo 70 metros
			t.Error("Punto en el area lo marca afuera")
		}

		//Puntos Afuera
		if area.InArea(Point{X: -58.37313, Y: -34.59616}) ||
			area.InArea(Point{X: -58.3966175, Y: -34.606401}) {
			t.Error("Punto FUERA del area lo marca adentro")
		}
	}
}

func TestDist(t *testing.T) {
	d := DistToSegment(
		Point{
			X: -58.374699330856885,
			Y: -34.59556314353189,
		},
		Point{
			X: -58.37522279680157,
			Y: -34.595603737796054,
		},
		Point{
			X: -58.37482830384703,
			Y: -34.595082277056385,
		})
	if d > 50 {
		t.Error("DistToSegment deberia dar tipo 40 ", d)
	}

}
