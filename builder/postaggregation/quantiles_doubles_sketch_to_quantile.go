package postaggregation

// QuantilesDoublesSketchToQuantile struct based on
// PostAggregator section in http://druid.apache.org/docs/latest/development/extensions-core/datasketches-quantiles.html
type QuantilesDoublesSketchToQuantile struct {
	Base
	Fraction  float64                        			   `json:"fraction,omitempty"`
	Field     *QuantilesDoublesSketchToQuantileField `json:"field,omitempty"`
}

// QuantilesDoublesSketchToQuantileField struct for Field in QuantilesDoublesSketchToQuantile
type QuantilesDoublesSketchToQuantileField struct {
	Type      string `json:"type,omitempty"`
	FieldName string `json:"fieldName,omitempty"`
}

// NewQuantilesDoublesSketchToQuantile new instance of QuantilesDoublesSketchToQuantile
func NewQuantilesDoublesSketchToQuantile() *QuantilesDoublesSketchToQuantile {
	q := &QuantilesDoublesSketchToQuantile{}
	q.SetType("quantilesDoublesSketchToQuantile")
	return q
}

// SetName set name
func (q *QuantilesDoublesSketchToQuantile) SetName(name string) *QuantilesDoublesSketchToQuantile {
	q.Base.SetName(name)
	return q
}

// SetFraction set fraction
func (q *QuantilesDoublesSketchToQuantile) SetFraction(fraction float64) *QuantilesDoublesSketchToQuantile {
	q.Fraction = fraction
	return q
}

// SetField set QuantilesDoublesSketchToQuantileField
func (q *QuantilesDoublesSketchToQuantile) SetField(field *QuantilesDoublesSketchToQuantileField) *QuantilesDoublesSketchToQuantile {
	q.Field = field
	return q
}

// NewQuantilesDoublesSketchToQuantileField new instance of QuantilesDoublesSketchToQuantileField
func NewQuantilesDoublesSketchToQuantileField() *QuantilesDoublesSketchToQuantileField {
	qf := &QuantilesDoublesSketchToQuantileField{}
	return qf
}

// SetType set type
func (qf *QuantilesDoublesSketchToQuantileField) SetType(typ string) *QuantilesDoublesSketchToQuantileField {
	qf.Type = typ
	return qf
}

// SetFieldName set fieldName
func (qf *QuantilesDoublesSketchToQuantileField) SetFieldName(fieldName string) *QuantilesDoublesSketchToQuantileField {
	qf.FieldName = fieldName
	return qf
}
