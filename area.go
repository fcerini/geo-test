package main

import (
	"fmt"
	"strconv"
	"strings"
)

type AreaType int

const (
	POLYGON AreaType = iota
	CIRCLE
	LINESTRING
)

type Point struct {
	X float64
	Y float64
}

// Area: circulo, poligono o linea
type Area struct {
	Points []Point
	Lat    float64
	Lon    float64
	Buffer float64
	Type   AreaType
}

func NewCircle(lat float64, lon float64, buffer float64) *Area {
	area := &Area{
		Points: []Point{},
		Lat:    lat,
		Lon:    lon,
		Buffer: buffer,
		Type:   CIRCLE,
	}

	return area
}

func NewArea(wkt string, buffer float64) *Area {
	area := &Area{Buffer: buffer}

	wkt = strings.ToUpper(wkt)
	if strings.Contains(wkt, "POLYGON") {
		area.Type = POLYGON
	} else if strings.Contains(wkt, "LINESTRING") {
		area.Type = LINESTRING
	} else {
		return nil
	}

	wkt = strings.ReplaceAll(wkt, "  ", " ")
	wkt = strings.ReplaceAll(wkt, "POLYGON", "")
	wkt = strings.ReplaceAll(wkt, "LINESTRING", "")
	wkt = strings.ReplaceAll(wkt, " (", "")
	wkt = strings.ReplaceAll(wkt, "(", "")
	wkt = strings.ReplaceAll(wkt, ")", "")
	wkt = strings.ReplaceAll(wkt, ", ", ",")

	lonLats := strings.Split(wkt, ",")
	// OJO funciona para GEOMETRY X,Y o sea Lon,Lat
	// NO para GEOGRAPHY que es Lat,Lon
	for _, lonLat := range lonLats {
		xy := strings.Split(lonLat, " ")

		x, _ := strconv.ParseFloat(xy[0], 64)
		y, _ := strconv.ParseFloat(xy[1], 64)

		area.Points = append(area.Points, Point{
			X: x,
			Y: y,
		})
	}

	return area
}

func (a *Area) InArea(test Point) bool {
	var result bool

	switch a.Type {
	case POLYGON:
		result = a.inPolygon(test)
	case CIRCLE:
		result = a.inCircle(test)
	default:
		panic(fmt.Sprintf("unexpected main.AreaType: %#v", a.Type))
	}
	return result
}

func (a *Area) inPolygon(test Point) bool {
	nverts := len(a.Points)
	inside := false
	j := nverts - 1
	for i := 0; i < nverts; i++ {
		if (a.Points[i].Y > test.Y) != (a.Points[j].Y > test.Y) &&
			(test.X < (a.Points[j].X-a.Points[i].X)*(test.Y-a.Points[i].Y)/(a.Points[j].Y-a.Points[i].Y)+a.Points[i].X) {
			inside = !inside
		}
		j = i
	}

	// si no esta dentro del poligono, me fijo si esta cerca de alguna de las lineas del borde
	if a.Buffer > 0 &&
		!inside {
		dist := 999.0
		j = nverts - 1
		for i := 0; i < nverts; i++ {
			daux := DistToSegment(test, a.Points[i], a.Points[j])
			if daux < dist {
				dist = daux
			}
			j = i
		}

		if dist <= a.Buffer {
			inside = true
		}
	}

	return inside
}
func (a *Area) inCircle(test Point) bool {
	result := false

	distance := FlatDistance(test.Y, test.X, a.Lat, a.Lon)
	if distance <= a.Buffer {
		result = true
	}

	return result
}
