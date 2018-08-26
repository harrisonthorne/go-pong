package main

import (
	mgl "github.com/go-gl/mathgl/mgl32"
)

// Camera /
type Camera struct {
	orthoMatrix mgl.Mat4
}

// UpdateScreenDimensions takes two integers that are the new
// screenWidth and screenHeight. This also updates the camera
// matrix.
func (c *Camera) UpdateScreenDimensions(width, height int) {
	c.updateOrthoMatrix(width, height)
}

func (c *Camera) updateOrthoMatrix(screenWidth, screenHeight int) {
	c.orthoMatrix = mgl.Ortho2D(0, float32(screenWidth), float32(screenHeight), 0)
}

// GetMatrix returns an orthogonal matrix.
func (c Camera) GetMatrix() mgl.Mat4 {
	return c.orthoMatrix
}
