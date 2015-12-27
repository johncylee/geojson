package geojson

import (
	"encoding/json"
	"fmt"
)

// Normally [lon, lat]
type Position []float64

// Generic Geometry Object used to Unmarshal
type Geometry struct {
	Type        string          `json:"type"`
	Coordinates json.RawMessage `json:"coordinates,omitempty"`
	Geometries  []Geometry      `json:"geometries,omitempty"`
}

type Point struct {
	Type        string   `json:"type"`
	Coordinates Position `json:"coordinates"`
}

type MultiPoint struct {
	Type        string     `json:"type"`
	Coordinates []Position `json:"coordinates"`
}

type LineString struct {
	Type        string     `json:"type"`
	Coordinates []Position `json:"coordinates"`
}

type MultiLineString struct {
	Type        string       `json:"type"`
	Coordinates [][]Position `json:"coordinates"`
}

type Polygon struct {
	Type        string       `json:"type"`
	Coordinates [][]Position `json:"coordinates"`
}

type MultiPolygon struct {
	Type        string         `json:"type"`
	Coordinates [][][]Position `json:"coordinates"`
}

type GeometryCollection struct {
	Type       string        `json:"type"`
	Geometries []interface{} `json:"geometries"`
}

// Cast into a specific geometry object.
func Cast(g *Geometry) (interface{}, error) {
	switch g.Type {
	case "Point":
		var coord Position
		err := json.Unmarshal(g.Coordinates, &coord)
		if err != nil {
			return nil, err
		}
		return &Point{
			Type:        g.Type,
			Coordinates: coord,
		}, nil
	case "MultiPoint":
		var coord []Position
		err := json.Unmarshal(g.Coordinates, &coord)
		if err != nil {
			return nil, err
		}
		return &MultiPoint{
			Type:        g.Type,
			Coordinates: coord,
		}, nil
	case "LineString":
		var coord []Position
		err := json.Unmarshal(g.Coordinates, &coord)
		if err != nil {
			return nil, err
		}
		return &LineString{
			Type:        g.Type,
			Coordinates: coord,
		}, nil
	case "MultiLineString":
		var coord [][]Position
		err := json.Unmarshal(g.Coordinates, &coord)
		if err != nil {
			return nil, err
		}
		return &MultiLineString{
			Type:        g.Type,
			Coordinates: coord,
		}, nil
	case "Polygon":
		var coord [][]Position
		err := json.Unmarshal(g.Coordinates, &coord)
		if err != nil {
			return nil, err
		}
		return &Polygon{
			Type:        g.Type,
			Coordinates: coord,
		}, nil
	case "MultiPolygon":
		var coord [][][]Position
		err := json.Unmarshal(g.Coordinates, &coord)
		if err != nil {
			return nil, err
		}
		return &MultiPolygon{
			Type:        g.Type,
			Coordinates: coord,
		}, nil
	case "GeometryCollection":
		geometries := make([]interface{}, 0, len(g.Geometries))
		for _, obj := range g.Geometries {
			o, err := Cast(&obj)
			if err != nil {
				return nil, err
			}
			geometries = append(geometries, o)
		}
		return &GeometryCollection{
			Type:       g.Type,
			Geometries: geometries,
		}, nil
	default:
		return nil, fmt.Errorf("Unknown type of Geometry Object: %s",
			g.Type)
	}
}
