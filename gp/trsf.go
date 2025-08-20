package gp

// #cgo LDFLAGS: -loccwrapper
// #include <occwrapper/gp_trsf.h>
import "C"

type GPTrsf C.gpTrsf

type Trsf struct {
	Trsf C.gpTrsf
}

func NewTrsf() Trsf {
	return Trsf{C.gpTrsf_Init()}
}

func NewTrsfFromRef(ref GPTrsf) Trsf {
	return Trsf{C.gpTrsf(ref)}
}

func (t Trsf) Rotation() Quaternion {
	return Quaternion{C.gpTrsf_GetRotation(t.Trsf)}
}

func (t Trsf) SetTransformation(from Ax3, to Ax3) {
	C.gpTrsf_SetTransformation(t.Trsf, from.ax3, to.ax3)
}

func (t Trsf) SetRotation(axis Ax1, radians float64) {
	C.gpTrsf_SetRotation(t.Trsf, axis.Ax1, C.double(radians))
}

func (t Trsf) SetTranslation(vec Vec) {
	C.gpTrsf_SetTranslation(t.Trsf, vec.Vec)
}

func (t Trsf) SetMirrorAx1(axis Ax1) {
	C.gpTrsf_SetMirrorAx1(t.Trsf, axis.Ax1)
}

func (t Trsf) SetMirrorAx2(axis Ax2) {
	C.gpTrsf_SetMirrorAx2(t.Trsf, axis.Ax2)
}

func (t Trsf) Free() {
	C.gpTrsf_Free(t.Trsf)
}
