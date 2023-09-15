package gp

// #cgo LDFLAGS: -loccwrapper
// #include <occwrapper/gp_circ.h>
import "C"

type GPCirc C.gpCirc

type Circ struct {
	Circ C.gpCirc
}

func NewCirc(center Ax2, radius float64) Circ {
	return Circ{C.gpCirc_Init(center.Ax2, C.double(radius))}
}

func NewCircFromRef(ref GPCirc) Circ {
	return Circ{C.gpCirc(ref)}
}

func (c Circ) Radius() float64 {
	return float64(C.gpCirc_Radius(c.Circ))
}

func (c Circ) Location() Pnt {
	return Pnt{C.gpCirc_Location(c.Circ)}
}
