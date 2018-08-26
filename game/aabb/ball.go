package aabb

import "go-pong/game/space"

// Ball ball
type Ball struct {
	AABB
	velocity space.Velocity
}
