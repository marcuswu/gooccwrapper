package brepbuilderapi

// #cgo LDFLAGS: -loccwrapper
// #include <occwrapper/brep_builder.h>
import "C"
import (
	"github.com/marcuswu/gooccwrapper/gp"
	"github.com/marcuswu/gooccwrapper/topods"
)

type Transform struct {
	transform C.BRepBuilderAPITransform
}

func NewTransform(shape topods.Shape, trsf gp.Trsf) Transform {
	return Transform{C.BRepBuilderAPITransform_InitWithShapeTrsf(C.TopoDSShape(shape.Shape), C.gpTrsf(trsf.Trsf))}
}

func (t Transform) Shape() topods.Shape {
	return topods.NewShapeFromRef(topods.TopoDSShape(C.BRepBuilderAPITransform_Shape(t.transform)))
}

func (t Transform) Free() {
	C.BRepBuilderAPITransform_Free(t.transform)
}
