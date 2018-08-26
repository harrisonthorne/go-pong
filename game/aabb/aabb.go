package aabb

// AABB is an Axis-Aligned Bounding Box
type AABB struct {
	X float32
	Y float32
	Width float32
	Height float32
}

// CollidesWith return true if the receiving AABB
// and the parameter AABB are in contact with each other
func (a AABB) CollidesWith(aabb AABB) bool {
	return a.X < aabb.X + aabb.Width &&
		a.X + a.Width > aabb.X &&
		a.Y < aabb.Y + aabb.Height &&
		a.Y + a.Height > aabb.Y
}
