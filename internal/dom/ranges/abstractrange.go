package ranges

// AbstractRange represents the base interface for ranges
// per WHATWG DOM Section 5.3
type AbstractRange interface {
	// StartContainer returns range's start node
	StartContainer() Node

	// StartOffset returns range's start offset
	StartOffset() int

	// EndContainer returns range's end node
	EndContainer() Node

	// EndOffset returns range's end offset
	EndOffset() int

	// Collapsed returns true if range is collapsed; otherwise false
	Collapsed() bool
}

// AbstractRangeImpl provides a base implementation of AbstractRange
type AbstractRangeImpl struct {
	start BoundaryPoint
	end   BoundaryPoint
}

// StartContainer getter steps are to return this's start node
func (r *AbstractRangeImpl) StartContainer() Node {
	return r.start.Node
}

// StartOffset getter steps are to return this's start offset
func (r *AbstractRangeImpl) StartOffset() int {
	return r.start.Offset
}

// EndContainer getter steps are to return this's end node
func (r *AbstractRangeImpl) EndContainer() Node {
	return r.end.Node
}

// EndOffset getter steps are to return this's end offset
func (r *AbstractRangeImpl) EndOffset() int {
	return r.end.Offset
}

// Collapsed getter steps are to return true if this is collapsed; otherwise false
func (r *AbstractRangeImpl) Collapsed() bool {
	return r.start.Node == r.end.Node && r.start.Offset == r.end.Offset
}

// SetStart sets the start boundary point
func (r *AbstractRangeImpl) SetStart(node Node, offset int) {
	r.start = BoundaryPoint{Node: node, Offset: offset}
}

// SetEnd sets the end boundary point
func (r *AbstractRangeImpl) SetEnd(node Node, offset int) {
	r.end = BoundaryPoint{Node: node, Offset: offset}
}

// GetStart returns the start boundary point
func (r *AbstractRangeImpl) GetStart() BoundaryPoint {
	return r.start
}

// GetEnd returns the end boundary point
func (r *AbstractRangeImpl) GetEnd() BoundaryPoint {
	return r.end
}
