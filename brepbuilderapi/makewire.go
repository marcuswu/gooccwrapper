package brepbuilderapi

// #cgo LDFLAGS: -loccwrapper
// #include <occwrapper/brep_builder.h>
import "C"
import (
	"github.com/marcuswu/gooccwrapper/topods"
)

type MakeWire struct {
	makeWire C.BRepBuilderAPIMakeWire
}

func NewMakeWire() MakeWire {
	return MakeWire{C.BRepBuilderAPIMakeWire_Init()}
}

func NewMakeWireWithEdge(edge topods.Edge) MakeWire {
	return MakeWire{C.BRepBuilderAPIMakeWire_InitWithTopoDSEdge(C.TopoDSEdge(edge.Edge))}
}

func (me MakeWire) AddEdge(edge topods.Edge) {
	C.BRepBuilderAPIMakeWire_AddEdge(me.makeWire, C.TopoDSEdge(edge.Edge))
}

func (me MakeWire) ToTopoDSWire() topods.Wire {
	return topods.NewWireFromRef(topods.TopoDSWire(C.BRepBuilderAPIMakeWire_ToTopoDSWire(me.makeWire)))
}

func (me MakeWire) Free() {
	C.BRepBuilderAPIMakeWire_Free(me.makeWire)
}
