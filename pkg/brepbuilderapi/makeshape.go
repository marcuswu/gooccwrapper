package brepbuilderapi

// #cgo LDFLAGS: -loccwrapper
// #include <occwrapper/brep_builder.h>
import "C"
import (
	"github.com/marcuswu/gooccwrapper/pkg/topods"
)

type MakeShape struct {
	makeShape C.BRepBuilderAPIMakeShape
}

func (me MakeShape) Shape() topods.Shape {
	return topods.NewShapeFromRef(topods.TopoDSShape(C.BRepBuilderAPIMakeShape_Shape(me.makeShape)))
}

// I didn't define this in the C API -- seems like an oversight
// func (me MakeShape) Free() {
// 	C.BRepBuilderAPIMakeShape_Free(me.makeShape)
// }
