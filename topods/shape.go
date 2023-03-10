package topods

// #cgo LDFLAGS: -loccwrapper
// #include <occwrapper/topods_shape.h>
import "C"

type TopoDSShape C.TopoDSShape

type Shape struct {
	Shape C.TopoDSShape
}

func NewShapeFromRef(ref TopoDSShape) Shape {
	return Shape{C.TopoDSShape(ref)}
}

func (s Shape) Free() {
	C.TopoDSShape_Free(s.Shape)
}
