// Copyright (c) 2019, The Emergent Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package eve

import (
	"cogentcore.org/core/mat32"
)

// Rigid contains the full specification of a given object's basic physics
// properties including position, orientation, velocity.  These
type Rigid struct {

	// 1/mass -- 0 for no mass
	InvMass float32

	// COR or coefficient of restitution -- how elastic is the collision i.e., final velocity / initial velocity
	Bounce float32 `min:"0" max:"1"`

	// friction coefficient -- how much friction is generated by transverse motion
	Friction float32

	// record of computed force vector from last iteration
	Force mat32.Vec3

	// Last calculated rotational inertia matrix in local coords
	RotInertia mat32.Mat3
}

// Defaults sets defaults only if current values are nil
func (ps *Rigid) Defaults() {
}
