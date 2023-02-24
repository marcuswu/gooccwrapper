package brepbuilderapi

// #cgo LDFLAGS: -loccwrapper
// #include <occwrapper/brep_builder.h>
import "C"
import (
	"github.com/marcuswu/gooccwrapper/pkg/geom"
	"github.com/marcuswu/gooccwrapper/pkg/topods"
)

type MakeEdge struct {
	makeEdge C.BRepBuilderAPIMakeEdge
}

func NewMakeEdge(curve geom.Curve) MakeEdge {
	return MakeEdge{C.BRepBuilderAPIMakeEdge_InitWithGeomCurve(C.GeomCurve(curve.Curve))}
}

func (me MakeEdge) ToTopoDSEdge() topods.Edge {
	return topods.NewEdgeFromRef(topods.TopoDSEdge(C.BRepBuilderAPIMakeEdge_ToTopoDSEdge(me.makeEdge)))
}

func (me MakeEdge) Free() {
	C.BRepBuilderAPIMakeEdge_Free(me.makeEdge)
}
