package brepbuilderapi

// #cgo LDFLAGS: -loccwrapper
// #include <occwrapper/brep_builder.h>
import "C"
import (
	"github.com/marcuswu/gooccwrapper/topods"
)

type MakeFace struct {
	makeFace C.BRepBuilderAPIMakeFace
}

func NewMakeFace(wire topods.Wire) MakeFace {
	return MakeFace{C.BRepBuilderAPIMakeFace_InitWithWire(C.TopoDSWire(wire.Wire))}
}

func (me MakeFace) ToTopoDSFace() topods.Face {
	return topods.NewFaceFromRef(topods.TopoDSFace(C.BRepBuilderAPIMakeFace_ToTopoDSFace(me.makeFace)))
}

func (me MakeFace) Free() {
	C.BRepBuilderAPIMakeFace_Free(me.makeFace)
}
