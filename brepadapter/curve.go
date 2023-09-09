package brepadapter

// #cgo LDFLAGS: -loccwrapper
// #include <occwrapper/brep_adapter.h>
// #include <occwrapper/occ_types.h>
import "C"
import (
	"github.com/marcuswu/gooccwrapper/gp"
	"github.com/marcuswu/gooccwrapper/topods"
)

type Curve struct {
	curve C.BRepAdapterCurve
}

func NewCurve(edge topods.Edge) Curve {
	return Curve{C.BRepAdapterCurve_Init(C.TopoDSEdge(edge.Edge))}
}

func (c Curve) IsLine() bool {
	return bool(C.BRepAdapterCurve_IsLine(c.curve))
}

func (c Curve) IsCircle() bool {
	return bool(C.BRepAdapterCurve_IsCircle(c.curve))
}

func (c Curve) IsEllipse() bool {
	return bool(C.BRepAdapterCurve_IsEllipse(c.curve))
}

func (c Curve) ToCircle() gp.Circ {
	return gp.NewCircFromRef(gp.GPCirc(C.BRepAdapterCurve_ToCircle(c.curve)))
}
