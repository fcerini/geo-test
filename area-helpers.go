package main

import (
	"math"
)

// Helpers para calculos basicos

// Distancia en metros entre el punto p y el segmento v-w
func DistToSegment(p, v, w Point) float64 {
	grados := math.Sqrt(distToSegmentSquared(p, v, w))
	return grados * 111000 // +- un grado en metros
}

// Distancia plana en metros (aprox.) entre 2 Lat Long
func FlatDistance(lat1, lng1, lat2, lng2 float64) float64 {
	a := (lat1 - lat2) * flatDistLng(lat1)
	b := (lng1 - lng2) * flatDistLat(lat1)
	return math.Sqrt(a*a + b*b)
}

func sqr(x float64) float64 {
	return x * x
}

func dist2(v, w Point) float64 {
	return sqr(v.X-w.X) + sqr(v.Y-w.Y)
}

func distToSegmentSquared(p, v, w Point) float64 {
	l2 := dist2(v, w)
	if l2 == 0 {
		return dist2(p, v)
	}
	t := ((p.X-v.X)*(w.X-v.X) + (p.Y-v.Y)*(w.Y-v.Y)) / l2
	if t < 0 {
		return dist2(p, v)
	}
	if t > 1 {
		return dist2(p, w)
	}
	aux := Point{X: v.X + t*(w.X-v.X), Y: v.Y + t*(w.Y-v.Y)}
	return dist2(p, aux)
}

func flatDistLng(lat float64) float64 {
	return 0.0003121092*math.Pow(lat, 4) +
		0.0101182384*math.Pow(lat, 3) -
		17.2385140059*lat*lat +
		5.5485277537*lat + 111301.967182595
}

func flatDistLat(lat float64) float64 {
	return -0.000000487305676*math.Pow(lat, 4) -
		0.0033668574*math.Pow(lat, 3) +
		0.4601181791*lat*lat -
		1.4558127346*lat + 110579.25662316
}
