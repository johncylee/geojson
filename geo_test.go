package geojson

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"testing"
)

const (
	SAMPLES = "samples.json"
)

func validateGeoObj(t *testing.T, g *Geometry, o interface{}) error {
	typeName := reflect.TypeOf(o).Elem().Name()
	if typeName != g.Type {
		return fmt.Errorf("Casted and original objects do not match. "+
			"Got: %s, want: %s", typeName, g.Type)
	}
	if g.Type != "GeometryCollection" {
		return nil
	}
	t.Log("Found GeometryCollection, validating geometries")
	c := o.(*GeometryCollection)
	for i, obj := range c.Geometries {
		err := validateGeoObj(t, &g.Geometries[i], obj)
		if err == nil {
			continue
		}
		return fmt.Errorf("Failed at geometries[%d]: %s",
			i, err.Error())
	}
	return nil
}

func TestGeometry(t *testing.T) {
	f, err := os.Open(SAMPLES)
	defer f.Close()
	var geoObjs []Geometry
	dec := json.NewDecoder(f)
	err = dec.Decode(&geoObjs)
	if err != nil {
		t.Fatalf("Failed to load from %s: %s", SAMPLES, err.Error())
	}
	for i, g := range geoObjs {
		t.Logf("Testing type: %s", g.Type)
		o, err := Cast(&g)
		if err != nil {
			t.Errorf("Cast failed at #%d: %s", i, err.Error())
		}
		err = validateGeoObj(t, &g, o)
		if err != nil {
			t.Errorf("Validation failed at %d: %s", i, err.Error())
		}
	}
}
