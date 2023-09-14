package gp

// #cgo LDFLAGS: -loccwrapper
// #include <occwrapper/gp_quaternion.h>
import "C"

type Quaternion struct {
	quaternion C.gpQuaternion
}

func (q Quaternion) Inverted() Quaternion {
	return Quaternion{C.gpQuaternion_Inverted(q.quaternion)}
}

func (q Quaternion) Multiplied(o Quaternion) Quaternion {
	return Quaternion{C.gpQuaternion_Multiplied(q.quaternion, o.quaternion)}
}

func (q Quaternion) RotationAngle() float64 {
	return float64(C.gpQuaternion_GetRotationAngle(q.quaternion))
}

func (q Quaternion) Free() {
	C.gpQuaternion_Free(q.quaternion)
}
