package aggregation

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuantilesSketch(t *testing.T) {
	quantilesSketch := NewQuantilesSketch()
	quantilesSketch.SetName("output_name").SetFieldName("metric_name").SetK(123)

	// "omitempty" will ignore boolean=false
	expected := `{"type":"quantilesDoublesSketch", "k":123, "name":"output_name", "fieldName":"metric_name"}`

	quantilesSketchJSON, err := json.Marshal(quantilesSketch)
	assert.Nil(t, err)
	assert.JSONEq(t, expected, string(quantilesSketchJSON))
}
