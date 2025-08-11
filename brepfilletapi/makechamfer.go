package brepfilletapi

// #cgo LDFLAGS: -loccwrapper
// #include <occwrapper/brepfilletapi.h>
import "C"
import (
	"github.com/marcuswu/gooccwrapper/topods"
)

type MakeChamfer struct {
	makeChamfer C.BRepFilletAPIMakeChamfer
}

func NewMakeChamfer(shape topods.TopoDSShape) MakeChamfer {
	return MakeChamfer{C.BRepFilletAPIMakeChamfer_Init(
		C.TopoDSShape(shape),
	)}
}

func (mf MakeChamfer) AddEdge(edge topods.TopoDSEdge, radius float64) {
	C.BRepFilletAPIMakeChamfer_Add(mf.makeChamfer, C.TopoDSEdge(edge), C.double(radius))
}

func (mf MakeChamfer) Shape() topods.Shape {
	return topods.NewShapeFromRef(topods.TopoDSShape(C.BRepFilletAPIMakeChamfer_Shape(mf.makeChamfer)))
}

func (mf MakeChamfer) Free() {
	C.BRepFilletAPIMakeChamfer_Free(mf.makeChamfer)
}
