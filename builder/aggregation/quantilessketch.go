package aggregation

// QuantilesSketch holds the quantiles struct based on
// Aggregator section in http://druid.apache.org/docs/latest/development/extensions-core/datasketches-quantiles.html
type QuantilesSketch struct {
	Base
	FieldName string `json:"fieldName,omitempty"`
	K         int64  `json:"k,omitempty"`
}

// NewQuantilesSketch create a new instance of QuantilesSketch
func NewQuantilesSketch() *QuantilesSketch {
	t := &QuantilesSketch{}
	t.Base.SetType("quantilesDoublesSketch")
	return t
}

// SetName set name
func (t *QuantilesSketch) SetName(name string) *QuantilesSketch {
	t.Base.SetName(name)
	return t
}

// SetFieldName set fieldName
func (t *QuantilesSketch) SetFieldName(fieldName string) *QuantilesSketch {
	t.FieldName = fieldName
	return t
}

// SetK set quantiles k
func (t *QuantilesSketch) SetK(k int64) *QuantilesSketch {
	t.K = k
	return t
}
