package brepprimapi

// #cgo LDFLAGS: -loccwrapper
// #include <occwrapper/brepprimapi.h>
import "C"
import (
	"github.com/marcuswu/gooccwrapper/gp"
	"github.com/marcuswu/gooccwrapper/topods"
)

type MakeRevol struct {
	makeRevol C.BRepPrimAPIMakeRevol
}

func NewMakeRevol(face topods.Face, axis gp.Ax1, degrees float64) MakeRevol {
	return MakeRevol{C.BRepPrimAPIMakeRevol_Init(
		C.TopoDSFace(face.Face),
		C.gpAx1(axis.Ax1),
		C.double(degrees),
	)}
}

func (mp MakeRevol) Shape() topods.Shape {
	return topods.NewShapeFromRef(topods.TopoDSShape(C.BRepPrimAPIMakeRevol_Shape(mp.makeRevol)))
}

func (mr MakeRevol) Free() {
	C.BRepPrimAPIMakeRevol_Free(mr.makeRevol)
}
