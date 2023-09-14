package gp

// #cgo LDFLAGS: -loccwrapper
// #include <occwrapper/gp_trsf.h>
import "C"

type GPTrsf C.gpTrsf

type Trsf struct {
	trsf C.gpTrsf
}

func NewTrsf() Trsf {
	return Trsf{C.gpTrsf_Init()}
}

func NewTrsfFromRef(ref GPTrsf) Trsf {
	return Trsf{C.gpTrsf(ref)}
}

func (t Trsf) Rotation() Quaternion {
	return Quaternion{C.gpTrsf_GetRotation(t.trsf)}
}

func (t Trsf) SetTransformation(from Ax3, to Ax3) {
	C.gpTrsf_SetTransformation(t.trsf, from.ax3, to.ax3)
}

func (t Trsf) Free() {
	C.gpTrsf_Free(t.trsf)
}
