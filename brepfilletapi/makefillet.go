package brepfilletapi

// #cgo LDFLAGS: -loccwrapper
// #include <occwrapper/brepfilletapi.h>
import "C"
import (
	"github.com/marcuswu/gooccwrapper/topods"
)

type MakeFillet struct {
	makeFillet C.BRepFilletAPIMakeFillet
}

func NewMakeFillet(shape topods.TopoDSShape) MakeFillet {
	return MakeFillet{C.BRepFilletAPIMakeFillet_Init(
		C.TopoDSShape(shape),
	)}
}

func (mf MakeFillet) AddEdge(edge topods.TopoDSEdge, radius float64) {
	C.BRepFilletAPIMakeFillet_Add(mf.makeFillet, C.TopoDSEdge(edge), C.double(radius))
}

func (mf MakeFillet) Shape() topods.Shape {
	return topods.NewShapeFromRef(topods.TopoDSShape(C.BRepFilletAPIMakeFillet_Shape(mf.makeFillet)))
}

func (mf MakeFillet) Free() {
	C.BRepFilletAPIMakeFillet_Free(mf.makeFillet)
}
