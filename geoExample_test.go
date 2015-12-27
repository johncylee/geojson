package geojson

import (
	"encoding/json"
	"fmt"
)

func ExampleCast() {
	jsonGeoObj := []byte(`{
  "type": "Point",
  "coordinates": [121.562321, 25.033908],
  "_comment": "台北101"
}`)
	var g Geometry
	err := json.Unmarshal(jsonGeoObj, &g)
	if err != nil {
		panic(err)
	}
	v, err := Cast(&g)
	if err != nil {
		panic(err)
	}
	switch g.Type {
	case "Point":
		o := v.(*Point)
		// May access fields with specific types now
		fmt.Printf("Type: %s, coordinates: [%f, %f]",
			o.Type, o.Coordinates[0], o.Coordinates[1])
	default:
		// can add more cases
		panic(fmt.Errorf("Unsupported type: %s", g.Type))
	}
	// Output:
	// Type: Point, coordinates: [121.562321, 25.033908]
}
